package routers

import (
	"lazyper/web/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Configure(b *iris.Application) {

	mvc.New(b.Party("/")).
		Register().
		Handle(new(controllers.MainController))
}
