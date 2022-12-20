package views

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
	fmt.Printf("%#v", addrs)
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

}

var HtmlPage = `
<html>
	<head>
		<meta charset="utf-8">
		<title>遥控器</title>
	</head>
	<body>
		<hr />
		<button id="volume1">静音</button>
		<button id="volume2">音量+</button>
		<button id="volume3">音量-</button>
		<hr />
		<button id="keyboard1">上</button>
		<button id="keyboard2">下</button>
		<button id="keyboard3">左</button>
		<button id="keyboard4">右</button>
		<button id="keyboard5">回车</button>
		<hr />
		<input id="keyboard6" />
		<button id="keyboard7">任意键按下</button>
		<hr />
		<input id="sys1" />秒后
		<button id="sys2">关机</button>
		<button id="volume3">取消关机</button>
		<hr />
	</body>
</html>
`
