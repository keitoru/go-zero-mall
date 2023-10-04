package svc

import (
	"go-zero-mall/service/user/api/internal/config"
	"go-zero-mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
