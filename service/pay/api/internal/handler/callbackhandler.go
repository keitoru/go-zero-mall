package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-mall/service/pay/api/internal/logic"
	"go-zero-mall/service/pay/api/internal/svc"
	"go-zero-mall/service/pay/api/internal/types"
)

func CallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CallbackRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCallbackLogic(r.Context(), svcCtx)
		resp, err := l.Callback(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
