package user

import (
	"context"
	"errors"

	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"douyin/apps/rpc/user/types/user"
	authcrypto "douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	// 预处理username 去除空格
	filter := []byte(req.Username)
	for i := 0; i < len(filter); i++ {
		if filter[i] == 32 {
			filter = append(filter[:i], filter[i+1:]...)
		}
	}

	// 判断是否用户名和密码长度是否超过32
	if len(req.Username) > 32 || len(req.Password) > 32 {
		return &types.LoginResp{Status_code: 1, Status_msg: "用户名或密码超过32位", User_id: -1, Token: "null"}, errors.New("用户名或密码超过32位")
	}

	// 发送给rpc的login处理
	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{Username: req.Username, Password: req.Password})
	if err != nil {
		return &types.LoginResp{Status_code: 1, Status_msg: err.Error(), User_id: -1, Token: "null"}, nil
	} else {
		token, err := authcrypto.GetAuthToken(loginResp.UserId)
		if err != nil {
			return &types.LoginResp{Status_code: 1, Status_msg: err.Error(), User_id: -1, Token: "null"}, nil
		}
		return &types.LoginResp{Status_code: 0, Status_msg: "登录成功", User_id: loginResp.UserId, Token: token}, nil
	}
}
