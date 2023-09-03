package logic

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"

	"douyin/apps/rpc/user/internal/svc"
	"douyin/apps/rpc/user/types/user"
	userdb "douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// todo: add your logic here and delete this line
	// 数据库连接
	// db, err := userdb.MysqlInit()
	// if err != nil {
	// 	return nil, err
	// }
	db := l.svcCtx.DbEngin

	// 判断用户名和密码是否正确
	info := userdb.User{}
	err := db.Where("username = ? AND password = ?", in.Username, fmt.Sprintf("%x", md5.Sum([]byte(in.Password)))).First(&info).Error
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	} else {
		return &user.LoginResp{UserId: info.ID}, nil
	}
}
