package sms

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	commonApi "sms_backend/api/v1/common"
	"sms_backend/internal/model"
)

type SubTaskListReq struct {
	model.PageReq
	g.Meta    `path:"/sub/task/list" tags:"子平台群发短信" method:"get" dc:"查看群发短信列表" `
	ProjectID int      `json:"project_id" dc:"所属项目id" example:"1"`
	TaskName  string   `json:"task_name" dc:"任务名称"`
	DateRange []string `json:"date_range" p:"dateRange" description:"日期范围"`
}

type SubTaskListRes struct {
	commonApi.ListRes
	Data struct {
		ID                int    `json:"id" dc:"序号"`
		TaskID            int    `json:"task_id" dc:"任务id"`
		TaskName          string `json:"task_name" dc:"任务名称"`
		FileName          string `json:"file_name" dc:"文件名"`
		DeviceQuota       int    `json:"device_quota" dc:"执行设备"`
		TaskStatus        string `json:"task_status" dc:"任务状态"`
		SmsQuantity       string `json:"sms_quantity" dc:"SMS Quantity 短信总条数"`
		SurplusQuantity   string `json:"surplus_quantity" dc:"剩余数量"`
		QuantitySent      string `json:"quantity_sent" dc:"以发送数量"`
		AssociatedAccount string `json:"associated_account" dc:"所属子账号"`
		IntervalTime      string `json:"interval_time" dc:"间隔时间"`
		StartTime         string `json:"start_time" dc:"开始时间"`
		CreateTime        string `json:"create_time" dc:"创建时间"`
	} `json:"data"`
}

// Create Task
type SubTaskCreateReq struct {
	g.Meta          `path:"/sub/task" mine:"multipart/form-data" tags:"子平台群发短信" method:"post" dc:"创建任务" `
	File            *ghttp.UploadFile `json:"file" v:"required"`
	TaskName        string            `json:"task_name" v:"required"`
	GroupID         int               `json:"group_id" dc:"选择分组" v:"required"`
	IntervalTime    string            `json:"interval_time" dc:"间隔时间" v:"required"`
	TimingStartTime string            `json:"timing_start_time"  dc:"定时启动时间"`
}

type SubTaskCreateRes struct{}

// Recall Task

type SubTaskDeleteReq struct {
	g.Meta `path:"/sub/task"  tags:"子平台群发短信" method:"delete" dc:"撤销任务" `
	TaskID int `json:"task_id" v:"required" dc:"任务id"`
}

type SubTaskDeleteRes struct{}

// 群发记录

type SubTaskRecordReq struct {
	model.PageReq
	g.Meta            `path:"/sub/task/record" tags:"群发记录" method:"get" dc:"查看群发短信列表" `
	ProjectID         int      `json:"project_id" dc:"所属项目id" example:"1"`
	SmsStatus         string   `json:"sms_status" dc:"SMS Status 发送状态"`
	TaskName          string   `json:"task_name" dc:"任务名称"`
	TargetPhoneNumber string   `json:"target_phone_number" dc:"Target Phone Number 目标手机号"`
	DeviceNumber      string   `json:"device_number" dc:"Device Number 执行设备号"`
	SentDateRange     []string `json:"sent_date_range" p:"dateRange 1" description:"发送日期范围"`
	CreateDateRange   []string `json:"create_date_range" p:"dateRange 2" description:"创建日期范围"`
}

type SubTaskRecordRes struct {
	commonApi.ListRes
	Data struct {
		ID                int    `json:"id" dc:"序号"`
		TaskName          string `json:"task_name" dc:"任务名称"`
		SubTaskID         string `json:"sub_task_id" dc:"Sub Task ID 子任务ID"`
		TargetPhoneNumber string `json:"target_phone_number" dc:"Target Phone Number 目标手机号"`
		DeviceNumber      string `json:"device_number" dc:"Device Number 执行设备号"`
		SmsTopic          string `json:"sms_topic" dc:"SMS Topic 主题"`
		SmsContent        string `json:"sms_content" dc:"SMS Content 短信内容"`
		SmsStatus         string `json:"sms_status" dc:"SMS Status 短信发送状态"`
		AssociatedAccount string `json:"associated_account" dc:"所属子账号"`
		ProjectName       string `json:"project_name" dc:"Project Name 所属项目"`
		Note              string `json:"note" dc:"失败原因"`
		StartTime         string `json:"start_time" dc:"开始时间"`
		CreateTime        string `json:"create_time" dc:"创建时间"`
	} `json:"data"`
}
