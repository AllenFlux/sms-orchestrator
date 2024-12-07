package sms

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "sms_backend/api/v1/common"
	"sms_backend/internal/model"
)

type TaskListReq struct {
	model.PageReq
	g.Meta    `path:"/task/list" tags:"群发短信" method:"get" dc:"查看群发短信列表" `
	ProjectID int      `json:"project_id" dc:"所属项目id" example:"1"`
	TaskName  string   `json:"task_name" dc:"任务名称"`
	DateRange []string `json:"date_range" p:"dateRange" description:"日期范围"`
}

type TaskListRes struct {
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

type TaskRecordReq struct {
	model.PageReq
	g.Meta            `path:"/task/record" tags:"群发记录" method:"get" dc:"查看群发短信列表" `
	ProjectID         int      `json:"project_id" dc:"所属项目id" example:"1"`
	SmsStatus         string   `json:"sms_status" dc:"SMS Status 发送状态"`
	TaskName          string   `json:"task_name" dc:"任务名称"`
	TargetPhoneNumber string   `json:"target_phone_number" dc:"Target Phone Number 目标手机号"`
	DeviceNumber      string   `json:"device_number" dc:"Device Number 执行设备号"`
	SentDateRange     []string `json:"sent_date_range" p:"dateRange 1" description:"发送日期范围"`
	CreateDateRange   []string `json:"create_date_range" p:"dateRange 2" description:"创建日期范围"`
}

type TaskRecordRes struct {
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
		StartTime         string `json:"start_time" dc:"开始时间"`
		CreateTime        string `json:"create_time" dc:"创建时间"`
	} `json:"data"`
}

type ConversationListReq struct {
	g.Meta `path:"/conversation/list" tags:"消息对话" method:"get" dc:"查看群发短信列表" `
	model.PageReq
	ProjectID          int    `json:"project_id" v:"required"`
	SearchWord         string `json:"search_word" dc:"搜索对话"`
	ConversationStatus int    `json:"conversation_status" dc:"对话状态 1-所有对话 2-正在对话" default:"1"`
}

type ConversationListRes struct {
	commonApi.ListRes
	Data struct {
		ConversationID      int    `json:"conversation_id"`
		ConversationNumber  string `json:"conversation_number" dc:"目标电话号码"`
		SmsTime             string `json:"sms_time"`
		ConversationContent string `json:"conversation_content" dc:"消息内容"`
	} `json:"data"`
}

type ConversationRecordReq struct {
	model.PageReq
	g.Meta         `path:"/conversation/record" tags:"消息对话" method:"get" dc:"单点对话记录" `
	ConversationID int `json:"conversation_id"`
}

type ConversationRecord struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Time     string `json:"time"`
	Content  string `json:"content"`
}
type ConversationRecordRes struct {
	commonApi.ListRes
	ReceiverNumber string `json:"receiver_number" dc:"Receiver 字段代表对方 人类"`
	Data           struct {
		History []ConversationRecord `json:"history"`
	} `json:"data"`
}