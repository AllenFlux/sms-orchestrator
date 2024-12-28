package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) SubProjectListForFront(ctx context.Context, req *sms.SubProjectListForFrontReq) (res *sms.SubProjectListForFrontRes, err error) {
	return service.SubControllerDeviceManagement().GetProjectList(ctx, req)
}
