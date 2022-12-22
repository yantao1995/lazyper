package common

import (
	"fmt"
	"lazyper/web/views"
	"net"
	"strings"
)

func init() {
	Localhost = GetOutBoundIp()
	fmt.Println("访问地址:", Localhost)
	views.HTML = strings.ReplaceAll(views.HTML, "连接地址:", "连接地址:"+Localhost)
}

//出站流量网卡
func GetOutBoundIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return strings.Split(localAddr.String(), ":")[0]
}

//遍历ip
func SearchIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

var (
	Localhost  = "127.0.0.1"
	ListenPort = 80
)
