// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"golang.org/x/net/context"
	"golang.org/x/tools/container/intsets"

	"fuya-ark/internal/consts"
	"fuya-ark/internal/dao/internal"
	"fuya-ark/internal/model"
	"fuya-ark/internal/model/bo"
	"fuya-ark/internal/model/do"
	"fuya-ark/internal/model/entity"
)

// internalRoomLiveDao is internal type for wrapping internal DAO implements.
type internalRoomLiveDao = *internal.RoomLiveDao

// roomLiveDao is the data access object for table room_live.
// You can define custom methods on it to extend its functionality as you wish.
type roomLiveDao struct {
	internalRoomLiveDao
}

var (
	// RoomLive is globally public accessible object for table room_live operations.
	RoomLive = roomLiveDao{
		internal.NewRoomLiveDao(),
	}
)

// Fill with you ideas below.

// CreateLiveRoom 创建直播间
func (m *roomLiveDao) CreateLiveRoom(ctx context.Context, data entity.RoomLive) (int64, error) {
	data.CreateTime = time.Now().Unix() //创建时间就是预开播时间
	data.UpdateTime = time.Now().Unix()
	data.BeginTime = time.Now().Unix()
	data.Status = consts.RoomStatusPre
	res, err := m.DB().Ctx(ctx).Model().Data(data).Insert()
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func (m *roomLiveDao) RestartLiveRoom(ctx context.Context, id int64, status int) error {
	_, err := m.DB().Ctx(ctx).Model().Data(
		do.RoomLive{Id: id, UpdateTime: time.Now().Unix(), Status: status}).Update()
	if err != nil {
		return err
	}
	return nil
}

// GetActiveLiveRoomList 查询直播中的直播间
func (m *roomLiveDao) GetActiveLiveRoomList(ctx context.Context, ids []int64) ([]*entity.RoomLive, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	db := m.DB().Ctx(ctx).Model().
		Fields("id,user_id,room_name,image_path,room_id,stream_id").
		Where("status = ? and last_heartbeat_time >= ?", consts.RoomStatusLiving, time.Now().Unix()-60).
		WhereIn("user_id", ids)
	var res []*entity.RoomLive
	err := db.Scan(&res)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return res, nil
}

// GetLiveRoomListByUid 根据主播ID查询直播间列表
func (m *roomLiveDao) GetLiveRoomListByUid(ctx context.Context, userId int64) ([]*entity.RoomLive, error) {
	var res []*entity.RoomLive
	err := m.DB().Ctx(ctx).Model().
		Fields("`id`,`room_name`,`image_path`,`member_count`,`status`,`coin_amount`,`begin_time`,`end_time`").
		Where("user_id = ?", userId).
		Scan(&res)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return res, nil
}

// CloseLiveRoom 主播关闭直播间更新相关信息
func (m *roomLiveDao) CloseLiveRoom(ctx context.Context, live entity.RoomLive) error {
	curTime := time.Now().Unix()
	robotMemberCount, err := g.Redis().ZCard(ctx, consts.RoomRobotExisted+live.RoomId)
	data := g.Map{
		m.Columns().Status:                   live.Status,
		m.Columns().Score:                    live.Score,
		m.Columns().MemberCount:              live.MemberCount - robotMemberCount,
		m.Columns().SendGiftPersonCount:      live.SendGiftPersonCount,
		m.Columns().SendLuckyGiftPersonCount: live.SendLuckyGiftPersonCount,
		m.Columns().CoinAmount:               live.CoinAmount,
		m.Columns().FocusCount:               live.FocusCount,
		m.Columns().EndTime:                  curTime,
		m.Columns().LastHeartbeatTime:        curTime,
		m.Columns().UpdateTime:               curTime,
	}
	_, err = m.DB().Ctx(ctx).Model().
		Where(m.Columns().RoomId, live.RoomId).
		Cache(gdb.CacheOption{Duration: -1, Name: consts.MysqlRoomInfoByRoomId + live.RoomId}).
		Data(data).
		Update()
	return err
}

// GetWillCloseLiveRoom 将直播中、预开播的直播间查询出来
func (m *roomLiveDao) GetWillCloseLiveRoom(ctx context.Context, userId int64) (liveRoomList []*entity.RoomLive, err error) {
	err = m.DB().Ctx(ctx).Model().
		Where("user_id = ? and (status = ? or status = ?)", userId, consts.RoomStatusLiving,
			consts.RoomStatusPre).
		Scan(&liveRoomList)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return liveRoomList, nil
}

// BatchCloseLiveRoom 批量关闭直播间 （开播中、预开播）
func (m *roomLiveDao) BatchCloseLiveRoom(ctx context.Context, roomId string, roomStatus int) error {
	_, err := m.DB().Ctx(ctx).Model().
		Data("status = ?,end_time = ?", roomStatus, time.Now().Unix()).
		Where("room_id = ? and (status = ? or status = ?)",
			roomId, consts.RoomStatusLiving, consts.RoomStatusPre).
		Update()
	return err
}

// BeginLive 开始直播
func (m *roomLiveDao) BeginLive(ctx context.Context, userId int64, req *model.BeginLiveReq) (int64, error) {
	notTime := time.Now().Unix()
	updateData := g.Map{"last_heartbeat_time": notTime}
	updateData["status"] = consts.RoomStatusLiving
	updateData["room_name"] = req.RoomName
	updateData["image_path"] = req.ImagePath
	updateData["begin_time"] = notTime
	updateData["game_id"] = req.GameId
	updateData["private_status"] = req.PrivateStatus
	updateData["password"] = req.Password
	updateData["unlock_price"] = req.UnlockPrice
	updateData["wish_gift_id"] = req.WishGiftId
	updateData["wish_gift_count"] = req.WishGiftCount
	updateData["room_type"] = req.RoomType
	updateData["max_multiple"] = req.MaxMultiple
	updateData["layout_set"] = req.LayoutSet
	sqlResult, err := m.DB().Ctx(ctx).Model().
		Data(updateData).
		Where("room_id = ? and user_id = ? and status = ?", req.RoomId, userId, consts.RoomStatusPre).
		Update()
	if err != nil {
		return 0, err
	}
	affected, _ := sqlResult.RowsAffected()
	return affected, nil
}

func (m *roomLiveDao) UpdateRoomLive(ctx context.Context, roomId string,
	userId int64, password string, privateStatus, unlockPrice int) {
	notTime := time.Now().Unix()
	updateData := g.Map{"last_heartbeat_time": notTime}
	updateData["private_status"] = privateStatus
	updateData["password"] = password
	updateData["unlock_price"] = unlockPrice
	m.DB().Ctx(ctx).Model().
		Data(updateData).
		Cache(gdb.CacheOption{Duration: -1, Name: consts.MysqlRoomInfoByRoomId + roomId}).
		Where("room_id = ? and user_id = ? and status = ?", roomId, userId, consts.RoomStatusLiving).
		Update()

}

func (m *roomLiveDao) UpdateStatusByRoomIdAndUserId(ctx context.Context, userId int64, roomId string, status int) (int64, error) {
	sqlResult, err := m.DB().Ctx(ctx).Model().
		Where("room_id = ? and user_id = ?", roomId, userId).
		Data("status = ?", status).Update()
	if err != nil {
		return 0, err
	}
	affected, _ := sqlResult.RowsAffected()
	return affected, nil
}

func (m *roomLiveDao) UpdateDataByScore(ctx context.Context, roomId string, score int) error {
	_, err := m.DB().Ctx(ctx).Model().
		Where("room_id = ", roomId).Increment("score", score)
	if err != nil {
		return err
	}
	return nil
}

// UpdateImagePath 更新图片地址
func (m *roomLiveDao) UpdateImagePath(ctx context.Context, userId int64, roomId string, path string) error {
	_, err := m.DB().Ctx(ctx).Model().
		Where("room_id = ? and user_id = ?", roomId, userId).
		Data("image_path = ?", path).
		Update()
	return err
}

// RoomHeartBeat 更新心跳包
func (m *roomLiveDao) RoomHeartBeat(ctx context.Context, roomId string) error {
	_, err := m.DB().Ctx(ctx).Model().
		Where("room_id = ?", roomId).
		Cache(gdb.CacheOption{
			Duration: -1,
			Name:     consts.MysqlRoomInfoByRoomId + roomId,
		}).
		Data("last_heartbeat_time = ?", time.Now().Unix()).
		Update()
	return err
}

// AddMember 更新观看人数
func (m *roomLiveDao) AddMember(ctx context.Context, userId int64, roomId string, amt int64) error {
	_, err := m.DB().Ctx(ctx).Model().
		Where("room_id = ? and user_id = ? and status = ?", roomId, userId, consts.RoomStatusLiving).
		Increment("member_count", amt)
	return err
}

// AddFans 更新粉丝数
func (m *roomLiveDao) AddFans(ctx context.Context, roomId string, amt int64) error {
	_, err := m.DB().Ctx(ctx).Model().
		Where("room_id = ? and status = ?", roomId, consts.RoomStatusLiving).
		Increment("focus_count", amt)
	return err
}

// AddCoinAmount 更新金币数
func (m *roomLiveDao) AddCoinAmount(ctx context.Context, roomId string, amt int64) error {
	_, err := m.DB().Ctx(ctx).Model().
		Where("room_id = ? and status = ?", roomId, consts.RoomStatusLiving).
		Increment("coin_amount", amt)
	return err
}

// GetRoomByStreamId 根据GetRoomByStreamId获取直播间详情
func (m *roomLiveDao) GetRoomByStreamId(ctx context.Context, streamId string) (*entity.RoomLive, error) {
	var res *entity.RoomLive
	err := m.DB().Ctx(ctx).Model().
		Where("stream_id = ?", streamId).
		Cache(gdb.CacheOption{
			Duration: time.Second * 10, //缓存10秒
			Name:     consts.MysqlRoomInfoByRoomId + streamId,
		}).Scan(&res)
	if err != nil && err != sql.ErrNoRows {
		g.Log().Errorf(ctx, "GetRoomByStreamId error = %v", err)
		return nil, err
	}
	return res, nil
}

// GetRoomDetailByRoomId 根据roomId获取直播间详情
func (m *roomLiveDao) GetRoomDetailByRoomId(ctx context.Context, roomId string) (*entity.RoomLive, error) {
	var res *entity.RoomLive
	err := m.DB().Ctx(ctx).Model().
		Where("room_id = ?", roomId).
		Cache(gdb.CacheOption{
			Duration: time.Second * 30, Name: consts.MysqlRoomInfoByRoomId + roomId,
		}).Scan(&res)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		g.Log().Errorf(ctx, "GetRoomDetailByRoomId error = %v", err)
		return nil, err
	}
	return res, nil
}

