package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"sms_backend/api/v1/career"
)

func (c *ControllerCareer) ReportTaskResult(ctx context.Context, req *career.ReportTaskResultReq) (res *career.ReportTaskResultRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}