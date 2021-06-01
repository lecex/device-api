package handler

import (
	"context"

	client "github.com/lecex/core/client"

	pb "github.com/lecex/device-api/proto/device"
)

// Device 设备结构
type Device struct {
	ServiceName string
}

// All 权限列表
func (srv *Device) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Devices.All", req, res)
}

// List 设备列表
func (srv *Device) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Devices.List", req, res)
}

// Get 获取设备
func (srv *Device) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Devices.Get", req, res)
}

// Create 创建设备
func (srv *Device) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Devices.Create", req, res)
}

// Update 更新设备
func (srv *Device) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Devices.Update", req, res)
}

// Delete 删除设备
func (srv *Device) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Devices.Delete", req, res)
}
