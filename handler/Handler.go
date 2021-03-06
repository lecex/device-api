package handler

import (
	"context"
	"time"

	server "github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/util/log"

	client "github.com/lecex/core/client"

	"github.com/lecex/device-api/config"
	cashierPB "github.com/lecex/device-api/proto/cashier"
	devicePB "github.com/lecex/device-api/proto/device"

	PB "github.com/lecex/user/proto/permission"
)

var Conf = config.Conf

// Register 注册
func Register(Server server.Server) {
	cashierPB.RegisterCashiersHandler(Server, &Cashier{Conf.Service["device"]})
	devicePB.RegisterDevicesHandler(Server, &Device{Conf.Service["device"]})
	go Sync() // 同步前端权限
}

// Sync 同步
func Sync() {
	time.Sleep(5 * time.Second)
	req := &PB.Request{
		Permissions: Conf.Permissions,
	}
	res := &PB.Response{}
	err := client.Call(context.TODO(), Conf.Service["user"], "Permissions.Sync", req, res)
	if err != nil {
		log.Log(err)
		Sync()
	}
}
