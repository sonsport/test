// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"fuya-ark/internal/model"
	"fuya-ark/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IBalances interface {
		// GetByUserId 根据用户id获取余额
		GetByUserId(ctx context.Context, userId int64) (balancesInfo *entity.BalanceInfo)
		// GetByUserIdWithTx 根据用户id获取余额
		GetByUserIdWithTx(ctx context.Context, userId int64, tx gdb.TX) (balancesInfo *entity.BalanceInfo)
		// SumByDepletionTypeAndTime 根据用户id和流水类型统计
		SumByDepletionTypeAndTime(ctx context.Context, userId int64, depletionType int, createTime, endTime int64) (float64, error)
		// SumByDepletionTypesAndTime 根据用户id和流水类型统计
		SumByDepletionTypesAndTime(ctx context.Context, userId int64, depletionType []int, gtTime int64) (float64, error)
		// SumDiamondsByTimeGroupByType 根据用户流水分类统计
		SumDiamondsByTimeGroupByType(ctx context.Context, createTime, endTime int64) (map[int64]int64, error)
		// SumDiamondsByTimeAndType 根据用户流水分类统计
		SumDiamondsByTimeAndType(ctx context.Context, appChannel string, createTime, endTime int64, recordType int) (int64, error)
		// GetLiveAwardRecord 获取直播奖励记录
		GetLiveAwardRecord(ctx context.Context, userId int64, page, size int) []*model.LiveAwardRecord
		// GetListByTimeAndGtId 获取一段时间余额流水
		GetListByTimeAndGtId(ctx context.Context, minId, startOfDay int64, endOfDay int64, size int) []*entity.BalanceRecord
		// MigrateBatch 迁移数据到月份归档
		MigrateBatch(ctx context.Context, month string, records []*entity.BalanceRecord) error
		// DeleteRecordByIds 删除流水记录
		DeleteRecordByIds(ctx context.Context, balanceRecordIds []int64)
		// SumRemainDiamonds 获取总库存
		SumRemainDiamonds(ctx context.Context) (remainDiamonds int64)
		// GetWeekNowAddDiamond 获取周流水是否显示填写whatsapp
		GetWeekNowAddDiamond(ctx context.Context, userId int64) (isShowWhatsApp bool)
		// Recharge 加钻/充值
		// 参数说明:userId 当前用户Id
		// fee: 花费金额（印尼盾，单位分，没有花费传递0）
		// diamonds: 钻石数
		// linkeId: 业务Id
		// depletionType: 业务类型
		// tx: 事务对象
		Recharge(ctx context.Context, userId, otherUserId, fee, diamonds, linkId int64, depletionType uint, tx gdb.TX) (logId int64, err error)
		// Charge 消费/扣钻
		// 参数说明:userId 当前用户Id
		// targetUserId:目标用户（收益方Id,没有则传递0）
		// diamonds: 钻石数
		// linkId: 业务Id
		// depletionType: 业务类型
		// tx: 事务对象
		Charge(ctx context.Context, userId, targetUserId, diamonds, linkId int64, depletionType uint, tx gdb.TX) (logId int64, err error)
	}
)

var (
	localBalances IBalances
)

func Balances() IBalances {
	if localBalances == nil {
		panic("implement not found for interface IBalances, forgot register?")
	}
	return localBalances
}

func RegisterBalances(i IBalances) {
	localBalances = i
}