// GetRoomsByRoomIds 根据roomId获取直播间详情
func (m *roomLiveDao) GetRoomsByRoomIds(ctx context.Context, roomIds []string) ([]*entity.RoomLive, error) {
	var res []*entity.RoomLive
	err := m.DB().Ctx(ctx).Model().WhereIn("room_id", roomIds).Scan(&res)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		g.Log().Errorf(ctx, "GetRoomDetailByRoomId error = %v", err)
		return nil, err
	}
	return res, nil
}

func (m *roomLiveDao) GetRoomDetailByRoomIdAndUserId(ctx context.Context, roomId string, userId int64) (*entity.RoomLive, error) {
	var res *entity.RoomLive
	err := m.DB().Ctx(ctx).Model().
		Fields("coin_amount,begin_time,focus_count,member_count,game_id").
		Where("room_id = ? and user_id=?", roomId, userId).
		Scan(&res)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return res, nil
}

// GetActiveLiveRoom 查询直播中的直播间
func (m *roomLiveDao) GetActiveLiveRoom(ctx context.Context,
	size, page int, role []int, userIds, blackIds []int64, homeType int) ([]*entity.RoomLive, error) {

	var res []*entity.RoomLive
	limit := size
	offset := page * size

	var orderBy string
	if homeType == consts.NewList {
		orderBy = "create_time"
	} else {
		orderBy = "score"
	}

	var timeNow = time.Now().Unix() - 90
	var condition = "status = ? and last_heartbeat_time >= ? "
	var args []interface{}
	args = append(args, consts.RoomStatusLiving, timeNow)
	if len(blackIds) == 0 {
		if len(role) == 1 {
			if len(userIds) == 0 {
				condition += "and role = ?"
				args = append(args, role[0])
			} else {
				condition += "and role = ? and user_id in (?)"
				args = append(args, role[0], userIds)
			}
		} else {
			if len(userIds) == 0 {
				condition += "and role in (?)"
				args = append(args, role)
			} else {
				condition += "and role in (?) and user_id in (?)"
				args = append(args, role, userIds)
			}
		}
	} else {
		if len(role) == 1 {
			condition += "and role = ? and user_id not in (?)"
			args = append(args, role[0], blackIds)
		} else {
			condition += "and role in (?) and user_id not in (?)"
			args = append(args, role, blackIds)
		}
	}

	var fields = "user_id,room_name,image_path,room_id,stream_id,member_count,battle_status,private_status"
	err := m.DB().Ctx(ctx).Model().
		Fields(fields).
		Where(condition, args).
		Limit(offset, limit).
		OrderDesc(orderBy).
		Scan(&res)
	return res, err
}

