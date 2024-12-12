package role

import (
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/api/v1/allUser"
	"sms_backend/internal/model/entity"
)

type ListReq struct {
	g.Meta `path:"/list" tags:"角色管理" method:"post" sm:"获取角色列表"`
	allUser.GeneralReq
}

type ListRes struct {
	g.Meta `mime:"application/json"`
	allUser.GeneralRes
	List []*entity.Role
}

type CreatedReq struct {
	g.Meta     `path:"/create" tags:"角色管理" method:"post" sm:"创建角色"`
	Name       string `json:"name" v:"required" dc:"角色名称"`
	Permission []int  `json:"permission" v:"required" dc:"权限id,空数组会返回错误"`
	Note       string `json:"note" dc:"备注"`
}

type CreatedRes struct{}

type UpdatedReq struct {
	g.Meta     `path:"/update" tags:"角色管理" method:"put" sm:"修改角色"`
	ID         int    `json:"id" v:"required" dc:"角色id"`
	Name       string `json:"name" v:"required" dc:"角色名称"`
	Permission []int  `json:"permission" v:"required" dc:"权限id,空数组会返回错误"`
	Note       string `json:"note" dc:"备注"`
}

type UpdatedRes struct{}

type DeletedReq struct {
	g.Meta `path:"/delete" tags:"角色管理" method:"delete" sm:"删除角色"`
	ID     int `json:"id" v:"required" dc:"角色id"`
}

type DeletedRes struct{}