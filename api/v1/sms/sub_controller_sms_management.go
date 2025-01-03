package sms

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	commonApi "sms_backend/api/v1/common"
	"sms_backend/internal/model"
)

type SubTaskListReq struct {
	model.PageReq
	g.Meta    `path:"/sub/task/list" tags:"子平台群发短信" method:"get" dc:"查看群发短信列表" `
	SubUserID int      `json:"sub_user_id" v:"required" dc:"子平台账号id"`
	ProjectID int      `json:"project_id" dc:"所属项目id" example:"1"`
	TaskName  string   `json:"task_name" dc:"任务名称"`
	DateRange []string `json:"date_range" p:"dateRange" description:"日期范围"`
}

type SubTaskListResData struct {
	ID                int    `json:"id" dc:"序号"`
	ProjectID         int    `json:"project_id" dc:"项目id"`
	ProjectName       string `json:"project_name" dc:"项目name"`
	TaskName          string `json:"task_name" dc:"任务名称"`
	FileName          string `json:"file_name" dc:"文件名"`
	DeviceQuota       string `json:"device_quota" dc:"执行设备"`
	TaskStatus        int    `json:"task_status" dc:"任务状态"`
	SmsQuantity       int    `json:"sms_quantity" dc:"SMS Quantity 短信总条数"`
	SurplusQuantity   int    `json:"surplus_quantity" dc:"剩余数量"`
	QuantitySent      int    `json:"quantity_sent" dc:"以发送数量"`
	AssociatedAccount string `json:"associated_account" dc:"所属子账号"`
	IntervalTime      string `json:"interval_time" dc:"间隔时间"`
	StartTime         string `json:"start_time" dc:"开始时间"`
	CreateTime        string `json:"create_time" dc:"创建时间"`
}

type SubTaskListRes struct {
	commonApi.ListRes
	Data []SubTaskListResData `json:"data"`
}

// Create Task
type SubTaskCreateReq struct {
	g.Meta          `path:"/sub/task" mine:"multipart/form-data" tags:"子平台群发短信" method:"post" dc:"创建任务" `
	File            *ghttp.UploadFile `json:"file" v:"required"`
	SubUserId       int               `json:"sub_user_id" v:"required"`
	TaskName        string            `json:"task_name" v:"required"`
	GroupID         int               `json:"group_id" dc:"选择分组" v:"required"`
	IntervalTime    string            `json:"interval_time" dc:"间隔时间" v:"required"`
	TimingStartTime *gtime.Time       `json:"timing_start_time"  dc:"定时启动时间"`
}

type SubTaskCreateRes struct {
	ID int64 `json:"id"`
}

// Task File Download
type TaskFileDownloadReq struct {
	g.Meta   `path:"/task/file"  tags:"子平台群发短信" method:"post" dc:"任务文件下载" `
	FileName string `json:"file_name" v:"required"`
}

type TaskFileDownloadRes struct {
	JsonData string `json:"json_data" description:"文件内容"`
}

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
	SubUserID         int      `json:"sub_user_id" v:"required"`
	ProjectID         int      `json:"project_id" dc:"所属项目id" example:"1"`
	SmsStatus         int      `json:"sms_status" dc:"SMS Status 发送状态 1发送成功 2发送失败" `
	TaskName          string   `json:"task_name" dc:"任务名称"`
	TargetPhoneNumber string   `json:"target_phone_number" dc:"Target Phone Number 目标手机号"`
	DeviceNumber      string   `json:"device_number" dc:"Device Number 执行设备号"`
	SentDateRange     []string `json:"sent_date_range" p:"dateRange 1" description:"发送日期范围"`
	CreateDateRange   []string `json:"create_date_range" p:"dateRange 2" description:"创建日期范围"`
}

type SubTaskRecordResData struct {
	ID                int    `json:"id" dc:"序号"`
	TaskName          string `json:"task_name" dc:"任务名称"`
	SubTaskID         int    `json:"sub_task_id" dc:"Sub Task ID 子任务ID"`
	TargetPhoneNumber string `json:"target_phone_number" dc:"Target Phone Number 目标手机号"`
	DeviceNumber      string `json:"device_number" dc:"Device Number 执行设备号"`
	SmsTopic          string `json:"sms_topic" dc:"SMS Topic 主题"`
	SmsContent        string `json:"sms_content" dc:"SMS Content 短信内容"`
	SmsStatus         string `json:"sms_status" dc:"SMS Status 短信发送状态"`
	AssociatedAccount string `json:"associated_account" dc:"所属子账号"`
	Reason            string `json:"reason"`
	ProjectName       string `json:"project_name" dc:"Project Name 所属项目"`
	StartTime         string `json:"start_time" dc:"开始时间"`
	CreateTime        string `json:"create_time" dc:"创建时间"`
}

type SubTaskRecordRes struct {
	commonApi.ListRes
	Data []SubTaskRecordResData `json:"data"`
}

type SubGetConversationRecordReq struct {
	g.Meta `path:"/sub/conversation/record" tags:"子平台消息对话" method:"get" dc:"查看消息对话内容" `
	//SubUserID int `json:"sub_user_id" v:"required"`
	ChatLogID int `json:"chat_log_id" v:"required" dc:"需要查看的chart id"`
}

type SubGetConversationRecordRes struct {
	Data []SubGetConversationRecordListResData `json:"data"`
}

type SubGetConversationRecordListReq struct {
	g.Meta `path:"/sub/conversation/record/list" tags:"子平台消息对话" method:"get" dc:"查看消息对话列表" `
	model.PageReq
	SubUserID int `json:"sub_user_id" v:"required"`
	ProjectID int `json:"project_id" v:"required"`
}
type SubGetConversationRecordListResData struct {
	TargetPhoneNumber string      `json:"target_phone_number"`
	Content           string      `json:"content"`
	SentOrReceive     int         `json:"sent_or_receive" dc:"1表示此条短信是发送 2表示此条短信是接收"`
	RecordTime        *gtime.Time `json:"record_time" dc:"接收到信息的时间"`
	ChatLogID         int         `json:"chat_log_id"  dc:"chart id"`
	DeviceNumber      string      `json:"device_number" dc:"Device Number"`
}
type SubGetConversationRecordListRes struct {
	commonApi.ListRes
	Data []SubGetConversationRecordListResData `json:"data"`
}

type SubPostConversationRecordReq struct {
	g.Meta `path:"/sub/conversation/record" tags:"子平台消息对话" method:"post" dc:"在对话中发送消息" `
	//SubUserID  int    `json:"sub_user_id" v:"required"`
	ChatLogID int64  `json:"chat_log_id" v:"required" dc:"需要回复的chat id"`
	Content   string `json:"content" v:"required"`
}

type SubPostConversationRecordRes struct {
	UnSendTaskNum int64  `json:"un_send_task_num" dc:"此种任务类型 当前设备未发送数量"`
	TaskItemName  string `json:"task_item_name"`
}

type SubPostConversationRecordData struct {
	TaskID            int64       `json:"task_id"`
	Content           string      `json:"content"`
	DeviceNumber      string      `json:"device_number"`
	TargetPhoneNumber string      `json:"target_phone_number"`
	Interval          string      `json:"interval"`
	StartAt           *gtime.Time `json:"start_at"`
}