// GetRoomByUseIdAndStatus 根据房间和状态获取直播
func (m *roomLiveDao) GetRoomByUseIdAndStatus(ctx context.Context, userId int64, status int) (res *entity.RoomLive, err error) {
	err = m.DB().Ctx(ctx).Model().
		Where("user_id = ? and status = ?", userId, status).
		OrderDesc("last_heartbeat_time").
		Limit(1).
		Scan(&res)
	return
}

// GetAllGameRoom 获取所有开播中的游戏房
func (m *roomLiveDao) GetAllGameRoom(ctx context.Context) (roomIds []string, err error) {
	resData, err := m.DB().Ctx(ctx).Model().
		Where("status = 1 and game_id>0").
		OrderDesc("last_heartbeat_time").
		Fields("room_id").All()
	for _, i2 := range resData.List() {
		roomIds = append(roomIds, gconv.String(i2["room_id"]))
	}
	return
}

// UpdateRoomBattleStatus 更新房间pk状态
func (m *roomLiveDao) UpdateRoomBattleStatus(ctx context.Context, roomId string, battleStatus int) error {
	_, err := m.DB().Ctx(ctx).Model().
		Where("room_id = ?", roomId).
		Cache(gdb.CacheOption{
			Duration: -1,
			Name:     consts.MysqlRoomInfoByRoomId + roomId,
		}).
		Data("battle_status = ?", battleStatus).
		Update()
	return err
}

