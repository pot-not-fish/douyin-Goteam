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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
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
		return &types.RegisterResp{Status_code: 1, Status_msg: "用户名或密码超过32位", User_id: -1, Token: "null"}, errors.New("用户名或密码超过32位")
	}

	// 发送给rpc的register处理
	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{Username: string(filter), Password: req.Password})
	if err != nil {
		return &types.RegisterResp{Status_code: 1, Status_msg: err.Error(), User_id: -1, Token: "null"}, nil
	} else {
		token, err := authcrypto.GetAuthToken(registerResp.UserId)
		if err != nil {
			return &types.RegisterResp{Status_code: 1, Status_msg: err.Error(), User_id: -1, Token: "null"}, nil
		}
		return &types.RegisterResp{Status_code: 0, Status_msg: "注册成功", User_id: registerResp.UserId, Token: token}, nil
	}
}
