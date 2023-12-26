package balances

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"fuya-ark/internal/consts"
	constsError "fuya-ark/internal/consts/error"
	"fuya-ark/internal/dao"
	"fuya-ark/internal/model"
	"fuya-ark/internal/model/entity"
	"fuya-ark/internal/service"
	"fuya-ark/utility"
)

type sBalances struct{}

func init() {
	service.RegisterBalances(New())
}

func New() *sBalances {
	return &sBalances{}
}

// GetByUserId 根据用户id获取余额
func (s *sBalances) GetByUserId(ctx context.Context, userId int64) (balancesInfo *entity.BalanceInfo) {
	return dao.BalanceInfo.GetByUserId(ctx, userId)
}

// GetByUserIdWithTx 根据用户id获取余额
func (s *sBalances) GetByUserIdWithTx(ctx context.Context, userId int64, tx gdb.TX) (balancesInfo *entity.BalanceInfo) {
	balancesInfo = dao.BalanceInfo.GetByUserIdWithTx(ctx, userId, tx)
	return
}

// SumByDepletionTypeAndTime 根据用户id和流水类型统计
func (s *sBalances) SumByDepletionTypeAndTime(ctx context.Context,
	userId int64, depletionType int, createTime, endTime int64) (float64, error) {
	return dao.BalanceRecord.SumByDepletionTypeAndTime(ctx, userId, depletionType, createTime, endTime)
}

// SumByDepletionTypesAndTime 根据用户id和流水类型统计
func (s *sBalances) SumByDepletionTypesAndTime(ctx context.Context,
	userId int64, depletionType []int, gtTime int64) (float64, error) {
	return dao.BalanceRecord.SumByDepletionTypesAndTime(ctx, userId, depletionType, gtTime)
}

// SumDiamondsByTimeGroupByType 根据用户流水分类统计
func (s *sBalances) SumDiamondsByTimeGroupByType(ctx context.Context, createTime, endTime int64) (map[int64]int64, error) {
	return dao.BalanceRecord.SumDiamondsByTimeGroupByType(ctx, createTime, endTime)
}

// SumDiamondsByTimeAndType 根据用户流水分类统计
func (s *sBalances) SumDiamondsByTimeAndType(ctx context.Context, appChannel string, createTime, endTime int64, recordType int) (int64, error) {
	return dao.BalanceRecord.SumDiamondsByTimeAndType(ctx, appChannel, createTime, endTime, recordType)
}

// GetLiveAwardRecord 获取直播奖励记录
func (s *sBalances) GetLiveAwardRecord(ctx context.Context, userId int64, page, size int) []*model.LiveAwardRecord {
	liveAwardRecords, err := dao.BalanceRecord.GetLiveAwardRecord(ctx, userId, page, size)
	if err != nil {
		return nil
	}
	for _, liveAwardRecord := range liveAwardRecords {
		roomId := gconv.String(liveAwardRecord.UserId) + gconv.String(liveAwardRecord.LinkId)
		liveRoom, _ := dao.RoomLive.GetRoomDetailByRoomId(ctx, roomId)
		if liveRoom != nil {
			liveAwardRecord.LiveTime = liveRoom.EndTime - liveRoom.BeginTime
			liveAwardRecord.EndTime = liveRoom.EndTime
		}
	}
	return liveAwardRecords
}

// GetListByTimeAndGtId 获取一段时间余额流水
func (s *sBalances) GetListByTimeAndGtId(ctx context.Context, minId, startOfDay int64, endOfDay int64, size int) []*entity.BalanceRecord {
	return dao.BalanceRecord.GetListByTimeAndGtId(ctx, minId, startOfDay, endOfDay, size)
}

// MigrateBatch 迁移数据到月份归档
func (s *sBalances) MigrateBatch(ctx context.Context, month string, records []*entity.BalanceRecord) error {
	return dao.BalanceRecord.MigrateBatch(ctx, month, records)
}

// DeleteRecordByIds 删除流水记录
func (s *sBalances) DeleteRecordByIds(ctx context.Context, balanceRecordIds []int64) {
	dao.BalanceRecord.DeleteRecordByIds(ctx, balanceRecordIds)
}

// SumRemainDiamonds 获取总库存
func (s *sBalances) SumRemainDiamonds(ctx context.Context) (remainDiamonds int64) {
	remainDiamonds = dao.BalanceInfo.SumRemainDiamonds(ctx)
	return
}

// GetWeekNowAddDiamond 获取周流水是否显示填写whatsapp
func (s *sBalances) GetWeekNowAddDiamond(ctx context.Context, userId int64) (isShowWhatsApp bool) {
	weekBeginTime := utility.DateUtils.GetFirstDateOfWeek(time.Now())
	weekEndTime := utility.DateUtils.GetLastDateOfWeek(time.Now())

	rechargeDiamonds, _ := dao.BalanceRecord.SumByDepletionTypeAndTime(ctx, userId, consts.DepletionTypeRecharge, weekBeginTime.Unix(), weekEndTime.Unix())
	agentRechargeDiamonds, _ := dao.BalanceRecord.SumByDepletionTypeAndTime(ctx, userId, consts.AgentRecharge, weekBeginTime.Unix(), weekEndTime.Unix())
	if rechargeDiamonds+agentRechargeDiamonds >= 78000 {
		return true
	}
	return false
}