// GetRoomIdsByStatus 根据状态获取房间id列表
func (m *roomLiveDao) GetRoomIdsByStatus(ctx context.Context, status int) (roomIds []string) {
	result, err := m.DB().Ctx(ctx).Model().
		Ctx(ctx).
		Fields("room_id").
		Where("`status` = ?", status).
		Where("private_status", consts.PrivateStatusNormal).
		All()
	if err != nil {
		return nil
	}
	roomIds = make([]string, 0, result.Size())
	for _, roomId := range result.Array() {
		roomIds = append(roomIds, roomId.String())
	}
	return roomIds
}

func (m *roomLiveDao) GetUserIdsByTime(ctx context.Context, startTime, endTime int64) (userIds []int64, err error) {
	result, err := m.DB().Ctx(ctx).Model().
		Fields("user_id").
		WhereGTE("update_time", startTime).
		WhereOr("status", consts.RoomStatusLiving).
		//WhereLTE("create_time", endTime).
		Distinct().
		All()
	if err != nil {
		return nil, err
	}
	userIds = make([]int64, 0, result.Size())
	for _, roomId := range result.Array() {
		userIds = append(userIds, roomId.Int64())
	}
	return userIds, nil
}

func (m *roomLiveDao) GetUserIdsByTimeAndLabel(ctx context.Context, startTime int64, labels []string) (userIds []int64, err error) {
	result, err := m.DB().Ctx(ctx).Model(m.Table(), "rl").
		LeftJoin("guild_anchor ga", "rl.user_id=ga.user_id").
		Where("NOT EXISTS(SELECT al.user_id from anchor_level al WHERE al.user_id=rl.user_id and al.level in ('A','AA'))").
		Fields("rl.user_id").WhereIn("ga.label", labels).
		Where("rl.update_time>? or rl.status=?", startTime, consts.RoomStatusLiving).
		//WhereLTE("create_time", endTime).
		Distinct().
		All()
	if err != nil {
		return nil, err
	}
	userIds = make([]int64, 0, result.Size())
	for _, roomId := range result.Array() {
		userIds = append(userIds, roomId.Int64())
	}
	return userIds, nil
}

