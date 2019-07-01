<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title><<< .siteName >>></title>
    <link rel="stylesheet" href="/static/layui/css/layui.css">
    <link rel="stylesheet" href="/static/lau/lau.css">
    <script>(window.top === window.self) || (window.top.location.href = window.self.location.href);</script>
</head>
<body class="layui-layout-body layui-unselect">
<div class="layui-layout layui-layout-admin">
    <!--顶部导航开始-->
    <div class="layui-header">
        <a class="lau-logo-mini"><i class="layui-icon layui-icon-release"></i></a>
        <a class="layui-logo"><<< .siteName >>></a>
        
        <ul class="layui-nav layui-layout-right">
            <li class="layui-nav-item" lay-unselect>
                <a href="javascript:;"><img src="<<< .me.Portrait >>>" class="layui-nav-img"> <<< .me.RealName >>></a>
                <dl class="layui-nav-child">
                    <dd><a lau-href="/self/info" lau-title="基本资料">基本资料</a></dd>
                    <dd><a lau-href="/self/pass" lau-title="安全设置">安全设置</a></dd>
                    <dd><a href="/auth/logout">安全退出</a></dd>
                </dl>
            </li>
        </ul>
    </div>
    <!--顶部导航结束-->

    <!--侧边菜单开始-->
    <div class="layui-side">
        <div class="lau-side-fold"><i class="layui-icon layui-icon-shrink-right"></i></div>
        <div class="layui-side-scroll">
            <ul class="layui-nav layui-nav-tree">
                <div class="layui-side">
                    <div class="lau-side-fold"><i class="layui-icon layui-icon-shrink-right"></i></div>
                    <div class="layui-side-scroll">
                        <ul class="layui-nav layui-nav-tree">

                            <<< range $_, $item := .menus >>>
                                <<< if eq $item.ParentID  0 >>>
                                    <li class="lau-nav-item">
                                        <a class="lau-nav-header">
                                            <i class="layui-icon layui-icon-right"></i>
                                            <cite><<< $item.Name >>></cite>
                                        </a>
                                        <dl class="lau-nav-child">
                                            <<< range $_, $subItem := $.menus >>>
                                                <<< if eq $subItem.ParentID $item.ID >>>
                                                    <dd>
                                                        <a lau-href="<<< $subItem.Href >>>">
                                                            <i class="layui-icon"></i>
                                                            <cite><<< $subItem.Name >>></cite>
                                                        </a>
                                                    </dd>
                                                <<< end >>>
                                            <<< end >>>
                                        </dl>
                                    </li>
                                <<< end >>>
                            <<< end >>>

                            <li class="lau-nav-item">
                                <a class="lau-nav-header">
                                    <i class="layui-icon layui-icon-right"></i>
                                    <cite>财务管理</cite>
                                </a>
                                <dl class="lau-nav-child">
                                    <dd>
                                        <a lau-href="/admin/orderList.html">
                                            <i class="layui-icon layui-icon-danxuankuanghouxuan"></i>
                                            <cite>订单列表</cite>
                                        </a>
                                    </dd>
                                </dl>
                                <dl class="lau-nav-child">
                                    <dd>
                                        <a lau-href="/admin/packageList.html">
                                            <i class="layui-icon layui-icon-danxuankuanghouxuan"></i>
                                            <cite>套餐列表</cite>
                                        </a>
                                    </dd>
                                </dl>
                                <dl class="lau-nav-child">
                                    <dd>
                                        <a lau-href="/admin/packageAdd.html">
                                            <i class="layui-icon layui-icon-danxuankuanghouxuan"></i>
                                            <cite>新增套餐</cite>
                                        </a>
                                    </dd>
                                </dl>
                            </li>

                            <li class="lau-nav-item">
                                <a class="lau-nav-header">
                                    <i class="layui-icon layui-icon-home"></i>
                                    <cite>无子菜单</cite>
                                </a>
                            </li>
                            <!--菜单结构结束-->
                        </ul>
                    </div>
                </div>
            </ul>
        </div>
    </div>
    <!--侧边菜单结束-->

    <!--内容主体区域开始-->
    <div class="layui-body" data-type="" data-title="控制台" data-icon="layui-icon-home" data-href="/admin/dashboard"></div>
    <!--内容主体区域结束-->
</div>
</body>
<script src="/static/layui/layui.js"></script>
<script>
    var lau;
    layui.config({base: '/static/'}).extend({lau: 'lau/lau'}).use(['lau'], function () {
        lau = layui.lau;
        var layer = layui.layer,
                $ = layui.$;

        //监听事件,这个不一定要用lau-event,可以自己写
        $(document).on('click', '[lau-event]', function () {
            var _this = $(this);
            switch (_this.attr('lau-event')) {
                case 'about':
                    $.get('/html/about.html', function (html) {
                        lau.drawer({content: html});
                    });
                    break;
                case 'download':
                    layer.confirm('下载的源码仅供学习使用，如需用于正式项目，请购买授权！', function (index) {
                        window.open('https://github.com/carolkey/lying-admin/');
                        layer.close(index);
                    });
                    break;
                case "sideMenu0":
                    lau.sideMenuChange(0);
                    break;
                case "sideMenu1":
                    lau.sideMenuChange(1);
                    break;
            }
        });

    });
</script>
</html>