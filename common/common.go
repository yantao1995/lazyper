package common

import (
	"fmt"
	"lazyper/web/views"
	"net"
	"strings"
)

func init() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && ipnet.IP.DefaultMask().String() == "ff000000" {
				fmt.Println("访问地址:", ipnet.IP.String())
				Localhost = ipnet.IP.String()
			}
		}
	}
	views.HTML = strings.ReplaceAll(views.HTML, "连接地址:", "连接地址:"+Localhost)
}

var (
	Localhost  = "127.0.0.1"
	ListenPort = 80
)