func (m *roomLiveDao) GetUserIdsByTimeAndLevel(ctx context.Context, startTime int64, levels []string) (userIds []int64, err error) {
	result, err := m.DB().Ctx(ctx).Model(m.Table(), "rl").
		LeftJoin("anchor_level al", "al.user_id=rl.user_id").
		Fields("rl.user_id").
		WhereIn("al.level", levels).
		Where("rl.update_time>? or rl.status=?", startTime, consts.RoomStatusLiving).
		//WhereLTE("create_time", endTime).
		Distinct().
		All()
	if err != nil {
		return nil, err
	}
	userIds = make([]int64, 0, result.Size())
	for _, roomId := range result.Array() {
		userIds = append(userIds, roomId.Int64())
	}
	return userIds, nil
}

// GetRoomByTimeout 获取超时房间的数据
//func (m *roomLiveDao) GetRoomByTimeout(ctx context.Context, status int, interval int64) ([]*LiveRoomModel, error) {
//	var res []*LiveRoomModel
//	err := storage.HimiMysqlClient.Model(m.TableName()).Ctx(ctx).
//		Where("last_heartbeat_time <= ? and status=?", time.Now().Unix()-interval, status).
//		Scan(&res)
//	if err != nil {
//		return nil, err
//	}
//	return res, nil
//}

func (m *roomLiveDao) GetOldRoomNameAndImage(ctx context.Context, userId int64) (room *roomLiveDao, err error) {
	err = m.DB().Ctx(ctx).Model(m.Table()).
		Where("room_name <> ''").
		Where("user_id", userId).
		Where("private_status", 0).
		OrderDesc("create_time").
		Scan(&room)
	return
}

