
<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title><<< .siteName >>></title>
    <link rel="stylesheet" href="/static/layui/css/layui.css">
    <link rel="stylesheet" href="/static/css/sign.css">
    <script>(window.top === window.self) || (window.top.location.href = window.self.location.href);</script>
</head>
<body class="layui-unselect lau-sign-body">

<form action="/admin/login" method="POST" enctype="multipart/form-data">
    <div class="layui-form layui-form-pane lau-sign-form">
        <h1 class="lau-sign-title">SIGN IN <<< .siteName >>></h1>
        <p class="lau-sign-subtitle"><<< .siteSubName >>></p>
        <div class="layui-form-item">
            <label class="layui-form-label"><i class="layui-icon layui-icon-username"></i> 账　号</label>
            <div class="layui-input-block">
                <input type="text" name="username" placeholder="请输入用户名" autocomplete="off" class="layui-input" value="admin">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label"><i class="layui-icon layui-icon-password"></i> 密　码</label>
            <div class="layui-input-block">
                <input type="password" name="password" placeholder="请输入密码" autocomplete="off" class="layui-input" value="admin123!">
            </div>
        </div>
        
        <div class="layui-form-item">
            <input type="checkbox" name="remember" lay-skin="primary" title="记住密码">
            <a class="lau-sign-forgot lau-sign-link" >忘记密码？</a>
        </div>
        <div class="layui-form-item">
            <button type="button" class="layui-btn layui-btn-fluid" lay-submit lay-filter="login">登 入</button>
        </div>
    </div>
    <div class="layui-trans lau-sign-footer">
        <p>@ 2018 <a href="http://lau.revoke.cc/" target="_blank">lau.revoke.cc</a> SATA License</p>
        <p>
            <span><a href="https://jq.qq.com/?_wv=1027&k=5qarx5y" target="_blank">获取授权</a></span>
            <span><a href="http://lau.revoke.cc/" target="_blank">在线演示</a></span>
            <span><a href="https://github.com/carolkey/lying-admin" target="_blank">源码下载</a></span>
            <span><a href="mailto:su@revoke.cc" target="_blank">联系作者</a></span>
        </p>
    </div>
</form>

</body>
<script src="/static/layui/layui.js"></script>
<script>
    layui.config({base: '/static/admin/js/'}).define(['layer', 'form', 'tips'], function(exports) {
        var form = layui.form,
                layer = layui.layer,
                $ = layui.$,
                tips = layui.tips;

        //ajax请求出错提示
        $(document).ajaxError(function (event, request, setting) {
            if (request.status === 200) {
                tips.error('Invalid response');
            } else {
                tips.error(request.status + ': ' + request.statusText);
            }
        });

        //登陆
        form.on('submit(login)', function (data) {

            if (data.field.username == '') {
                tips.warning('用户名不能为空');
                return false;
            } else if (data.field.password == '') {
                tips.warning('密码不能为空');
                return false;
            }

            //登陆中
            tips.loading('登陆中...', 0, -1);

            //发送登陆表单
            $.post('/auth/login', data.field, function (json) {
                if (json.code == 0) {
                    tips.success(json.msg, function () {
                        location.href = '/auth/';
                    });
                } else {
                    console.log(json)
                    tips.error(json.msg);
                }
            }, 'json');

            return false;
        });

        exports('login', {});

        <<< if .ErrMsg >>>
        tips.error("<<< .ErrMsg >>>");
        <<< end >>>
    });
</script>
</html>