package bo

import (
	"fuya-ark/internal/model"
)

type RobotJoinObject struct {
	UserId   int64             `json:"user_id"`  //用户id
	RoomId   string            `json:"room_id"`  //房间id
	Level    int               `json:"level"`    //用户等级
	Name     string            `json:"name"`     //用户昵称
	Portrait string            `json:"portrait"` //用户头像
	MsgType  int               `json:"msg_type"` //类型。1加，2减
	Robot    *model.RoomViewer `json:"robot"`    //机器人信息
}
