<!DOCTYPE html>
<html lang="en" class="am-touch js cssanimations">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title><<< .siteName >>></title>
    <link rel="stylesheet" href="/static/layui/css/layui.css">
    <link rel="stylesheet" href="/static/lau/lau.css">
</head>

<body style="padding: 15px;">
<p class="layui-elem-quote">
<<< if .RBAC_USER_MOD >>>
    <a href="/auth/user/add" class="layui-btn layui-btn-normal">添加员工</a>
<<< end >>>
</p>
<table class="layui-table" lay-data="{url:'/auth/user/list', page:true, id:'test'}" lay-filter="test">
    <thead>
    <tr>
        <th lay-data="{field:'ID', width:80, sort: true}">ID</th>
        <th lay-data="{field:'RealName', width:150}">姓名</th>
        <th lay-data="{field:'UserName', width:150}">账户</th>
        <th lay-data="{field:'Roles', width:250, templet: '#rolesTpl'}">角色</th>
        <th lay-data="{field:'Signature'}">签名</th>
        <th lay-data="{field:'LastLoginTime', width:180}">最后登录时间</th>
        <th lay-data="{fixed: 'right', width:150, align:'center', toolbar: '#barDemo'}">操作</th>
    </tr>
    </thead>
</table>

<script type="text/html" id="rolesTpl">
    {{# layui.each(d.Roles, function(index, item) { }}
    <span class="layui-badge layui-bg-black">{{ item.Name  }}</span>
    {{# }); }}
</script>

<script type="text/html" id="barDemo">
<<< if .RBAC_USER_MOD >>>
{{# if(d.UserName != "admin" ) { }}
    {{# if(d.State == 0) { }}
        <a class="layui-btn layui-btn-xs" lay-event="enable">启用</a>
    {{# } else { }}
        <a class="layui-btn layui-btn-xs layui-btn-danger" lay-event="disable">禁用</a>
    {{# } }}
{{# } else { }}
    <span class="layui-btn layui-btn-xs" >启用</span>
{{# } }}
    <a class="layui-btn layui-btn-xs" lay-event="edit">编辑</a>
<<< end >>>
</script>

</body>
<script src="/static/layui/layui.js"></script>
<script>
    layui.config({base: '/static/admin/js/'}).use(['table', 'tips'], function(){
        var table = layui.table;
        var tips = layui.tips;
        var $ = layui.$;

        table.on('tool(test)', function(obj){
            var data = obj.data; //获得当前行数据
            var layEvent = obj.event; //获得 lay-event 对应的值（也可以是表头的 event 参数对应的值）
            var tr = obj.tr; //获得当前行 tr 的DOM对象

            if(layEvent === 'disable'){
                $.ajax({
                    url: "/auth/user/modState?state=0&id=" + data.ID, 
                    type: 'post',
                    success: function(resp) {
                        console.log(resp);
                        window.parent.lau.reload();
                    }
                })
            } else if(layEvent === 'enable') {
                $.ajax({
                    url: "/auth/user/modState?state=1&id=" + data.ID, 
                    type: 'post',
                    success: function(resp) {
                        console.log(resp);
                        window.parent.lau.reload();
                    }
                })
            } else if(layEvent === 'edit'){ //编辑
                window.location.href = "/auth/user/mod?id=" + data.ID;
            }
        });;
    });
</script>
</html>