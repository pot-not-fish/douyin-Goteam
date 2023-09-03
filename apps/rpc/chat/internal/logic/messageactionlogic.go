package logic

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"douyin/apps/rpc/chat/internal/svc"
	"douyin/apps/rpc/chat/types/chat"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageActionLogic {
	return &MessageActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageActionLogic) MessageAction(in *chat.MessageActionReq) (*chat.MessageActionResp, error) {
	// todo: add your logic here and delete this line

	// 缓存连接
	// redisDb, err := pkg.RedisInit()
	// if err != nil {
	// 	return nil, err
	// }
	redisDb := l.svcCtx.DbRedis

	// 缓存创建
	messagechat := pkg.Chat{
		Content:    in.Content,
		ToUserId:   in.ToUserId,
		FromUserId: in.UserId,
	}
	userid_str := strconv.FormatInt(in.UserId, 10)
	to_userid_str := strconv.FormatInt(in.ToUserId, 10)
	messagelist1, err := redisDb.LRange(userid_str+"_"+to_userid_str, 0, -1).Result()
	if err != nil {
		return nil, errors.New("failed to find userid touserid redis")
	}
	messagelist2, err := redisDb.LRange(to_userid_str+"_"+userid_str, 0, -1).Result()
	if err != nil {
		return nil, errors.New("failed to find touserid userid redis")
	}
	if (len(messagelist1) == 0 && len(messagelist2) == 0) || (len(messagelist1) != 0 && len(messagelist2) == 0) {
		messagechat.ID, err = redisDb.LLen(userid_str + "_" + to_userid_str).Result()
		if err != nil {
			return nil, err
		}
		messagechat.CreatedAt = time.Now()
		messagedata, err := json.Marshal(messagechat)
		if err != nil {
			return nil, err
		}
		redisDb.RPush(userid_str+"_"+to_userid_str, string(messagedata))
	} else if len(messagelist1) == 0 && len(messagelist2) != 0 {
		messagechat.ID, err = redisDb.LLen(to_userid_str + "_" + userid_str).Result()
		if err != nil {
			return nil, err
		}
		messagechat.CreatedAt = time.Now()
		messagedata, err := json.Marshal(messagechat)
		if err != nil {
			return nil, err
		}
		redisDb.RPush(to_userid_str+"_"+userid_str, string(messagedata))
	}

	return &chat.MessageActionResp{StatusMsg: "消息发布成功"}, nil
}
