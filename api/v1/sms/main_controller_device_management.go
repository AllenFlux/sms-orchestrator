package sms

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "sms_backend/api/v1/common"
	"sms_backend/internal/model"
)

type ProjectListReq struct {
	g.Meta `path:"/project/list" tags:"项目管理" method:"get" dc:"查看项目列表" `
	model.PageReq
}
type ProjectListRes struct {
	Data []struct {
		ID                int    `json:"id"`
		Name              string `json:"name" `
		QuantityDevice    int    `json:"quantity_device"`
		AssociatedAccount string `json:"associated_account"`
		Note              string `json:"note"`
		UpdateTime        string `json:"update_time"`
	} `json:"data"`
	commonApi.ListRes
}

type ProjectCreateReq struct {
	g.Meta      `path:"/project" tags:"项目管理" method:"post" dc:"新增项目" `
	ProjectName string `json:"project_name" v:"required"`
	Note        string `json:"note" `
}

type ProjectCreateRes struct {
}

type ProjectUpdateReq struct {
	g.Meta      `path:"/project" tags:"项目管理" method:"put" dc:"编辑项目" `
	ProjectID   int    `json:"project_id" v:"required"`
	ProjectName string `json:"project_name" `
	Note        string `json:"note" `
}

type ProjectUpdateRes struct {
}

type ProjectDeleteReq struct {
	g.Meta    `path:"/project" tags:"项目管理" method:"delete" dc:"删除项目" `
	ProjectID int `json:"project_id" v:"required"`
}
type ProjectDeleteRes struct {
}

type DeviceListReq struct {
	g.Meta `path:"/device/list" tags:"设备列表" method:"get" dc:"查看设备列表" `
	model.PageReq
	DateRange    []string `json:"date_range" p:"dateRange" description:"日期范围"`
	ProjectID    string   `json:"project_id" dc:"查询条件中的项目ID"`
	SentStatus   int      `json:"sent_status" dc:"需要查询的设备状态"`
	TaskName     string   `json:"task_name" dc:"任务名称"`
	Number       string   `json:"number" dc:"设备号码"`
	DeviceNumber string   `json:"device_number" dc:"设备序列号"`
}
type DeviceListRes struct {
	Data struct {
		ID            int    `json:"id"`
		DeviceID      string `json:"device_id"`
		DeviceNumber  string `json:"device_number"`
		Number        string `json:"number"`
		ActiveDays    int    `json:"active_days"`
		OwnerAccount  string `json:"owner_account"`
		AssignedItems string `json:"assigned_items"`
		SentStatus    int    `json:"sent_status" dc:"1 空闲 2 异常 3 占用"`
		QuantitySent  string `json:"quantity_sent"`
		DeviceStatus  int    `json:"device_status" dc:"1 空闲 2 异常 3 占用"`
		ActiveTime    string `json:"active_time"`
	} `json:"data"`
	commonApi.ListRes
}

type AllocateDevice2ProjectReq struct {
	g.Meta       `path:"/device/project" tags:"设备列表" method:"post" dc:"分配设备给项目" `
	DeviceIdList []string `json:"device_id_list"`
}
type AllocateDevice2ProjectRes struct {
}