package sys

import "github.com/go-vgo/robotgo"

//按回车
func TapEnter() {
	robotgo.KeyTap("enter")
}

//按回车
func TapSpace() {
	robotgo.KeyTap("space")
}

//方向上
func TapUp() {
	robotgo.KeyTap("up")
}

//方向下
func TapDown() {
	robotgo.KeyTap("down")
}

//方向左
func TapLeft() {
	robotgo.KeyTap("left")
}

//方向右
func TapRight() {
	robotgo.KeyTap("right")
}

//Esc
func TapEsc() {
	robotgo.KeyTap("esc")
}

//任意键
func TapAny(any string) {
	robotgo.KeyTap(any)
}
