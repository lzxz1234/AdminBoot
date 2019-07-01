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

<form class="layui-form" action="/self/info" method="post">
<input type="hidden" name="id" value="<<< .me.ID >>>">
  <div class="layui-form-item">
    <label class="layui-form-label">姓名</label>
    <div class="layui-input-block">
      <input value="<<< .me.RealName >>>" type="text" name="RealName" required lay-verify="required" placeholder="请输入姓名" autocomplete="off" class="layui-input">
    </div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">头像</label>
    <div class="layui-input-block">
      <div class="layui-upload">
        <button type="button" class="layui-btn" id="test1">上传图片</button>
        <div class="layui-upload-list">
          <img class="layui-upload-img" id="demo1" src="<<< .me.Portrait >>>" width="200px" height="200px">
          <p id="demoText"></p>
          <input type="hidden" name="Portrait" id="Portrait" value="<<< .me.Portrait >>>">
        </div>
      </div>   
    </div>
  </div>
  <div class="layui-form-item layui-form-text">
    <label class="layui-form-label">状态签名</label>
    <div class="layui-input-block">
      <textarea name="Signature" placeholder="请输入内容" class="layui-textarea"><<< .me.Signature >>></textarea>
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
    layui.config({base: '/static/admin/js/'}).use(['tips', 'upload'], function(){
        var tips = layui.tips;
        var $ = layui.$;
        var form = layui.form;
        var upload = layui.upload;
        
        var uploadInst = upload.render({
          elem: '#test1'
          ,url: '/cmn/upload'
          ,before: function(obj){
            //预读本地文件示例，不支持ie8
            obj.preview(function(index, file, result){
              $('#demo1').attr('src', result);
            });
          }
          ,done: function(res){
            //如果上传失败
            if(res.code > 0){
              return layer.msg('上传失败');
            }
            //上传成功
            console.log(res.data.src);
            $("#Portrait").val(res.data.src);
          }
          ,error: function(){
            //演示失败状态，并实现重传
            var demoText = $('#demoText');
            demoText.html('<span style="color: #FF5722;">上传失败</span> <a class="layui-btn layui-btn-xs demo-reload">重试</a>');
            demoText.find('.demo-reload').on('click', function(){
              uploadInst.upload();
            });
          }
        });
        
        <<< if .ErrMsg >>>
        tips.error("<<< .ErrMsg >>>");
        <<< end >>>
        
    });
</script>
</html>