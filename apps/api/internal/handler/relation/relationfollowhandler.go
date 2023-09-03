package relation

import (
	"net/http"

	"douyin/apps/api/internal/logic/relation"
	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RelationfollowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RelationFollowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := relation.NewRelationfollowLogic(r.Context(), svcCtx)
		resp, err := l.Relationfollow(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
