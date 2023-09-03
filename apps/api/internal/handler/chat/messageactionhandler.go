package chat

import (
	"net/http"

	"douyin/apps/api/internal/logic/chat"
	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MessageactionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageActionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := chat.NewMessageactionLogic(r.Context(), svcCtx)
		resp, err := l.Messageaction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
