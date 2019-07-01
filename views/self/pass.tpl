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

<body>
<p class="layui-elem-quote">
    <a onclick="javascript: window.top.lau.tabCloseThis();" class="layui-btn layui-btn-normal">取消修改</a>
</p>

<form id="demo1" class="layui-form" action="/self/pass" method="post">
<input type="hidden" name="id" value="<<< .me.ID >>>">
  <div class="layui-form-item">
    <label class="layui-form-label">输入旧密码</label>
    <div class="layui-input-block">
      <input type="password" name="OldPassword" required lay-verify="required" placeholder="请输入旧密码" autocomplete="off" class="layui-input">
    </div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">输入新密码</label>
    <div class="layui-input-block">
      <input type="password" name="NewPassword" required lay-verify="required" placeholder="请输入新密码" autocomplete="off" class="layui-input">
    </div>
  </div><div class="layui-form-item">
    <label class="layui-form-label">再次输入新密码</label>
    <div class="layui-input-block">
      <input type="password" name="NewPassword2" required lay-verify="required" placeholder="再次输入新密码" autocomplete="off" class="layui-input">
    </div>
  </div>
  <div class="layui-form-item">
    <div class="layui-input-block">
      <button class="layui-btn" lay-submit="" lay-filter="demo1">立即提交</button>
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
        
        form.on('submit(demo1)', function(data){
            
            if(data.field.NewPassword != data.field.NewPassword2) {
              tips.error("两次密码不一致！");
              return false;
            }
            return true;
        });
        <<< if .ErrMsg >>>
        tips.error("<<< .ErrMsg >>>");
        <<< end >>>
    });
</script>
</html>