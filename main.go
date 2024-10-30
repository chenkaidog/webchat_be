// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/binding/go_playground"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/swagger"
	_ "github.com/hertz-contrib/swagger/example/basic/docs"
	swaggerFiles "github.com/swaggo/files"
	"webchat_be/biz/config"
	"webchat_be/biz/db"
	"webchat_be/biz/middleware"
	"webchat_be/biz/util/logger"
)

//	@title			HertzTest
//	@version		1.0
//	@description	This is a demo using Hertz.

//	@contact.name	hertz-contrib
//	@contact.url	https://github.com/hertz-contrib

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8888
// @BasePath	/
// @schemes	http
func main() {
	logger.Init()
	config.Init()
	db.Init()

	vd := go_playground.NewValidator()

	h := server.Default(
		server.WithHostPorts("0.0.0.0:8000"),
		server.WithCustomValidator(vd),
	)

	h.Use(middlewareSuite()...)
	register(h)

	// swagger document
	url := swagger.URL("/swagger/doc.json")
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))

	h.Spin()
}

func middlewareSuite() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),     // panic handler
		middleware.TraceContext(), // 链路ID
		accesslog.New(
			accesslog.WithAccessLogFunc(hlog.CtxInfof),
			accesslog.WithFormat("${status} - ${latency} ${method} ${path} ${queryParams}"),
		), // 接口日志
		cors.Default(),       // 跨域请求
		middleware.Session(), // 会话
	}
}
