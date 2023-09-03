package logic

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"douyin/apps/rpc/chat/internal/svc"
	"douyin/apps/rpc/chat/types/chat"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLogic {
	return &MessageChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageChatLogic) MessageChat(in *chat.MessageChatReq) (*chat.MessageChatResp, error) {
	// todo: add your logic here and delete this line

	// 缓存连接
	// redisDb, err := pkg.RedisInit()
	// if err != nil {
	// 	return nil, err
	// }
	redisDb := l.svcCtx.DbRedis

	// 整理缓存的数据
	var messagelist []*chat.MessageList
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
	if len(messagelist1) == 0 && len(messagelist2) == 0 {
		return &chat.MessageChatResp{MessageList: messagelist}, nil
	} else if len(messagelist1) == 0 && len(messagelist2) != 0 {
		for _, v := range messagelist2 {
			var messagedata pkg.Chat
			err = json.Unmarshal([]byte(v), &messagedata)
			if err != nil {
				return nil, err
			}
			// 通过时间判断该消息是否已经被发送
			if messagedata.CreatedAt.Unix() > in.PreMsgTime {
				messagelist = append(messagelist, &chat.MessageList{
					Id:         messagedata.ID,
					ToUserId:   messagedata.ToUserId,
					FromUserId: messagedata.FromUserId,
					Content:    messagedata.Content,
					CreateTime: messagedata.CreatedAt.Unix(),
				})
			} else {
				break
			}
		}
	} else if len(messagelist1) != 0 && len(messagelist2) == 0 {
		for _, v := range messagelist1 {
			var messagedata pkg.Chat
			err = json.Unmarshal([]byte(v), &messagedata)
			if err != nil {
				return nil, err
			}
			// 通过时间判断该消息是否已经被发送
			if messagedata.CreatedAt.Unix() > in.PreMsgTime {
				messagelist = append(messagelist, &chat.MessageList{
					Id:         messagedata.ID,
					ToUserId:   messagedata.ToUserId,
					FromUserId: messagedata.FromUserId,
					Content:    messagedata.Content,
					CreateTime: messagedata.CreatedAt.Unix(),
				})
			} else {
				break
			}
		}
	}

	return &chat.MessageChatResp{MessageList: messagelist}, nil
}
