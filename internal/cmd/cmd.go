package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gmode"
	v1 "sms_backend/internal/controller/v1"
	"sms_backend/library/libOss"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gmode.SetProduct()
			InitFun(ctx)
			apiServer := g.Server("api")
			apiServer.SetLogger(g.Log())
			//s.Group("/", func(group *ghttp.RouterGroup) {
			//	group.Middleware(ghttp.MiddlewareHandlerResponse, MiddlewareCORS)
			//	group.Bind(
			//		v1.NewCommon(),
			//	)
			//})
			apiServer.Group("/api/v1/sms", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, MiddlewareCORS)
				//service.SysUser().CheckLogin,
				//service.SystemMiddleware().PrintAndHideError, service.SystemMiddleware().Ctx, service.SystemMiddleware().Auth, service.SystemMiddleware().PrintParams)
				group.Bind(
					v1.NewSms(),
				)
			})

			apiServer.Group("/api/v1/career", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, MiddlewareCORS)
				//service.SysUser().CheckLogin,
				//service.SystemMiddleware().PrintAndHideError, service.SystemMiddleware().Ctx, service.SystemMiddleware().Auth, service.SystemMiddleware().PrintParams)
				group.Bind(
					v1.NewCareer(),
				)
			})

			apiServer.Group("/api/v1/user", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, MiddlewareCORS)
				group.Bind(
					v1.NewUser(),
				)
			})

			apiServer.Group("/api/v1/sub-user", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, MiddlewareCORS)
				group.Bind(
					v1.NewSubUser(),
				)
			})

			apiServer.Group("/api/v1/all-user", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, MiddlewareCORS)
				group.Bind(
					v1.NewAllUser(),
				)
			})

			apiServer.Group("/api/v1/role", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, MiddlewareCORS)
				group.Bind(
					v1.NewRole(),
				)
			})

			apiServer.Group("/api/v1/log", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, MiddlewareCORS)
				group.Bind(
					v1.NewLog(),
				)
			})
			apiServer.Run()
			return nil
		},
	}
)

func InitFun(ctx context.Context) {
	CompressAndClearLog(ctx)
	libOss.InitOssClient(ctx)
	DBUpgrade(ctx)
}
