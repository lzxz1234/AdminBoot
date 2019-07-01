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
<<< if .RBAC_ROLE_MOD >>>
    <a href="/auth/role/add" class="layui-btn layui-btn-normal">添加角色</a>
<<< end >>>
</p>
<table class="layui-table" lay-data="{url:'/auth/role/list', page:true, id:'test'}" lay-filter="test">
    <thead>
    <tr>
        <th lay-data="{field:'ID', width:80, sort: true}">ID</th>
        <th lay-data="{field:'Name', width:150}">名称</th>
        <th lay-data="{field:'Actions', width:400, templet: '#actionTpl'}">权限集</th>
        <th lay-data="{field:'Actions', width:400, templet: '#userTpl'}">人员</th>
        <th lay-data="{field:'Description'}">描述</th>
        <th lay-data="{fixed: 'right', width:150, align:'center', toolbar: '#barDemo'}">操作</th>
    </tr>
    </thead>
</table>

<script type="text/html" id="actionTpl">
    {{# layui.each(d.Actions, function(index, item) { }}
    <span class="layui-badge layui-bg-blue">{{ item.Group  }}</span>
    {{# }); }}
</script>

<script type="text/html" id="userTpl">
    {{# layui.each(d.Users, function(index, item) { }}
    <span class="layui-badge layui-bg-black">{{ item.RealName  }}</span>
    {{# }); }}
</script>

<script type="text/html" id="barDemo">
<<< if .RBAC_ROLE_MOD >>>
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

            if(layEvent === 'edit'){ //编辑
                window.location.href = "/auth/role/mod?id=" + data.ID;
            }
        });;
    });
</script>
</html>