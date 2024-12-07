package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) ConversationList(ctx context.Context, req *sms.ConversationListReq) (res *sms.ConversationListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
