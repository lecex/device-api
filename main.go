package main

import (
	// 公共引入
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	_ "github.com/lecex/core/plugins"
	"github.com/lecex/device-api/config"
	"github.com/lecex/device-api/handler"
)

func main() {
	var Conf = config.Conf
	service := micro.NewService(
		micro.Name(Conf.Name),
		micro.Version(Conf.Version),
		micro.WrapHandler(Conf.Middleware().Wrapper), //验证权限
	)
	service.Init()

	// 注册服务
	handler.Register(service.Server())
	// Run the server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
