// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sms_backend/api/v1/system"
	"sms_backend/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IOperateLog interface {
		// OperationLog 操作日志写入
		OperationLog(r *ghttp.Request)
		Invoke(ctx context.Context, data *model.SysOperLogAdd)
		List(ctx context.Context, req *system.SysOperLogSearchReq) (listRes *system.SysOperLogSearchRes, err error)
		GetByOperId(ctx context.Context, operId int64) (res *model.SysOperLogInfoRes, err error)
		DeleteByIds(ctx context.Context, ids []int64) (err error)
		ClearLog(ctx context.Context) (err error)
	}
)

var (
	localOperateLog IOperateLog
)

func OperateLog() IOperateLog {
	if localOperateLog == nil {
		panic("implement not found for interface IOperateLog, forgot register?")
	}
	return localOperateLog
}

func RegisterOperateLog(i IOperateLog) {
	localOperateLog = i
}
