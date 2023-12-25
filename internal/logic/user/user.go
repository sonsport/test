package user

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"

	"fuya-ark/internal/consts"
	"fuya-ark/internal/dao"
	"fuya-ark/internal/model"
	"fuya-ark/internal/model/do"
	"fuya-ark/internal/service"
	"fuya-ark/utility"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// Register 注册
func (s *sUser) Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error) {
	//处理加密盐和密码的逻辑
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	//插入数据返回id
	lastInsertID, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RegisterOutput{Id: uint(lastInsertID)}, err
}

// UpdatePassword 修改密码
func (*sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) (out model.UpdatePasswordOutput, err error) {
	//	验证密保问题
	//userInfo := do.UserInfo{}
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	//err = dao.UserInfo.Ctx(ctx).WherePri(userId).Scan(&userInfo)
	//if err != nil {
	//	return model.UpdatePasswordOutput{}, err
	//}
	//if gconv.String(userInfo.SecretAnswer) != in.SecretAnswer {
	//	return out, errors.New(consts.ErrSecretAnswerMsg)
	//}
	//userSalt := grand.S(10)
	//in.UserSalt = userSalt
	//in.Password = utility.EncryptPassword(in.Password, userSalt)
	//_, err = dao.UserInfo.Ctx(ctx).WherePri(userId).Update(in)
	//if err != nil {
	//	return model.UpdatePasswordOutput{}, err
	//}
	return model.UpdatePasswordOutput{Id: userId}, nil
}

// UserDetail 获取用户详情
func (*sUser) UserDetail(ctx context.Context, in model.UserDetailInput) (out model.UserDetailOutput, err error) {
	userInfo := do.UserInfo{}
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.UserInfo.Ctx(ctx).WherePri(dao.UserInfo.Columns().UserId, userId).Scan(&userInfo)
	if err != nil {
		return model.UserDetailOutput{}, err
	}
	err = gconv.Struct(userInfo, &out)
	return out, err
}
