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