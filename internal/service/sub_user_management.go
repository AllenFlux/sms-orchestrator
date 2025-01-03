// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sms_backend/api/v1/subUser"
)

type (
	ISubUser interface {
		GetList(ctx context.Context, req *subUser.SubGetListReq) (*subUser.SubGetListRes, error)
		CreatedSubUser(ctx context.Context, req *subUser.SubRegisterReq) (res *subUser.SubRegisterRes, err error)
		UpdateSubUser(ctx context.Context, req *subUser.SubUpdateReq) (res *subUser.SubUpdateRes, err error)
		ChangeStatus(ctx context.Context, req *subUser.SubChangeStatusReq) (res *subUser.SubChangeStatusRes, err error)
		DeleteSubUser(ctx context.Context, req *subUser.SubDeleteReq) (res *subUser.SubDeleteRes, err error)
	}
)

var (
	localSubUser ISubUser
)

func SubUser() ISubUser {
	if localSubUser == nil {
		panic("implement not found for interface ISubUser, forgot register?")
	}
	return localSubUser
}

func RegisterSubUser(i ISubUser) {
	localSubUser = i
}
