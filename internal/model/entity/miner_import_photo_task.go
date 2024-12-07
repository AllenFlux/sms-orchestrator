// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MinerImportPhotoTask is the golang structure for table miner_import_photo_task.
type MinerImportPhotoTask struct {
	Id        int         `json:"id"          orm:"id"          description:""`                                  //
	Name      string      `json:"name"        orm:"name"        description:"任务名称"`                              // 任务名称
	PhotoNum  int         `json:"photo_num"   orm:"photo_num"   description:"图片数量"`                              // 图片数量
	Status    int         `json:"status"      orm:"status"      description:"任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成"` // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	CreatedAt *gtime.Time `json:"created_at"  orm:"created_at"  description:""`                                  //
	ExecAt    *gtime.Time `json:"exec_at"     orm:"exec_at"     description:"执行时间"`                              // 执行时间
	Remark    string      `json:"remark"      orm:"remark"      description:"备注"`                                // 备注
	ClientId  int         `json:"client_id"   orm:"client_id"   description:"客户ID"`                              // 客户ID
	ExecEndAt *gtime.Time `json:"exec_end_at" orm:"exec_end_at" description:"执行结束时间"`                            // 执行结束时间
	DeviceNum int         `json:"device_num"  orm:"device_num"  description:"设备数量"`                              // 设备数量
	HasShow   int         `json:"has_show"    orm:"has_show"    description:"1-显示 2-不显示"`                        // 1-显示 2-不显示
	Type      int         `json:"type"        orm:"type"        description:"1-图片 2-视频"`                         // 1-图片 2-视频
}
