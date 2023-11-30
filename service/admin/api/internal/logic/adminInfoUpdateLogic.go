package logic

import (
	"context"

	"go-zero-mall/service/admin/api/internal/svc"
	"go-zero-mall/service/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminInfoUpdateLogic {
	return &AdminInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminInfoUpdateLogic) AdminInfoUpdate(req *types.AdminInfoUpdateReq) (resp *types.AdminInfoUpdateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
