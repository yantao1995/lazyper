package common

import (
	"fmt"
	"net"
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

}

var (
	Localhost  = "127.0.0.1"
	ListenPort = 80
)
