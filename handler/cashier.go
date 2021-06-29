package handler

import (
	"context"

	client "github.com/lecex/core/client"

	pb "github.com/lecex/cashier-api/proto/cashier"
)

// Cashier 设备结构
type Cashier struct {
	ServiceName string
}

// All 权限列表
func (srv *Cashier) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Cashiers.All", req, res)
}

// List 设备列表
func (srv *Cashier) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Cashiers.List", req, res)
}

// Get 获取设备
func (srv *Cashier) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Cashiers.Get", req, res)
}

// Create 创建设备
func (srv *Cashier) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Cashiers.Create", req, res)
}

// Update 更新设备
func (srv *Cashier) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Cashiers.Update", req, res)
}

// Delete 删除设备
func (srv *Cashier) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Cashiers.Delete", req, res)
}
