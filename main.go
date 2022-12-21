package main

import (
	"encoding/json"
	"fmt"
	"lazyper/common"
	"lazyper/web/middleware"
	"lazyper/web/routers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	//错误页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		resp, _ := json.Marshal("404")
		ctx.Write(resp)
		middleware.MiddleDone(ctx)
		return
	})
	app.RegisterView(iris.HTML("./web/views", ".html"))
	app.HandleDir("./", "./web/views")
	app.UseGlobal(recover.New())
	// 配置上面配置到所有url
	mvc.New(app.Party("/"))
	// 为所有handle配置Done
	app.SetExecutionRules(iris.ExecutionRules{
		Done: iris.ExecutionOptions{Force: true},
	})
	app.Done(middleware.MiddleDone)
	routers.Configure(app)
	app.Run(iris.Addr(fmt.Sprintf(":%d", common.ListenPort)))
}
