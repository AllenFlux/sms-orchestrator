// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EntDeviceLabel is the golang structure for table ent_device_label.
type EntDeviceLabel struct {
	Id        uint        `json:"id"         orm:"id"         description:""`      //
	Name      string      `json:"name"       orm:"name"       description:"标签名称"`  // 标签名称
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`      //
	CreatedBy int         `json:"created_by" orm:"created_by" description:"创建者id"` // 创建者id
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:"删除时间"`  // 删除时间
}
