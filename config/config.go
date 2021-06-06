package config

import (
	"github.com/lecex/core/config"
	"github.com/lecex/core/env"
	PB "github.com/lecex/user/proto/permission"
)

// 	Conf 配置
// 	Name // 服务名称
//	Method // 方法
//	Auth // 是否认证授权
//	Policy // 是否认证权限
//	Name // 权限名称
//	Description // 权限解释
var Conf config.Config = config.Config{
	Name:    env.Getenv("MICRO_API_NAMESPACE", "go.micro.api.") + "device-api",
	Version: "v1.0.3",
	Service: map[string]string{
		"device": env.Getenv("DEVICE_SERVICE", "go.micro.srv.device"),
		"user":   env.Getenv("USER_SERVICE", "go.micro.srv.user"),
	},
	Permissions: []*PB.Permission{ // 权限管理
		// 设备管理
		{Service: "device-api", Method: "device.Create", Auth: true, Policy: true, Name: "创建设备", Description: "创建新设备权限。"},
		{Service: "device-api", Method: "device.Delete", Auth: true, Policy: true, Name: "删除设备", Description: "删除设备。"},
		{Service: "device-api", Method: "device.Update", Auth: true, Policy: true, Name: "更新设备", Description: "更新设备信息。"},
		{Service: "device-api", Method: "device.Get", Auth: true, Policy: true, Name: "查询设备", Description: "查询设备信息权限。"},
		{Service: "device-api", Method: "device.List", Auth: true, Policy: true, Name: "设备列表", Description: "查询设备列表"},
	},
}
