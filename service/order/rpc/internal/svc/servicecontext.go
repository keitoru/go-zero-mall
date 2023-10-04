package svc

import (
	"go-zero-mall/service/order/model"
	"go-zero-mall/service/order/rpc/internal/config"
	"go-zero-mall/service/product/rpc/product"
	"go-zero-mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	OrderModel model.OrderModel

	UserRpc    user.User
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(conn, c.CacheRedis),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
