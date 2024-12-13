package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/user"
)

func (c *ControllerUser) GetList(ctx context.Context, req *user.GetListReq) (res *user.GetListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}