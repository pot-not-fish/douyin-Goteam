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

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line

	// 数据库连接
	// db, err := userdb.MysqlInit()
	// if err != nil {
	// 	return nil, err
	// }
	db := l.svcCtx.DbEngin

	// 判断用户名是否重复
	info := userdb.User{}
	err := db.Where("username = ?", in.Username).First(&info).Error
	if err != nil {

		// 创建用户
		info = userdb.User{Username: in.Username, Password: fmt.Sprintf("%x", md5.Sum([]byte(in.Password)))}
		err = db.Create(&info).Error
		if err != nil {
			return nil, errors.New("创建用户失败")
		}

		// 返回用户id
		return &user.RegisterResp{UserId: info.ID}, nil
	} else {
		return nil, errors.New("用户名重复")
	}
}
