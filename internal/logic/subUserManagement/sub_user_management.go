package subUserManagement

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"sms_backend/api/v1/sms"
	"sms_backend/api/v1/subUser"
	"sms_backend/internal/consts"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/do"
	"sms_backend/internal/service"
	"sms_backend/library/libUtils"
	"sms_backend/library/liberr"
	"sms_backend/utility"
)

type sSubUser struct{}

func New() *sSubUser { return &sSubUser{} }

func init() { service.RegisterSubUser(New()) }

func (s *sSubUser) GetList(ctx context.Context, req *subUser.SubGetListReq) (*subUser.SubGetListRes, error) {
	// 直接获取请求中的子请求
	userReq := req.GetListReq

	// 调用用户服务获取列表
	userRes, err := service.User().GetList(ctx, userReq)
	if err != nil {
		return nil, err // 提前返回错误，简化逻辑
	}

	// 构造响应并返回
	return &subUser.SubGetListRes{
		GetListRes: userRes,
	}, nil
}

func (s *sSubUser) CreatedSubUser(ctx context.Context, req *subUser.SubRegisterReq) (res *subUser.SubRegisterRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		g.Dump(req)
		orm := dao.User.Ctx(ctx)
		maxSystemId, err := orm.Fields("MAX(system_id)").Value()
		liberr.ErrIsNil(ctx, err)

		exist, err := orm.Where(dao.User.Columns().Name, req.Name).Exist()
		liberr.ErrIsNil(ctx, err)
		if exist {
			liberr.ErrIsNil(ctx, errors.New("该用户名已被注册"))
		}

		salt := grand.Letters(10)
		password := libUtils.EncryptPassword(req.Password, salt)
		subUserId, err := orm.Data(do.User{
			Name:     req.Name,
			Password: password,
			Salt:     salt,
			Status:   req.Status,
			RoleId:   2,
			SystemId: maxSystemId.Int() + 1,
			Note:     req.Note,
		}).InsertAndGetId()
		liberr.ErrIsNil(ctx, err)

		if len(req.Project) > 0 {
			for _, v := range req.Project {
				//_, err = dao.ProjectList.Ctx(ctx).Data(do.ProjectList{
				//	AssociatedAccountId: int(subUserId),
				//}).WherePri(v).Update()
				mainReq := &sms.AllocateAccount2ProjectReq{
					AccountId: int(subUserId),
					ProjectId: v,
				}
				_, err = service.MainControllerProjectManagement().AllocateAccount2Project(ctx, mainReq)
				liberr.ErrIsNil(ctx, err)
			}
		}

		data := do.Log{
			UserId:   service.Context().GetUserId(ctx),
			UserName: service.Context().GetUsername(ctx),
			ClientIp: libUtils.GetClientIp(ctx),
			Function: "子账号后台",
			Note:     "添加子账号" + req.Name,
			SystemId: service.Context().GetSystemId(ctx),
		}
		err = utility.CreatedLog(ctx, data)
		liberr.ErrIsNil(ctx, err)
	})
	return
}

func (s *sSubUser) UpdateSubUser(ctx context.Context, req *subUser.SubUpdateReq) (res *subUser.SubUpdateRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		orm := dao.User.Ctx(ctx)
		_, err = orm.WherePri(req.ID).Data(do.User{Note: req.Note}).Update()
		liberr.ErrIsNil(ctx, err)

		username, err := orm.Fields(dao.User.Columns().Name).WherePri(req.ID).Value()
		liberr.ErrIsNil(ctx, err)

		if len(req.Project) > 0 {
			for _, v := range req.Project {
				//_, err = dao.ProjectList.Ctx(ctx).Data(do.ProjectList{
				//	AssociatedAccountId: req.ID,
				//}).WherePri(v).Update()
				mainReq := &sms.AllocateAccount2ProjectReq{
					AccountId: req.ID,
					ProjectId: v,
				}
				_, err = service.MainControllerProjectManagement().AllocateAccount2Project(ctx, mainReq)
				liberr.ErrIsNil(ctx, err)
			}
		}

		data := do.Log{
			UserId:   service.Context().GetUserId(ctx),
			UserName: service.Context().GetUsername(ctx),
			ClientIp: libUtils.GetClientIp(ctx),
			Function: "子账号后台",
			Note:     "修改子账号" + username.String(),
			SystemId: service.Context().GetSystemId(ctx),
		}
		err = utility.CreatedLog(ctx, data)
		liberr.ErrIsNil(ctx, err)
	})
	return
}

func (s *sSubUser) ChangeStatus(ctx context.Context, req *subUser.SubChangeStatusReq) (res *subUser.SubChangeStatusRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		orm := dao.User.Ctx(ctx)
		_, err = orm.WherePri(req.ID).Data(do.User{
			Status: req.Status,
		}).Update()
		liberr.ErrIsNil(ctx, err)

		var status string
		if req.Status == consts.EnumUserStatusDisable {
			status = "停用"
		} else if req.Status == consts.EnumUserStatusEnable {
			status = "启用"
		}

		username, err := orm.Fields(dao.User.Columns().Name).WherePri(req.ID).Value()
		liberr.ErrIsNil(ctx, err)

		data := do.Log{
			UserId:   service.Context().GetUserId(ctx),
			UserName: service.Context().GetUsername(ctx),
			ClientIp: libUtils.GetClientIp(ctx),
			Function: "子账号后台",
			Note:     status + "子账号" + username.String(),
			SystemId: service.Context().GetSystemId(ctx),
		}
		err = utility.CreatedLog(ctx, data)
		liberr.ErrIsNil(ctx, err)
	})
	return
}

func (s *sSubUser) DeleteSubUser(ctx context.Context, req *subUser.SubDeleteReq) (res *subUser.SubDeleteRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		if req.ID != 0 {
			err = errors.New("暂不支持删除子账号")
			liberr.ErrIsNil(ctx, err)
		}
		_, err = dao.User.Ctx(ctx).WherePri(req.ID).Delete()
		liberr.ErrIsNil(ctx, err)

		username, err := dao.User.Ctx(ctx).Fields(dao.User.Columns().Name).WherePri(req.ID).Value()
		liberr.ErrIsNil(ctx, err)

		data := do.Log{
			UserId:   service.Context().GetUserId(ctx),
			UserName: service.Context().GetUsername(ctx),
			ClientIp: libUtils.GetClientIp(ctx),
			Function: "子账号后台",
			Note:     "删除子账号" + username.String(),
			SystemId: service.Context().GetSystemId(ctx),
		}
		err = utility.CreatedLog(ctx, data)
		liberr.ErrIsNil(ctx, err)
	})
	return
}
