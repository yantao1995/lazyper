package sys

//关闭所有定时关机任务
func CancelShutdown() []string {
	return []string{"shutdown", "-a"}
}

//second >= 0 时生效
func Shutdown(second string) []string {
	return []string{"shutdown", "-s", "-t", second}
}