// GetRoomLivesStatisticalUserId 获取用户昨日开播的数据
func (m *roomLiveDao) GetRoomLivesStatisticalUserId(ctx context.Context,
	userId int64, currentDayOfStart, yesterdayOfStart time.Time) ([]*bo.LiveRoomBO, error) {

	var res []*bo.LiveRoomBO
	sqlStr := `
		SELECT * FROM room_live WHERE begin_time <= ? AND end_time = 0 AND user_id = ? AND status = 1
		UNION ALL
		SELECT * FROM room_live WHERE begin_time <= ? AND end_time >= ? AND user_id = ? AND status IN ( 2, 4, 5 )
		UNION ALL
		SELECT * FROM room_live WHERE begin_time <= ? AND end_time <= ? AND end_time >= ? AND user_id = ? AND status IN ( 2, 4, 5) order by begin_time 
	`
	result, err := m.DB().Ctx(ctx).Query(ctx, sqlStr,
		currentDayOfStart.Unix(), userId,
		currentDayOfStart.Unix(), currentDayOfStart.Unix(), userId,
		currentDayOfStart.Unix(), currentDayOfStart.Unix(), yesterdayOfStart.Unix(), userId,
	)
	if err != nil {
		return nil, err
	}
	err = result.Structs(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetUserCloseLiveListByDay 查询用户某天关播的直播数据
func (m *roomLiveDao) GetUserCloseLiveListByDay(ctx context.Context, userId int64, yesterdayDate int64, todayDate int64) ([]*bo.LiveRoomBO, error) {
	var res []*bo.LiveRoomBO
	err := m.DB().Ctx(ctx).Model(m.Table()).
		Where("user_id", userId).
		WhereGT("end_time", yesterdayDate).
		WhereLT("begin_time", todayDate).
		WhereIn("status", g.Slice{consts.RoomStatusEnd, consts.RoomStatusPollOut, consts.RoomStatusPreOut}).
		OrderDesc("create_time").
		Scan(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetRoomByUseIdAndPrivateStatus 获取小黑屋房间
func (m *roomLiveDao) GetRoomByUseIdAndPrivateStatus(ctx context.Context, userId int64) (data *entity.RoomLive, err error) {
	err = m.DB().Ctx(ctx).Model(m.Table()).
		Where("user_id", userId).
		Where("private_status", 1).
		Scan(&data)
	if err != nil {
		return nil, err
	}
	if err != nil {
		g.Log().Errorf(ctx, "GetRoomByUseIdAndPrivateStatus 获取小黑屋房间 err%v", err)
		return nil, err
	}
	if data == nil {
		data = m.CreatePrivateLiveRoom(ctx, userId)
	}
	return data, err
}

func getRoomId(userId int64) string {
	//return fmt.Sprintf("%d%d%d", time.Now().UnixNano()/1000+rand.Int()%1000)
	//return fmt.Sprintf("%v", ztool.IdUtils.GenerateSnowflakeId())
	return fmt.Sprintf("%d%d%03d", userId, time.Now().Unix(), grand.Intn(intsets.MaxInt)%1000)
}

// CreatePrivateLiveRoom 创建小黑屋
func (m *roomLiveDao) CreatePrivateLiveRoom(ctx context.Context, userId int64) *entity.RoomLive {
	liveRoom := entity.RoomLive{
		UserId:        userId,
		RoomName:      "Small black room",
		RoomId:        getRoomId(userId),
		StreamId:      "",
		Role:          consts.UserRobot,
		Status:        consts.RoomStatusLeave,
		PrivateStatus: consts.PrivateStatusSmallBlack,
		CreateTime:    time.Now().Unix(),
		UpdateTime:    time.Now().Unix(),
	}
	lastId, err := m.DB().Ctx(ctx).Model().Data(liveRoom).InsertAndGetId()
	if err != nil {
		g.Log().Errorf(ctx, "CreatePrivateLiveRoom err:%v", err)
		return nil
	}
	liveRoom.Id = uint64(lastId)
	return &liveRoom
}

func (m *roomLiveDao) GetRoomIdsByClassifyType(ctx context.Context, classifyType, guildId int64, review bool) ([]string, error) {
	// classifyType 1 带标签的主播 2 不带标签的主播 3 密码房、钻石房 4 7天认证的主播 5 加工会未加星 6 用户 7 pesona 标签
	sqlStr := ""
	if classifyType == 0 { //不是私密房
		sqlStr = "SELECT rl.room_id FROM room_live rl WHERE rl.`status` =1  "
		if time.Now().Hour() > 3 && time.Now().Hour() < 20 { //特定时间点钻石房密码房开播在热门里显示
			sqlStr += " and rl.private_status=0"
		}
		if review {
			sqlStr = "SELECT rl.room_id FROM room_live rl WHERE not exists(select ga.user_id from guild_anchor ga where ga.user_id = rl.user_id and ga.label in ('%v')) and rl.`status` =1 and rl.private_status=0 "
			sqlStr = fmt.Sprintf(sqlStr, "Pretty','"+strings.Join(consts.CharmList, "','"))
		}
	} else if classifyType == 1 { //带标签的主播
		sqlStr = "SELECT rl.room_id FROM room_live rl LEFT JOIN guild_anchor ga ON ga.user_id = rl.user_id  WHERE ga.`status` = 1 AND rl.`status` =1 and ga.label_type in (3,6,7,8,9)"
	} else if classifyType == 2 { //不带标签的主播
		sqlStr = "SELECT rl.room_id FROM room_live rl LEFT JOIN guild_anchor ga ON ga.user_id = rl.user_id  WHERE ga.`status` = 1 AND rl.`status` =1 and ga.label_type=0"
	} else if classifyType == 3 { //密码房、钻石房
		sqlStr = "SELECT rl.room_id FROM room_live rl where private_status in (2,3) AND rl.`status` =1"
	} else if classifyType == 4 { //14天认证的主播
		weekBefore := time.Now().AddDate(0, 0, -14).Unix()
		sqlStr = "SELECT rl.room_id FROM room_live rl LEFT JOIN guild_anchor ga ON ga.user_id = rl.user_id  WHERE ga.star_level>0 and ga.create_time > " + gconv.String(weekBefore) + " AND rl.`status` =1"
	} else if classifyType == 5 { //加工会未加星
		sqlStr = "SELECT rl.room_id FROM room_live rl LEFT JOIN guild_anchor ga ON ga.user_id = rl.user_id  WHERE ga.guild_id>0 and ga.star_level=0 AND rl.`status` =1"
	} else if classifyType == 6 { //用户
		sqlStr = "SELECT rl.room_id FROM room_live rl where not exists(select ga.user_id from  guild_anchor ga where ga.user_id = rl.user_id) and rl.`status` =1"
	} else if classifyType == 7 { //pesona标签
		sqlStr = "SELECT rl.room_id FROM room_live rl where (exists(select ga.user_id from  guild_anchor ga where ga.user_id = rl.user_id and ga.label in ('%v')) or exists(select al.user_id from anchor_level al where al.user_id=rl.user_id and al.level in('A','AA'))) and rl.`status` =1"
		sqlStr = fmt.Sprintf(sqlStr, strings.Join(consts.CharmList, "','"))
	} else if classifyType == 8 { //同工会主播
		sqlStr = fmt.Sprintf("SELECT rl.room_id FROM room_live rl where exists(select ga.user_id from  guild_anchor ga where ga.user_id = rl.user_id and ga.guild_id = %v) and rl.`status` =1", guildId)
	} else if classifyType == 9 { //游戏房
		sqlStr = "SELECT rl.room_id FROM room_live rl where rl.`status` =1 and rl.game_id>0"
	} else {
		return nil, nil
	}
	result, err := m.DB().Ctx(ctx).Query(ctx, sqlStr)
	if err != nil {
		return nil, err
	}
	var roomIds []string
	for _, res := range result.Array("room_id") {
		roomIds = append(roomIds, res.String())
	}
	return roomIds, nil
}

func (m *roomLiveDao) UpdateRoomType(ctx context.Context, roomId string, roomType, layoutSet int64) {
	updateData := g.Map{"update_time": time.Now().Unix()}
	if roomType != -1 {
		updateData["room_type"] = roomType
	}
	if layoutSet != -1 {
		updateData["layout_set"] = layoutSet
	}
	m.DB().Ctx(ctx).Model().
		Where("room_id = ?", roomId).
		Cache(gdb.CacheOption{
			Duration: -1,
			Name:     consts.MysqlRoomInfoByRoomId + roomId,
		}).Data(updateData).Update()
}

func (m *roomLiveDao) UpdateMaxMultiple(ctx context.Context, roomId string, maxMultiple int64) {
	updateData := g.Map{"update_time": time.Now().Unix()}
	if maxMultiple != -1 {
		updateData["max_multiple"] = maxMultiple
	}
	m.DB().Ctx(ctx).Model().
		Where("room_id = ?", roomId).
		Cache(gdb.CacheOption{
			Duration: -1,
			Name:     consts.MysqlRoomInfoByRoomId + roomId,
		}).Data(updateData).Update()
}

func (m *roomLiveDao) GetLastLiveRoomByUserId(ctx context.Context, userId int64) (model *entity.RoomLive, err error) {
	err = m.DB().Ctx(ctx).Model().
		Where("user_id = ? and status = ?", userId, consts.RoomStatusLiving).
		OrderDesc("last_heartbeat_time").
		Limit(1).
		Scan(&model)
	if err != nil {
		g.Log().Warningf(ctx, "GetLastLiveRoom,UserId = %v, Err = %v", userId, err)
	}
	return
}

func (m *roomLiveDao) GetLiveTimeGtTime(ctx context.Context, nowTime int64) (res []*entity.RoomLive) {
	m.DB().Ctx(ctx).Model().
		Ctx(ctx).
		Where("status = ?", consts.RoomStatusLiving).
		Where("begin_time < ?", nowTime).
		OrderDesc("last_heartbeat_time").
		Scan(&res)
	return res
}