// Recharge 加钻/充值
// 参数说明:userId 当前用户Id
// fee: 花费金额（印尼盾，单位分，没有花费传递0）
// diamonds: 钻石数
// linkeId: 业务Id
// depletionType: 业务类型
// tx: 事务对象
func (s *sBalances) Recharge(ctx context.Context,
	userId, otherUserId, fee, diamonds, linkId int64, depletionType uint, tx gdb.TX) (logId int64, err error) {
	if userId <= 0 || diamonds <= 0 || linkId <= 0 || depletionType <= 0 || tx == nil {
		g.Log().Warningf(ctx,
			"添加加钻记录参数异常 userId:%v diamonds:%v linkId:%v depletionType:%v tx:%v",
			userId, diamonds, linkId, depletionType, tx,
		)
		return 0, fmt.Errorf("添加加钻记录参数异常 userId:%v diamonds:%v linkId:%v depletionType:%v tx:%v",
			userId, diamonds, linkId, depletionType, tx)
	}
	//添加日志
	logId, err = addBalanceRecord(ctx, uint64(userId), uint64(otherUserId), uint64(diamonds), depletionType, uint64(linkId), consts.BalanceTypeIncome, tx)
	if err != nil {
		g.Log().Infof(ctx, "添加加钻记录出现错误，userId:%v", userId)
		return 0, err
	}
	balanceInfo := &entity.BalanceInfo{}
	//只有充值才记入总购买钻数
	if depletionType == consts.DepletionTypeRecharge {
		balanceInfo.TotalPayDiamonds = uint64(diamonds)
	}
	balanceInfo.RemainDiamonds = uint64(diamonds)
	balanceInfo.TotalFee = uint64(fee)
	balanceInfo.UserId = userId
	err = dao.BalanceInfo.UpdateOrInsert(ctx, balanceInfo, tx)
	if err != nil {
		g.Log().Warning(ctx, "Recharge 改变diamonds值:%s ", err)
		return 0, err
	}
	return logId, nil
}

// Charge 消费/扣钻
// 参数说明:userId 当前用户Id
// targetUserId:目标用户（收益方Id,没有则传递0）
// diamonds: 钻石数
// linkId: 业务Id
// depletionType: 业务类型
// tx: 事务对象
func (s *sBalances) Charge(ctx context.Context,
	userId, targetUserId, diamonds, linkId int64, depletionType uint, tx gdb.TX) (logId int64, err error) {
	if userId <= 0 || diamonds <= 0 || linkId <= 0 || depletionType <= 0 || tx == nil {
		g.Log().Infof(ctx,
			"Charge 参数异常，userId:%v,linkId:%v,diamonds:%v,depletionType:%v,tx:%v",
			userId, linkId, diamonds, depletionType, tx,
		)
		return 0, fmt.Errorf("Charge 参数异常，userId:%v,linkId:%v,diamonds:%v,depletionType:%v,tx:%v",
			userId, linkId, diamonds, depletionType, tx)
	}
	oldBalanceInfo := dao.BalanceInfo.GetByUserId(ctx, userId)
	if oldBalanceInfo.RemainDiamonds < uint64(diamonds) { //余额不足
		return 0, constsError.InsufficientBalanceError
	}
	//添加日志
	logId, err = addBalanceRecord(ctx, uint64(userId), uint64(targetUserId), uint64(diamonds), depletionType, uint64(linkId), consts.BalanceTypeExpend, tx)
	if err != nil {
		g.Log().Infof(ctx, "添加消费记录出现错误，userId:%v", userId)
		return 0, err
	}
	affected, err := dao.BalanceInfo.UpdateBalance(ctx, uint64(userId), uint64(diamonds), tx)
	if err != nil {
		g.Log().Errorf(ctx, "扣钻失败 err:%v", err)
		return 0, constsError.InsufficientBalanceError
	}
	if affected == 1 {
		return logId, nil
	}
	g.Log().Infof(ctx, "扣钻失败userId:%v", userId)
	return 0, constsError.InsufficientBalanceError
}

func addBalanceRecord(ctx context.Context, sourceUserId, targetUserId, diamonds uint64,
	depletionType uint, linkId uint64, balanceRecordType uint, tx gdb.TX) (logId int64, err error) {
	balance, err := dao.BalanceInfo.GetBalanceInfoByUserid(ctx, int64(sourceUserId), tx)
	if err != nil {
		g.Log().Warningf(ctx, "GetBalanceInfoByUserid: %v ", err)
		return 0, err
	}
	if balance == nil {
		balance = &entity.BalanceInfo{}
	}
	balanceRecord := entity.BalanceRecord{}
	balanceRecord.SourceUserId = sourceUserId
	balanceRecord.TargetUserId = targetUserId
	balanceRecord.DepletionType = depletionType
	balanceRecord.Diamonds = diamonds
	balanceRecord.BeforeChange = balance.RemainDiamonds
	if balanceRecordType == consts.BalanceTypeIncome {
		balanceRecord.AfterChange = balance.RemainDiamonds + diamonds
	} else {
		balanceRecord.AfterChange = balance.RemainDiamonds - diamonds
	}

	balanceRecord.LinkId = linkId
	balanceRecord.Type = balanceRecordType
	balanceRecord.CreateTime = time.Now().Unix()
	balanceRecord.UpdateTime = time.Now().Unix()
	logId, err = dao.BalanceRecord.Insert(ctx, balanceRecord, tx)
	if err != nil {
		g.Log().Warningf(ctx, "addBalanceRecord: %v ", err)
		return 0, err
	}
	return logId, nil
}
