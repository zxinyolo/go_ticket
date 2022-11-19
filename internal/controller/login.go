package controller

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	v1 "go_ticket/api/v1"
	"go_ticket/internal/errorcode"
	"go_ticket/internal/model"
	"go_ticket/internal/service"
	"go_ticket/utility/token"
	"go_ticket/utility/utils"
	"strconv"
)

var Login = cLogin{}

type cLogin struct {
}

func (l *cLogin) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = &v1.LoginRes{}
	user, err := service.User().Login(ctx, model.UserLoginInput{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, errorcode.MyWrapCode(ctx, errorcode.LoginFailed, err)
	}
	res.User = &v1.UserGetInfoRes{}
	err = utils.MyCopy(ctx, res.User, user.User)
	if err != nil {
		return nil, err
	}
	userKey := strconv.Itoa(user.User.Id)
	MyToken, err := token.Instance().GenerateToken(ctx, userKey, user.User)
	res.Token = MyToken.Token
	glog.Info(ctx, "登录用户信息xxxx：", res)
	return

}
