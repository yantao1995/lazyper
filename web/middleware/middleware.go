package middleware

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

// 所有handle的Done
func MiddleDone(ctx iris.Context) {
	fmt.Println(ctx.Request().URL.Path)
}
