package video

import (
	"net/http"

	"douyin/apps/api/internal/logic/video"
	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavorlistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VideoFavorListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewFavorlistLogic(r.Context(), svcCtx)
		resp, err := l.Favorlist(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
