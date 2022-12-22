package views

import (
	"fmt"
)

var HTML = fmt.Sprintf(`
<html>
	<head>
		<meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=0.5, maximum-scale=2.0, user-scalable=yes" />
		<meta name="apple-mobile-web-app-capable" content="yes" />
		<style>%s</style>
		<script src="http://libs.baidu.com/jquery/2.1.1/jquery.min.js"></script>
		<script>
		%s
		</script>
		<title>遥控器</title>
	</head>
	<body>
		<div id="host">连接地址:</div>
		<hr />
		<button id="volume1">静音</button>
		<button id="volume3">音量-</button>
		<button id="volume2">音量+</button>
		<hr />
		<button id="keyboard1">↑</button>
		<button id="keyboard2">↓</button>
		<button id="keyboard3">←</button>
		<button id="keyboard4">→</button>
		<hr />
		<button id="keyboard9">Esc</button>
		<button id="keyboard8">空格</button>
		<button id="keyboard5">回车</button>
		<hr />
		<input id="keyboard6" placeholder="输入键的名称"/>
		<button id="keyboard7">任意键按下</button>
		<hr />
		<input id="sys1" placeholder="输入时间" onkeyup="numCheck($(this))" />秒后
		<button id="sys2">定时关机</button>
		<button id="sys3">取消关机</button>
		<hr />
	</body>
</html>
`, css, js)

var js = `
$(function(){			
    $("button").click(function(){			
        var id = $(this).attr("id");
        var host = $('#host').text();
        var p1 = $('#keyboard6').val();
        var p2 = $('#sys1').val();
        if (id =="sys2" && p2 =="") {
            alert("必须输入一个非负整数");
            return
        }
        if (id =="keyboard7" && p1 =="") {
            alert("必须输入一个键的名称");
            return
        }
        var url ="/sign?btn="+id+"&p1="+p1+"&p2="+p2
        $.get(url,function(data,status){
            console.log("数据: " + data + "\n状态: " + status);
        });
    })
})

function numCheck($this){
    var value = $this.val() ;
    if(isNaN(value)){
        $this.val(0) ;
    }else if(value==""){
        
    }else if(value<0){
        $this.val(0) ;
    }else if(value.indexOf(".")!=-1){
        $this.val(value.substring(0,value.indexOf(".")))
    }
}
`

var css = `
body{
	width: 100%;
	height: 100%;
	background-color: aliceblue;
	position: absolute;
}

input{
	border-radius: 7px;
	width: 100px;
	height: 80px;
	text-align: center;
	text-decoration: none;
	display: inline-block;
	font-size: 16px; 
}

button{
	text-decoration: none;
	border: none;
	padding: 40px 20px;
	font-size: 16px;
	background-color: green;
	color: #fff;
	border-radius: 5px;
	box-shadow: 7px 6px 26px 1px rgba(0, 0, 0, 0.24);
	cursor: pointer;
	outline: none;
	transition: 0.2s all;
}

button:active {
	transform: scale(0.90);
	box-shadow: 3px 2px 22px 1px rgba(0, 0, 0, 0.24);
}
`
