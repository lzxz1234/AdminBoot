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
    <a href="/auth/role/" class="layui-btn layui-btn-normal">返回</a>
</p>

<form class="layui-form" action="/auth/role/add" method="post">
  <div class="layui-form-item">
    <label class="layui-form-label">名称</label>
    <div class="layui-input-block">
      <input type="text" name="Name" required  lay-verify="required" placeholder="请输入姓名" autocomplete="off" class="layui-input">
    </div>
  </div>
  <<< range $k, $v := .actions >>>
  <div class="layui-form-item">
    <label class="layui-form-label"><<< $k >>></label>
      <div class="layui-input-block">
        <<< range $v >>>
        <input type="checkbox" name="actionIds" title="<<< .Name >>>" value="<<< .ID >>>">
        <<< end >>>
      </div>
  </div>
  <<< end >>>
  <div class="layui-form-item layui-form-text">
    <label class="layui-form-label">描述信息</label>
    <div class="layui-input-block">
      <textarea name="Description" placeholder="请输入内容" required class="layui-textarea"></textarea>
    </div>
  </div>
  <div class="layui-form-item">
    <div class="layui-input-block">
      <button class="layui-btn" lay-submit lay-filter="formDemo">立即提交</button>
      <button type="reset" class="layui-btn layui-btn-primary">重置</button>
    </div>
  </div>
</form>

</body>
<script src="/static/layui/layui.js"></script>
<script>
    layui.config({base: '/static/admin/js/'}).use(['tips', 'form'], function(){
        var tips = layui.tips;
        var $ = layui.$;
        var form = layui.form;
        
        <<< if .ErrMsg >>>
        tips.error("<<< .ErrMsg >>>");
        <<< end >>>
        
    });
</script>
</html>