package relation

import (
	"context"
	"strconv"

	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"douyin/apps/rpc/relation/types/relation"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationactionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationactionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationactionLogic {
	return &RelationactionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationactionLogic) Relationaction(req *types.RelationActionReq) (resp *types.RelationActionResp, err error) {
	// todo: add your logic here and delete this line

	// token鉴权
	userid, err := pkg.AuthToken(req.Token)
	if err != nil {
		return &types.RelationActionResp{
			Status_code: 1,
			Status_msg:  err.Error(),
		}, nil
	}

	// 不能自己关注自己
	touserid, _ := strconv.ParseInt(req.To_user_id, 10, 64)
	actiontype, _ := strconv.ParseInt(req.Action_type, 10, 64)
	if userid == touserid {
		return &types.RelationActionResp{
			Status_code: 1,
			Status_msg:  "不能关注自己",
		}, nil
	}

	// 发送给relationaction的rpc处理
	relationactionResp, err := l.svcCtx.RelationRpc.RelationAction(l.ctx, &relation.RelationActionReq{
		UserId:     userid,
		ToUserId:   touserid,
		ActionType: int32(actiontype),
	})
	if err != nil {
		return &types.RelationActionResp{
			Status_code: 1,
			Status_msg:  err.Error(),
		}, nil
	} else {
		return &types.RelationActionResp{
			Status_code: 0,
			Status_msg:  relationactionResp.StatusMsg,
		}, nil
	}
}
