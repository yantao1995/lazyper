package controllers

import (
	"fmt"
	"lazyper/sys"
	"lazyper/web/views"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

type MainController struct {
	Ctx iris.Context
}

//主页
func (c *MainController) Get() {
	c.Ctx.WriteString(views.HTML)
}

//获取请求
func (c *MainController) GetSign() {
	btn := c.Ctx.URLParam("btn")
	p1 := c.Ctx.URLParam("p1")
	p2 := c.Ctx.URLParam("p2")
	handle(btn, p1, p2)
	c.Ctx.WriteString("success")
}

//请求处理程序
func handle(btn string, params ...string) {
	switch btn {
	case "volume1":
		sys.VolumeMute()
	case "volume2":
		sys.VolumeUp()
	case "volume3":
		sys.VolumeDown()
	case "keyboard1":
		sys.TapUp()
	case "keyboard2":
		sys.TapDown()
	case "keyboard3":
		sys.TapLeft()
	case "keyboard4":
		sys.TapRight()
	case "keyboard5":
		sys.TapEnter()
	case "keyboard7":
		sys.TapAny(params[0])
	case "keyboard8":
		sys.TapSpace()
	case "sys2":
		iInt, err := strconv.Atoi(params[1])
		if err != nil || iInt < 0 {
			fmt.Println("多少秒后关机，秒设置不正常应该大于等于0")
			return
		}
		sys.Run(sys.CancelShutdown())
		time.Sleep(time.Second)
		sys.Run(sys.Shutdown(params[1]))
		time.Sleep(time.Second)
		sys.TapEnter()
	case "sys3":
		sys.Run(sys.CancelShutdown())
	}
}
