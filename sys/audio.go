package sys

import "github.com/go-vgo/robotgo"

//增加音量
func VolumeUp() {
	robotgo.KeyTap("audio_vol_up")
}

//降低音量
func VolumeDown() {
	robotgo.KeyTap("audio_vol_down")
}

//静音
func VolumeMute() {
	robotgo.KeyTap("audio_mute")
}
