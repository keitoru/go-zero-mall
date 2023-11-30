package logic

import (
	"context"

	"go-zero-mall/service/admin/api/internal/svc"
	"go-zero-mall/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminInfoLogic {
	return &AdminInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminInfoLogic) AdminInfo(req *types.AdminInfoReq) (resp *types.AdminInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
