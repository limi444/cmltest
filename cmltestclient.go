package cmltest

import (
	"git.pinquest.cn/qlb/brick/rpc"
)

var ServiceName = "cmltest"

func AddRole(ctx *rpc.Context, req *AddRoleReq) (*AddRoleRsp, error) {
	rsp := &AddRoleRsp{}
	return rsp, rpc.ClientCall(ctx, ServiceName, AddRoleCMDPath, req, rsp)
}

func UpdateRole(ctx *rpc.Context, req *UpdateRoleReq) (*UpdateRoleRsp, error) {
	rsp := &UpdateRoleRsp{}
	return rsp, rpc.ClientCall(ctx, ServiceName, UpdateRoleCMDPath, req, rsp)
}

func DelRole(ctx *rpc.Context, req *DelRoleReq) (*DelRoleRsp, error) {
	rsp := &DelRoleRsp{}
	return rsp, rpc.ClientCall(ctx, ServiceName, DelRoleCMDPath, req, rsp)
}

func CopyRole(ctx *rpc.Context, req *CopyRoleReq) (*CopyRoleRsp, error) {
	rsp := &CopyRoleRsp{}
	return rsp, rpc.ClientCall(ctx, ServiceName, CopyRoleCMDPath, req, rsp)
}

func SortRole(ctx *rpc.Context, req *SortRoleReq) (*SortRoleRsp, error) {
	rsp := &SortRoleRsp{}
	return rsp, rpc.ClientCall(ctx, ServiceName, SortRoleCMDPath, req, rsp)
}

func GetRoleList(ctx *rpc.Context, req *GetRoleListReq) (*GetRoleListRsp, error) {
	rsp := &GetRoleListRsp{}
	return rsp, rpc.ClientCall(ctx, ServiceName, GetRoleListCMDPath, req, rsp)
}

func SetRoleStaff(ctx *rpc.Context, req *SetRoleStaffReq) (*SetRoleStaffRsp, error) {
	rsp := &SetRoleStaffRsp{}
	return rsp, rpc.ClientCall(ctx, ServiceName, SetRoleStaffCMDPath, req, rsp)
}

func GetDefaultRoleList(ctx *rpc.Context, req *GetDefaultRoleListReq) (*GetDefaultRoleListRsp, error) {
	rsp := &GetDefaultRoleListRsp{}
	return rsp, rpc.ClientCall(ctx, ServiceName, GetDefaultRoleListCMDPath, req, rsp)
}

func GetPermission(ctx *rpc.Context, req *GetPermissionReq) (*GetPermissionRsp, error) {
	rsp := &GetPermissionRsp{}
	return rsp, rpc.ClientCall(ctx, ServiceName, GetPermissionCMDPath, req, rsp)
}

func AddPermission(ctx *rpc.Context, req *AddPermissionReq) (*AddPermissionRsp, error) {
	rsp := &AddPermissionRsp{}
	return rsp, rpc.ClientCall(ctx, ServiceName, AddPermissionCMDPath, req, rsp)
}

func SetPermission(ctx *rpc.Context, req *SetPermissionReq) (*SetPermissionRsp, error) {
	rsp := &SetPermissionRsp{}
	return rsp, rpc.ClientCall(ctx, ServiceName, SetPermissionCMDPath, req, rsp)
}

func GetPermissionList(ctx *rpc.Context, req *GetPermissionListReq) (*GetPermissionListRsp, error) {
	rsp := &GetPermissionListRsp{}
	return rsp, rpc.ClientCall(ctx, ServiceName, GetPermissionListCMDPath, req, rsp)
}

func DelPermission(ctx *rpc.Context, req *DelPermissionReq) (*DelPermissionRsp, error) {
	rsp := &DelPermissionRsp{}
	return rsp, rpc.ClientCall(ctx, ServiceName, DelPermissionCMDPath, req, rsp)
}
