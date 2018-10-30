<!DOCTYPE html>
<html lang='zh-CN'>
<link href="/assets/projects/application-1184fbb96d2e3b8ded59d23df2a94ce2.css" media="screen" rel="stylesheet" type="text/css" />
<head>
<title>ai_redis/redis_test.go · realwrtoff/ai_redis - 码云 Gitee.com</title>
<link href="/assets/favicon-e87ded4710611ed62adc859698277663.ico" rel="shortcut icon" type="image/vnd.microsoft.icon" />
<meta content='gitee.com/realwrtoff/ai_redis git https://gitee.com/realwrtoff/ai_redis.git' name='go-import'>
<meta charset='utf-8'>
<meta content='always' name='referrer'>
<meta content='码云' property='og:site_name'>
<meta content='Object' property='og:type'>
<meta content='http://gitee.com/realwrtoff/ai_redis/blob/master/ai_redis/redis_test.go' property='og:url'>
<meta content='https://gitee.com/big_logo_gitee.jpg' itemprop='image' property='og:image'>
<meta content='ai_redis/redis_test.go · realwrtoff/ai_redis - 码云 Gitee.com' itemprop='name' property='og:title'>
<meta content='码云(gitee.com)是开源中国推出的代码托管平台，支持 Git 和 SVN，提供免费的私有仓库托管。目前已有超过 300 万的开发者选择码云。' property='og:description'>
<meta content='码云,代码托管,git,Git@OSC,gitee.com,开源,项目管理,版本控制,开源代码,代码分享,项目协作,开源项目托管,免费代码托管,Git代码托管,Git托管服务' name='Keywords'>
<meta content='redis 客户端' itemprop='description' name='Description'>
<meta content="IE=edge" http-equiv="X-UA-Compatible" />
<meta content="authenticity_token" name="csrf-param" />
<meta content="dAsUd5Q3+QefknbrpOW8s3VNQQJi0xdtPIVOTvDmLuI=" name="csrf-token" />

<link href="/assets/application-14f5f827bc14bc7435e86d1fb5de459b.css" media="screen" rel="stylesheet" type="text/css" />
<script src="/assets/application-c0d1385daa4e0cff531000e14e5d274f.js" type="text/javascript"></script>
<script src="/assets/lib/jquery.timeago.zh-CN-bcd91c2c27a815fa9a395595874b592b.js" type="text/javascript"></script>

<script type="text/javascript">
//<![CDATA[
window.gon = {};gon.locale="zh-CN";gon.http_clone="https://gitee.com/realwrtoff/ai_redis.git";gon.user_project="realwrtoff/ai_redis";gon.manage_branch="管理分支";gon.manage_tag="管理标签";gon.enterprise_id=0;gon.ref="master";
//]]>
</script>
<script src="//res.wx.qq.com/open/js/jweixin-1.2.0.js" type="text/javascript"></script>
<script>
  var title = document.title.replace(/( - Gitee| - 码云)$/, ''),
    imgUrl = '',
    imgUrlEl = document.querySelector('meta[itemprop=image]')
  if (imgUrlEl) {
    imgUrl = imgUrlEl.getAttribute('content')
  } else {
    imgUrl = "https://gitee.com/logo_for_wechat.png"
  }
  wx.config({
    debug: false,
    appId: "wxff219d611a159737",
    timestamp: "1540868002",
    nonceStr: "50cde5f965642d12a032077439027ba5",
    signature: "e3b1fdfdafb03ca70a9712ef17e0bc1f9574f5c5",
    jsApiList: [
      'onMenuShareTimeline',
      'onMenuShareAppMessage'
    ]
  });
  wx.ready(function () {
    wx.onMenuShareTimeline({
      title: title, // 分享标题
      link: "https://gitee.com/realwrtoff/ai_redis/blob/master/ai_redis/redis_test.go", // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
      imgUrl: imgUrl // 分享图标
    });
    wx.onMenuShareAppMessage({
      title: title, // 分享标题
      link: "https://gitee.com/realwrtoff/ai_redis/blob/master/ai_redis/redis_test.go", // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
      desc: document.querySelector('meta[name=Description]').getAttribute('content'),
      imgUrl: imgUrl // 分享图标
    });
  });
  wx.error(function(res){
    console.error('err', res)
  });
</script>

<script type='text/x-mathjax-config'>
MathJax.Hub.Config({
  tex2jax: {
    inlineMath: [['$','$'], ['\\(','\\)']],
    displayMath: [["$$","$$"],["\\[","\\]"]],
    processEscapes: true,
    skipTags: ['script', 'noscript', 'style', 'textarea', 'pre', 'code'],
    ignoreClass: "container|files",
    processClass: "markdown-body"
  }
});
</script>
<script src="https://gitee.com/uploads/resources/MathJax-2.7.2/MathJax.js?config=TeX-AMS-MML_HTMLorMML" type="text/javascript"></script>

<script>
  var messages = {
    'zh-CN': {
      addResult: '增加 <b>{term}</b>',
      count: '已选择 {count}',
      maxSelections: '最多 {maxCount} 个选择',
      noResults: '未找到结果',
      serverError: '连接服务器时发生错误'
    },
    'zh-TW': {
      addResult: '增加 <b>{term}</b>',
      count: '已選擇 {count}',
      maxSelections: '最多 {maxCount} 個選擇',
      noResults: '未找到結果',
      serverError: '連接服務器時發生錯誤'
    }
  }
  
  if (messages[gon.locale]) {
    $.fn.dropdown.settings.message = messages[gon.locale]
  }
</script>

<!--[if lt IE 10]>
<script>
    window.location.href = "/incompatible.html";
</script>
<![endif]-->
</head>

<body class='git-project lang-zh-CN'>
<script src="/assets/projects/app-1917601997c5d088c56d479a64b39ac8.js" type="text/javascript"></script>
<header class='common-header fixed noborder' id='git-header-nav'>
<div class='ui container'>
<div class='ui menu'>
<div class='item gitosc-logo'>
<a href="/"><img class='ui inline image' height='28' src='/logo.svg?20171024' width='95'>
<img class='ui inline black image' height='28' src='/logo-black.svg?20171024' width='95'>
</a></div>
<a href="/explore" class="item " title="开源软件">开源软件
</a><a href="/enterprises" class="item " title="企业版">企业版
<sup class='ui red label'>
特惠
</sup>
</a><a href="/education" class="item " title="高校版">高校版
</a><a href="https://blog.gitee.com/" class="item" id="gitee-blog" target="_blank" title="博客">博客
</a><div class='right menu userbar' id='git-nav-user-bar'>
<div class='item'>
<form accept-charset="UTF-8" action="/search" autocomplete="on" data-text-filter="搜索格式不正确" data-text-require="搜索关键字不能少于1个" id="navbar-search-form" method="get"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /></div>
<div class='ui mini fluid input'>
<input id="navbar-search-input" name="q" placeholder="搜索项目、代码片段..." type="text" value="" />
<input id="navbar-search-type" name="type" type="hidden" />
</div>
</form>


</div>
<a href="/signup" class="item">注册
</a><a href="/login" class="item">登录
</a><script>
  $('.destroy-user-session').click(function() {
    $.cookie('access_token', null, { path: '/' });
  })
</script>

</div>
</div>
</div>
</header>

<script>
  Gitee.initNavbar()
  Gitee.initRepoRemoteWay()
</script>

<!--[if lt IE 10]>
<script>
  window.location.href = "/incompatible.html"
</script>
<![endif]-->

<div class='fixed-notice-messages'>
<div class='ui info message' id='git-bulletin'>
<a href="https:" />码云 Gitee IDE 全新上线——支持 Git 管理的轻量在线编码环境</a>
<i class='icon remove' id='remove-bulletin'></i>
</div>
<div class='ui container'>
<div class='flash-messages' id='messages-container'></div>
</div>
</div>
<script>
  (function() {
    $(function() {
      var $error_box, alertTip, notify_content, notify_options, setCookie, template;
  
      template = '<div data-notify="container" class="ui {0} message" role="alert">' + '<i data-notify="dismiss" class="close icon"></i>' + '<span data-notify="message">{2}</span>' + '</div>';
      notify_content = null;
      notify_options = {};
      alertTip = '';
      $error_box = $(".flash_error.flash_error_box");
      if (notify_options.type === 'error' && $error_box.length > 0 && !$.isEmptyObject(notify_content.message)) {
        if (notify_content.message === 'captcha_fail') {
          alertTip = "验证码不正确";
        } else if (notify_content.message === 'not_found_in_database') {
          alertTip = "帐号或者密码错误";
        } else if (notify_content.message === 'not_found_and_show_captcha') {
          alertTip = "帐号或者密码错误";
        } else {
          alertTip = notify_content.message;
        }
        $error_box.html(alertTip).show();
      } else if (notify_content) {
        notify_options.delay = 3000;
        notify_options.template = template;
        notify_options.offset = {
          x: 10,
          y: 30
        };
        notify_options.element = '#messages-container';
        $.notify(notify_content, notify_options);
      }
      setCookie = function(name, value) {
        $.cookie(name, value, {
          path: '/',
          expires: 365
        });
      };
      $('#remove-bulletin').on('click', function() {
        setCookie('remove_bulletin', "gitee-maintain-1539956609");
        $('#git-bulletin').hide();
      });
      $('#remove-member-bulletin').on('click', function() {
        setCookie('remove_member_bulletin', "gitee_member_bulletin");
        $(this).parent().hide();
      });
      return $('#remove-gift-bulletin').on('click', function() {
        setCookie('remove_gift_bulletin', "gitee-gift-bulletin");
        $(this).parent().hide();
      });
    });
  
  }).call(this);
</script>

<div class='git-project-header'>
<div class='fixed-notice-messages'>
<div class='ui info icon floating message green' id='fetch-ok' style='display: none'>
<div class='content'>
<div class='header status-title'>
<i class='info icon status-icon'></i>
代码拉取完成，页面将自动刷新
</div>
</div>
</div>
<div class='ui info icon floating message error' id='fetch-error' style='display: none'>
<div class='content'>
<div class='header status-title'>
<i class='info icon status-icon'></i>
<span class='error_msg'></span>
</div>
</div>
</div>
</div>
<div class='ui container'>

<div class='git-project-header-details'>
<div class='git-project-header-actions'>
<span class='ui basic buttons'>
<a class='ui button donate' id='project-donate' title='捐赠'>
<i class='iconfont icon-donation'></i>
捐赠
</a>
<a class='ui button action-social-count' href='/realwrtoff/ai_redis#project-donate-overview'>0</a>
<div class='ui tiny modal project-donate-modal' id='project-donate-modal'>
<i class='iconfont icon-close close'></i>
<div class='header'>捐赠</div>
<div class='content'>
该项目未开启捐赠功能，可发送私信通知作者开启
</div>
<div class='actions'>
<a class='ui blank button cancel'>取消</a>
<a class='ui orange ok button' href='/notifications/messages/1553405'>发送私信</a>
</div>
</div>
<div class='ui small modal wepay-qrcode'>
<i class='iconfont icon-close close'></i>
<div class='header'>
扫描微信二维码支付
<span class='wepay-cash'></span>
</div>
<div class='content weqcode-center'>
<img id='wepay-qrcode' src=''>
</div>
<div class='actions'>
<div class='ui cancel blank button'>取消</div>
<div class='ui ok orange button'>
支付完成
</div>
</div>
</div>
</span>
<div class='ui mini modal' id='confirm-alipay-modal'>
<div class='header'>支付提示</div>
<div class='content'>
将跳转至支付宝完成支付
</div>
<div class='actions'>
<div class='ui approve orange button'>
确定
</div>
<div class='ui blank cancel button'>
取消
</div>
</div>
</div>

<span class='ui basic buttons watch-container'>
<a href="/login" class="ui button watch" title="登录以Watch"><i class='iconfont icon-watch'></i>
Watch
</a><a href="/realwrtoff/ai_redis/watchers" class="ui button action-social-count" title="1">1
</a></span>
<span class='ui basic buttons star-container'>
<a href="/login" class="ui button star" title="你必须登录后才可以star一个仓库"><i class='iconfont icon-star'></i>
Star
</a><a href="/realwrtoff/ai_redis/stargazers" class="ui button action-social-count" title="0">0
</a></span>
<span class='ui basic buttons fork-container' data-content='无权 Fork 此项目'>
<a href="/login" class="ui button fork "><i class='iconfont icon-fork'></i>
Fork
</a><a href="/realwrtoff/ai_redis/members" class="ui button action-social-count" title="0">0
</a></span>
</div>
<h2 class='git-project-title'>
<span class="project-title"><i class="project-icon iconfont icon-project-public" title="这是一个公开项目"></i> <a href="/realwrtoff" class="author" title="realwrtoff">realwrtoff</a> / <a href="/realwrtoff/ai_redis" class="repository" style="padding-bottom: 0px" target="" title="ai_redis">ai_redis</a></span><span class="project-badges"><a href="/explore/recommend?lang=Go" class="ui small label" id="project-language" target="_blank" title="主要编程语言">Go</a><a href="/realwrtoff/ai_redis/tree/master/LICENSE" class="ui small license label" title="开源许可协议">BSD-3-Clause</a></span>

<input id="project_title" name="project_title" type="hidden" value="realwrtoff/ai_redis" />
</h2>
</div>
</div>
<div class='row' id='import-result-message' style='padding-top: 0px; display: none'>
<div class='sixteen wide column'>
<div class='ui icon yellow message status-color'>
<i class='info icon status-icon' style='width:60px;padding-right:12px;'></i>
<i class='close icon'></i>
<div class='header status-title'>
同步状态
</div>
<span class='status-message'></span>
</div>
</div>
</div>
<div class='ui small modal' id='modal-fork-project'>
<i class='icon-close iconfont close'></i>
<div class='header'>
Fork 项目
</div>
<div class='content'>
<div class='fork-info-content'>
<div class='ui segment fork_project_loader'>
<div class='ui active inverted dimmer'>
<div class='ui text loader'>加载中</div>
</div>
</div>
</div>
</div>
<div class='actions fork-action hide'>
<a class='cancel'>&emsp;取消&emsp;</a>
<div class='ui large button orange ok'>&emsp;确认&emsp;</div>
</div>
</div>
<script>
  (function() {
    $('.fork-container').popup({
      inline: true,
      hoverable: true,
      position: 'bottom left'
    });
  
  }).call(this);
</script>
<script>
  (function() {
    $('#modal-fork-project').modal({
      transition: 'fade'
    });
  
    $('.checkbox.sync-wiki').checkbox();
  
  }).call(this);
</script>
<style>
  i.loading{-webkit-animation:icon-loading 1.2s linear infinite;animation:icon-loading 1.2s linear infinite}.qrcode_cs{float:left}.check-sync-wiki{float:left;height:28px;line-height:28px}.check-sync-wiki .sync-wiki-warn{color:#e28560}
</style>

<div class='git-project-nav'>
<div class='ui container'>
<div class='ui secondary pointing menu'>
<a href="/realwrtoff/ai_redis" class="item active"><i class='iconfont icon-code'></i>
代码
</a><a href="/realwrtoff/ai_redis/issues" class="item "><i class='iconfont icon-issue'></i>
Issues
<span class='ui mini circular label'>
0
</span>
</a><a href="/realwrtoff/ai_redis/pulls" class="item "><i class='iconfont icon-pull-request'></i>
Pull Requests
<span class='ui mini circular label'>
0
</span>
</a><a href="/realwrtoff/ai_redis/attach_files" class="item "><i class='iconfont icon-annex'></i>
附件
<span class='ui mini circular label'>0</span>
</a><a href="/realwrtoff/ai_redis/wikis" class="item "><i class='iconfont icon-wiki'></i>
Wiki
<span class='ui mini circular label'>
0
</span>
</a><a href="/realwrtoff/ai_redis/graph/master" class="item  "><i class='iconfont icon-statistics'></i>
统计
</a><div class='item'>
<div class='ui pointing top right dropdown git-project-service'>
<div class='text'>
<i class='iconfont icon-service'></i>
服务
</div>
<i class='dropdown icon'></i>
<div class='menu' style='display:none'>
<a href="/realwrtoff/ai_redis/paas/code_pipeline" class="item"><img alt="Code_pipeline" src="/assets/code_pipeline-502dbffc863fcc0b6867628c9e8e163e.svg" />
<div class='item-title'>
阿里云CodePipeline
</div>
</a><a href="/realwrtoff/ai_redis/paas/huaweicloud_swr" class="item"><img alt="Huaweirqy" src="/assets/huaweirqy-c24d1adc35c0ecb2bc1e39f6e033835c.png" />
<div class='item-title'>
华为容器云
</div>
</a><a href="/realwrtoff/ai_redis/paas/huaweicloud_cse" class="item"><img alt="Hauweiwfw" src="/assets/hauweiwfw-9edbcf60ca0a5d0de81d2fc804efa9ac.png" />
<div class='item-title'>
华为微服务平台
</div>
</a><a href="/realwrtoff/ai_redis/paas/tencent_hub" class="item"><img alt="Tencent_hub" src="/assets/tencent_hub-e263083dfb38e2149d8a9eca5347f4d5.png" />
<div class='item-title'>
Tencent Hub
</div>
</a><a href="/realwrtoff/ai_redis/paas/select_platform" class="item"><img alt="Mopaas_mini" src="/assets/mopaas_mini-72f0d5aeae31630ae89f624dfb0c23ca.png" />
<div class='item-title'>
MoPaaS V3
</div>
</a><a href="/realwrtoff/ai_redis/quality_analyses?platform=sonar_qube" class="item"><img alt="Sonar_mini" src="/assets/sonar_mini-6270e37950512a0bf0a05ac5b9b11243.png" />
<div class='item-title'>
质量分析
</div>
</a><a href="/realwrtoff/ai_redis/quality_analyses?platform=codesafe" class="item"><img alt="Dmws" src="/assets/dmws-7fce33f3494048913a196a40f998a9ba.png" />
<div class='item-title'>
源代码缺陷检测
</div>
</a><a href="/realwrtoff/ai_redis/pages" class="item"><img alt="Gitee" src="/assets/gitee-d6fb391be28450a587df71dda0325f60.png" />
<div class='item-title'>
Gitee Pages
</div>
</a><button class='ui orange basic button quit-button' id='quiting-button'>
我知道了，不再自动展开
</button>
</div>
</div>
</div>
</div>
</div>
</div>
<script>
  $('.git-project-nav .ui.dropdown').dropdown({ action: 'nothing' });
</script>
<style>
  .git-project-nav i.checkmark.icon {
    color: green;
  }
  #quiting-button {
    display: none;
  }
</style>
<script>
  isSignIn = false
  isClickGuide = false
  $('#git-versions.dropdown').dropdown();
  $.ajax({
    url:"/realwrtoff/ai_redis/access/add_access_log",
    type:"GET"
  });
  $('#quiting-button').on('click',function() {
    $('.git-project-service').click();
    if (isSignIn) {
      $.post("/projects/set_service_guide")
    }
    $.cookie("Serve_State", true, { expires: 3650, path: '/'})
    $('#quiting-button').hide();
  });
  if (!(isClickGuide || $.cookie("Serve_State") == 'true')) {
    $('.git-project-service').click()
    $('#quiting-button').show()
  }
</script>

</div>
<div class='register-guide'>
<div class='register-container'>
<div class='regist'>
加入码云
</div>
<div class='description'>
与超过 300 万 开发者一起发现、参与优秀开源项目，私有仓库也完全免费 ：）
</div>
<a href="/signup?from=project-guide" class="ui orange button free-registion">免费加入</a>
<div class='login'>
已有帐号？
<a href="/login?from=project-guide">立即登录</a>
</div>
</div>
</div>

<div class='git-project-content-wrapper'>
<div class='ui container'>
<div class='git-project-content' id='git-project-content'>
<div class='row' id='git-detail-clone'>
<div class='git-project-desc-wrapper'>
<div class='git-project-desc'>
<span class='git-project-desc-text'>
redis 客户端
</span>
<span class='git-project-homepage'>
</span>
</div>
<div class='git-project-desc-edit ui form'>
<div class='fields'>
<div class='eight wide field'>
<div class='ui small input'>
<input name='project[description]' placeholder='描述' type='text' value='redis 客户端'>
</div>
</div>
<div class='four wide field'>
<div class='ui small input'>
<input data-regex-value='(^$)|(^(http|https):\/\/(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).*)|(^(http|https):\/\/[a-zA-Z0-9]+([_\-\.]{1}[a-zA-Z0-9]+)*\.[a-zA-Z]{2,10}(:[0-9]{1,10})?(\/.*)?$)' name='project[homepage]' placeholder='项目主页(eg: https://gitee.com)' type='text'>
</div>
</div>
<button class='ui orange button btn-save'>
保存更改
</button>
<button class='ui button btn-cancel-edit'>
取消
</button>
</div>
</div>
<script>
  $(function () {
    new ProjectInfoEditor({
      el: '.git-project-desc-wrapper',
      homepage: null,
      description: "redis 客户端",
      url: "/realwrtoff/ai_redis/update_description",
      modalHelper: Gitee.modalHelper
    })
  })
</script>
</div>

</div>
<div class='git-project-summary' id='git-project-summary'>
<div class='summary-viewer'>
<div class='viewer-wrapper'>
<ul>
<li>
<a href="/realwrtoff/ai_redis/commits/master"><i class='iconfont icon-commit'></i>
<b>3</b>
次提交
</a></li>
<li>
<a href="/realwrtoff/ai_redis/branches"><i class='iconfont icon-branches'></i>
<b>1</b>
个分支
</a></li>
<li>
<a href="/realwrtoff/ai_redis/tags"><i class='iconfont icon-tag'></i>
<b>0</b>
个标签
</a></li>
<li>
<a href="/realwrtoff/ai_redis/releases"><i class='iconfont icon-release'></i>
<b>0</b>
个发行版
</a></li>
<li>
<a href="/realwrtoff/ai_redis/contributors?ref=master"><i class='iconfont icon-collaborators'></i>
<b class='contributor-count' data-url='/realwrtoff/ai_redis/contributors_count?ref=master'></b>
<span class='contributor-text'>
正在获取贡献者
</span>
</a></li>
</ul>
<ul>
</ul>
</div>
</div>
<div class='summary-languages' title='点击查看语言占比'>
</div>

</div>
<div class='git-project-bread' id='git-project-bread'>
<div class='ui right floated horizontal list'>
<div class='item'>
<div class='ui orange button dropdown' id='btn-dl-or-clone'>
克隆/下载
<i class='dropdown icon'></i>
</div>
</div>
</div>
<div class='ui horizontal list'>
<div class='item item git-project-branch-item'>
<input id="path" name="path" type="hidden" value="ai_redis/redis_test.go" />
<div class='ui top left pointing dropdown gradient button dropdown-has-tabs' id='git-project-branch'>
<input id="ref" name="ref" type="hidden" value="master" />
<div class='default text'>
master
</div>
<i class='dropdown icon'></i>
<div class='menu'>
<div class='ui left icon search input'>
<i class='iconfont icon-search'></i>
<input placeholder='搜索分支' type='text'>
</div>
<div class='tab-menu'>
<div class='tab-menu-action' data-tab='branches'>
<a href="/realwrtoff/ai_redis/branches/recent" class="ui link button">管理</a>
</div>
<div class='tab-menu-action' data-tab='tags'>
<a href="/realwrtoff/ai_redis/tags" class="ui link button">管理</a>
</div>
<div class='tab-menu-item' data-placeholder='搜索分支' data-tab='branches'>
分支 (1)
</div>
</div>
<div class='tab scrolling menu' data-tab='branches'>
<div class='item' data-value='master'>master</div>
</div>
</div>
</div>
</div>
<div class='item'>
<div class='repo-index repo-none-index' style='margin-left:0px;'>
<div class='ui horizontal list repo-action-list'>
<div class='item'>
<div class='ui pointing right top dropdown gradient button' id='git-project-file'>
<div class='text'>文件</div>
<i class='dropdown icon'></i>
<div class='menu'>
<div class='disabled item'>新建文件</div>
<div class='disabled item'>上传文件</div>
<a class='item repo-action' id='search-files'>
搜索文件
</a>
</div>
</div>
</div>
<div class='item'>
<a href="/-/ide/project/realwrtoff/ai_redis/edit/master" class="ui gradient button webide" target="_blank">Web IDE</a>
</div>
<div class='item'>
<a href="/realwrtoff/ai_redis/widget" class="ui gradient button repo-action"><i class='iconfont icon-widget icon-orange'></i>
挂件
</a></div>
</div>
</div>
</div>
<div class='breadcrumb_path item path-breadcrumb-contrainer'>
<div class='ui breadcrumb path project-path-breadcrumb' id='path-breadcrumb'>
<a href="/realwrtoff/ai_redis/tree/master" class="section repo-name" data-direction="back" style="font-weight: bold">ai_redis
</a><div class='divider'>
/
</div>
<strong>
<a href="/realwrtoff/ai_redis/tree/master/ai_redis" class="section" data-direction="back"><span class='cblue'>
ai_redis
</span>
</a></strong>
<div class='divider'>
/
</div>
<strong>
redis_test.go
</strong>
</div>

</div>
</div>
<div class='ui popup bottom right transition hidden git-project-download-panel'>
<div class='ui small secondary pointing menu'>
<a class='item active' data-type='http' data-url='https://gitee.com/realwrtoff/ai_redis.git'>HTTPS</a>
<a class='item' data-type='ssh' data-url='git@gitee.com:realwrtoff/ai_redis.git'>SSH</a>
</div>
<div class='ui fluid right labeled small input'>
<input id="project_clone_url" name="project_clone_url" onclick="focus();select()" readonly="readonly" type="text" value="https://gitee.com/realwrtoff/ai_redis.git" />
<div class='ui basic label'>
<div class='ui small basic orange button' data-clipboard-target='#project_clone_url' id='btn-copy-clone-url'>
复制
</div>
</div>
</div>
<hr>
<a href="/login?url_to=/realwrtoff/ai_redis/repository/archive/master.zip" class="ui fluid download link button" data-confirm="为了避免盗链导致资源占用，登录后才能下载，确定登录?
强烈推荐: git clone &quot;https://gitee.com/realwrtoff/ai_redis.git&quot;"><i class='icon download'></i>
下载ZIP
</a><hr>

</div>
</div>
<style>
  .ui.dropdown .menu>.header{text-transform:none}
</style>
<script>
  $(document).ready(function () {
    var $gitProjectSummary = $('#git-project-summary');
    var $gitProjectLanguages = $gitProjectSummary.find('.summary-languages');
    var $statsSwitcherWrapper = $gitProjectSummary.find('.viewer-wrapper');
    var $contributorCount = $gitProjectSummary.find('.contributor-count');
    var $contributorText = $gitProjectSummary.find('.contributor-text');
    var contributorsCountUrl = $contributorCount.data('url');
  
    $gitProjectLanguages.on('click', function() {
      $statsSwitcherWrapper.toggleClass('js-lang');
    });
  
    $.ajax({
      url: contributorsCountUrl,
      method: 'GET',
      success: function(data) {
        if (data.status === 1) {
          $contributorCount.text(data.contributors_count);
          $contributorText.text('位贡献者')
        } else {
          $contributorText.text('获取失败')
        }
      }
    });
    var $tip = $('#apk-download-tip');
    if (!$tip.length) {
      return;
    }
    $tip.find('.btn-close').on('click', function () {
      $tip.slideUp();
    });
  });
  Gitee.initTabsInDropdown($('#git-project-branch').dropdown({
    fullTextSearch: true,
    onChange: function (value, text) {
      var path = $('#path').val();
      var href = ['/realwrtoff/ai_redis/tree', value, path].join('/');
      window.location.href = href;
    }
  }));
  $('#git-project-file').dropdown({ action: 'hide' });
  (function(){
    function pathAutoRender() {
      var $parent = $('#git-project-bread'),
          $child = $('#git-project-bread').children('.ui.horizontal.list'),
          mainWidth = 0;
      $child.each(function (i,item) {
        mainWidth += $(item).width()
        });
      $('.breadcrumb.path.fork-path').remove();
      if (mainWidth > 995) {
        $('#path-breadcrumb').hide();
        $parent.append('<div class="ui breadcrumb path fork-path">' + $('#path-breadcrumb').html() + '<div/>')
      } else {
        $('#path-breadcrumb').show();
      }
    }
    window.pathAutoRender = pathAutoRender;
    pathAutoRender();
  })();
</script>
<script>
  (function() {
    var $btnCopy, $input, $item, $panel, clipboard, remoteWay;
  
    $input = $('#project_clone_url');
  
    remoteWay = '';
  
    clipboard = new Clipboard('#btn-copy-clone-url');
  
    $panel = $('.git-project-download-panel');
  
    $btnCopy = $('#btn-copy-clone-url');
  
    $panel.find('.menu > .item').on('click', function() {
      var $item;
  
      $item = $(this).addClass('active');
      $item.siblings().removeClass('active');
      $input.val($item.attr('data-url'));
      return $.cookie('remote_way', $item.attr('data-type'), {
        expires: 365,
        path: '/'
      });
    });
  
    $('#btn-dl-or-clone').popup({
      popup: $panel,
      position: 'bottom right',
      on: 'click',
      lastResort: 'bottom right',
      maxSearchDepth: "10"
    });
  
    if (remoteWay) {
      $item = $panel.find('.item[data-type="' + remoteWay + '"]');
      if ($item.length === 0) {
        $item = $panel.find('.item[data-type="http"]');
      }
      if ($item.length === 0) {
        $item = $panel.find('.item[data-type="ssh"]');
      }
      if ($item.length === 0) {
        $item = $panel.find('.item[data-type="svn"]');
      }
      $item.addClass('active').siblings().removeClass('active');
      $input.val($item.attr('data-url'));
    }
  
    clipboard.on('success', function() {
      $btnCopy.popup({
        content: '已复制',
        position: 'right center',
        onHidden: function() {
          return $btnCopy.popup('destroy');
        }
      });
      return $btnCopy.popup('show');
    });
  
    clipboard.on('error', function() {
      $btnCopy.popup({
        content: '复制失败，请手动复制',
        position: 'right center',
        onHidden: function() {
          return $btnCopy.popup('destroy');
        }
      });
      return $btnCopy.popup('show');
    });
  
  }).call(this);
</script>

<div class='row column tree-holder' id='tree-holder'>
<div class='tree-content-holder' id='tree-content-holder'>
<div class='file_holder'>
<div class='file_title'>
<div class='blob-header-title'>
<div class='blob-description'>
<i class='iconfont icon-readme'></i>
<span class='file_name' title='redis_test.go'>
redis_test.go
</span>
<small>2.20 KB</small>
</div>
<div class='options'><div class='ui mini buttons basic'>
<textarea id="blob_raw" name="blob_raw" style="display:none;">
package ai_redis&#x000A;&#x000A;import (&#x000A;	&quot;testing&quot;&#x000A;)&#x000A;&#x000A;func TestAIRedis_Init(t *testing.T) {&#x000A;	aiRds := &amp;AIRedis{}&#x000A;&#x000A;	err := aiRds.Init(&quot;127.0.1&quot;, 21601, &quot;Mindata123&quot;, 0)&#x000A;	if err != nil {&#x000A;		t.Error(&quot;TestAIRedis_Init failed, err &quot;, err.Error())&#x000A;	} else {&#x000A;		t.Log(&quot;TestAIRedis_Init ok&quot;)&#x000A;	}&#x000A;	opt := aiRds.Cli.Options()&#x000A;	t.Log(opt.PoolSize)&#x000A;	t.Log(opt.PoolTimeout)&#x000A;}&#x000A;&#x000A;func TestAIRedis_Rpush(t *testing.T) {&#x000A;	aiRds := &amp;AIRedis{}&#x000A;	aiRds.Init(&quot;127.0.1&quot;, 21601, &quot;Mindata123&quot;, 0)&#x000A;	aiRds.Rpush(&quot;jim_test&quot;, &quot;1&quot;)&#x000A;	aiRds.Rpush(&quot;jim_test&quot;, &quot;1&quot;)&#x000A;	aiRds.Rpush(&quot;jim_test&quot;, &quot;1&quot;)&#x000A;	length, err := aiRds.LLen(&quot;jim_test&quot;)&#x000A;	if err != nil {&#x000A;		t.Error(&quot;TestAIRedis_Rpush failed, err &quot;, err.Error())&#x000A;	} else {&#x000A;		t.Log(&quot;queue length &quot;, length)&#x000A;	}&#x000A;	aiRds.Lpop(&quot;jim_test&quot;)&#x000A;	length, _ = aiRds.LLen(&quot;jim_test&quot;)&#x000A;	t.Log(&quot;queue length &quot;, length)&#x000A;}&#x000A;&#x000A;func TestAIRedis_SetNX(t *testing.T) {&#x000A;	aiRds := &amp;AIRedis{}&#x000A;	aiRds.Init(&quot;127.0.1&quot;, 21601, &quot;Mindata123&quot;, 0)&#x000A;	res, err := aiRds.SetNX(&quot;jim_test&quot;, &quot;1&quot;, 0)&#x000A;	if err != nil {&#x000A;		t.Error(&quot;TestAIRedis_Rpush failed, &quot;, err.Error())&#x000A;	} else {&#x000A;		t.Log(&quot;TestAIRedis_Rpush res &quot;, res)&#x000A;	}&#x000A;	aiRds.Delete(&quot;jim_test&quot;)&#x000A;	res, err = aiRds.SetNX(&quot;jim_test&quot;, &quot;1&quot;, 0)&#x000A;	if err != nil {&#x000A;		t.Error(&quot;TestRedisSetX failed, &quot;, err.Error())&#x000A;	} else {&#x000A;		t.Log(&quot;TestRedisSetX res &quot;, res)&#x000A;	}&#x000A;	res, err = aiRds.SetNX(&quot;jim_test&quot;, &quot;1&quot;, 0)&#x000A;	if err != nil {&#x000A;		t.Error(&quot;TestRedisSetX failed, &quot;, err.Error())&#x000A;	} else {&#x000A;		t.Log(&quot;TestRedisSetX res &quot;, res)&#x000A;	}&#x000A;}&#x000A;&#x000A;func TestAIRedis_Lpop(t *testing.T) {&#x000A;	aiRds := &amp;AIRedis{}&#x000A;	aiRds.Init(&quot;127.0.1&quot;, 21601, &quot;Mindata123&quot;, 0)&#x000A;	aiRds.Delete(&quot;jim_test&quot;)&#x000A;	result, err := aiRds.Lpop(&quot;jim_test&quot;)&#x000A;	if err != nil {&#x000A;		if err.Error() != &quot;redis: nil&quot; {&#x000A;			t.Errorf(&quot;TestAIRedis_Lpop failed, err [%s]&quot;, err.Error())&#x000A;		}&#x000A;	} else {&#x000A;		t.Log(&quot;TestAIRedis_Lpop res &quot;, result)&#x000A;	}&#x000A;	aiRds.Rpush(&quot;jim_test&quot;, &quot;hello&quot;)&#x000A;	result, err = aiRds.Lpop(&quot;jim_test&quot;)&#x000A;	if err != nil {&#x000A;		t.Error(&quot;TestAIRedis_Lpop failed, err &quot;, err.Error())&#x000A;	} else {&#x000A;		t.Log(&quot;TestAIRedis_Lpop res &quot;, result)&#x000A;	}&#x000A;}&#x000A;&#x000A;&#x000A;func TestAIRedis_LLen(t *testing.T) {&#x000A;	aiRds := &amp;AIRedis{}&#x000A;	aiRds.Init(&quot;127.0.1&quot;, 21601, &quot;Mindata123&quot;, 0)&#x000A;	aiRds.Delete(&quot;jim_test&quot;)&#x000A;	length, err := aiRds.LLen(&quot;jim_test&quot;)&#x000A;	if err != nil {&#x000A;		t.Error(&quot;TestAIRedis_LLen failed, err &quot;, err.Error())&#x000A;	} else {&#x000A;		t.Log(&quot;queue length &quot;, length)&#x000A;	}&#x000A;}&#x000A;</textarea>
<a href="#" class="ui button" id="copy-text" style="border-left: none;">一键复制</a>
<a href="/realwrtoff/ai_redis/edit/master/ai_redis/redis_test.go" class="ui button disabled has_tooltip edit-blob" title="无编辑权限">编辑</a>
<a href="/-/ide/project/realwrtoff/ai_redis/edit/master/-/ai_redis/redis_test.go" class="ui button web-ide" target="_blank">Web IDE</a>
<a href="/realwrtoff/ai_redis/raw/master/ai_redis/redis_test.go" class="ui button edit-raw" target="_blank">原始数据</a>
<a href="/realwrtoff/ai_redis/blame/master/ai_redis/redis_test.go" class="ui button edit-blame">按行查看</a>
<a href="/realwrtoff/ai_redis/commits/master/ai_redis/redis_test.go" class="ui button edit-history">历史</a>
</div>
<script>
  "use strict";
  try {
    if((gon.wait_fork!=undefined && gon.wait_fork==true) || (gon.wait_fetch!=undefined && gon.wait_fetch==true)){
      $('.edit-blob').popup({content:"当前项目正在后台处理中,暂时无法编辑", on: 'hover', delay: { show: 200, hide: 200 }});
      $('.edit-blob').click(function(e){
        e.preventDefault();
      })
    }
  
    var setUrl = function() {
      var params = window.location.search
      if (params==undefined || $.trim(params).length==0) return;
      $('span.options').children('.basic').find('a').each(function(index,ele){
        var origin_href = $(ele).attr('href');
        if (origin_href!="#" && origin_href.indexOf('?') == -1){
          $(ele).attr('href',origin_href+params);
        }
      });
    }
  
    setUrl();
  
    var clipboard = null,
        $btncopy  = $("#copy-text");
  
    clipboard = new Clipboard("#copy-text", {
      text: function(trigger) {
        return $("#blob_raw").val();
      }
    })
  
    clipboard.on('success', function(e) {
      $btncopy.popup('hide');
      $btncopy.popup('destroy');
      $btncopy.popup({content: '已复制', position: 'bottom center'});
      $btncopy.popup('show');
    })
  
    clipboard.on('error', function(e) {
      var giteeModal = new GiteeModalHelper({okText: '确定'});
      giteeModal.alert("一键复制", '复制失败，请手动复制');
    })
  
    $(function() {
      $btncopy.popup({
        content: '点击复制',
        position: 'bottom center'
      })
    })
  
  } catch (error) {
    console.log('blob/action error:' + error);
  }
</script>
</div>
</div>
<div class='contributor-description'><span class='recent-commit' style='margin-top: 0.7rem'>
<a class="commit-author-link" href="mailto:jim@mobvista.com">jim</a>
<span>提交于</span>
<span class='timeago commit-date' title='2018-06-23 23:34:41 +0800'>
2018-06-23 23:34
</span>
.
<a href="/realwrtoff/ai_redis/commit/8df39c870855b0f2d7815bab212888578654e555" title="init">init</a>
</span>
</div>
</div>
<div class='clearfix'></div>
<div class='file_content code'>
<div class='lines white'>
<div class='line-numbers'><a href='#L1' id='L1'>1</a><a href='#L2' id='L2'>2</a><a href='#L3' id='L3'>3</a><a href='#L4' id='L4'>4</a><a href='#L5' id='L5'>5</a><a href='#L6' id='L6'>6</a><a href='#L7' id='L7'>7</a><a href='#L8' id='L8'>8</a><a href='#L9' id='L9'>9</a><a href='#L10' id='L10'>10</a><a href='#L11' id='L11'>11</a><a href='#L12' id='L12'>12</a><a href='#L13' id='L13'>13</a><a href='#L14' id='L14'>14</a><a href='#L15' id='L15'>15</a><a href='#L16' id='L16'>16</a><a href='#L17' id='L17'>17</a><a href='#L18' id='L18'>18</a><a href='#L19' id='L19'>19</a><a href='#L20' id='L20'>20</a><a href='#L21' id='L21'>21</a><a href='#L22' id='L22'>22</a><a href='#L23' id='L23'>23</a><a href='#L24' id='L24'>24</a><a href='#L25' id='L25'>25</a><a href='#L26' id='L26'>26</a><a href='#L27' id='L27'>27</a><a href='#L28' id='L28'>28</a><a href='#L29' id='L29'>29</a><a href='#L30' id='L30'>30</a><a href='#L31' id='L31'>31</a><a href='#L32' id='L32'>32</a><a href='#L33' id='L33'>33</a><a href='#L34' id='L34'>34</a><a href='#L35' id='L35'>35</a><a href='#L36' id='L36'>36</a><a href='#L37' id='L37'>37</a><a href='#L38' id='L38'>38</a><a href='#L39' id='L39'>39</a><a href='#L40' id='L40'>40</a><a href='#L41' id='L41'>41</a><a href='#L42' id='L42'>42</a><a href='#L43' id='L43'>43</a><a href='#L44' id='L44'>44</a><a href='#L45' id='L45'>45</a><a href='#L46' id='L46'>46</a><a href='#L47' id='L47'>47</a><a href='#L48' id='L48'>48</a><a href='#L49' id='L49'>49</a><a href='#L50' id='L50'>50</a><a href='#L51' id='L51'>51</a><a href='#L52' id='L52'>52</a><a href='#L53' id='L53'>53</a><a href='#L54' id='L54'>54</a><a href='#L55' id='L55'>55</a><a href='#L56' id='L56'>56</a><a href='#L57' id='L57'>57</a><a href='#L58' id='L58'>58</a><a href='#L59' id='L59'>59</a><a href='#L60' id='L60'>60</a><a href='#L61' id='L61'>61</a><a href='#L62' id='L62'>62</a><a href='#L63' id='L63'>63</a><a href='#L64' id='L64'>64</a><a href='#L65' id='L65'>65</a><a href='#L66' id='L66'>66</a><a href='#L67' id='L67'>67</a><a href='#L68' id='L68'>68</a><a href='#L69' id='L69'>69</a><a href='#L70' id='L70'>70</a><a href='#L71' id='L71'>71</a><a href='#L72' id='L72'>72</a><a href='#L73' id='L73'>73</a><a href='#L74' id='L74'>74</a><a href='#L75' id='L75'>75</a><a href='#L76' id='L76'>76</a><a href='#L77' id='L77'>77</a><a href='#L78' id='L78'>78</a><a href='#L79' id='L79'>79</a><a href='#L80' id='L80'>80</a><a href='#L81' id='L81'>81</a><a href='#L82' id='L82'>82</a><a href='#L83' id='L83'>83</a><a href='#L84' id='L84'>84</a><a href='#L85' id='L85'>85</a><a href='#L86' id='L86'>86</a><a href='#L87' id='L87'>87</a><a href='#L88' id='L88'>88</a><a href='#L89' id='L89'>89</a><a href='#L90' id='L90'>90</a><a href='#L91' id='L91'>91</a><a href='#L92' id='L92'>92</a><a href='#L93' id='L93'>93</a><a href='#L94' id='L94'>94</a></div><div class='highlight'><pre><div class='line' id='LC1'><span class="k">package</span><span class="x"> </span><span class="n">ai_redis</span>&#x000A;</div><div class='line' id='LC2'>&#x000A;</div><div class='line' id='LC3'><span class="k">import</span><span class="x"> </span><span class="p">(</span>&#x000A;</div><div class='line' id='LC4'><span class="x">	</span><span class="s">"testing"</span>&#x000A;</div><div class='line' id='LC5'><span class="p">)</span>&#x000A;</div><div class='line' id='LC6'>&#x000A;</div><div class='line' id='LC7'><span class="k">func</span><span class="x"> </span><span class="n">TestAIRedis_Init</span><span class="p">(</span><span class="n">t</span><span class="x"> </span><span class="o">*</span><span class="n">testing</span><span class="o">.</span><span class="n">T</span><span class="p">)</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC8'><span class="x">	</span><span class="n">aiRds</span><span class="x"> </span><span class="o">:=</span><span class="x"> </span><span class="o">&amp;</span><span class="n">AIRedis</span><span class="p">{}</span>&#x000A;</div><div class='line' id='LC9'>&#x000A;</div><div class='line' id='LC10'><span class="x">	</span><span class="n">err</span><span class="x"> </span><span class="o">:=</span><span class="x"> </span><span class="n">aiRds</span><span class="o">.</span><span class="n">Init</span><span class="p">(</span><span class="s">"127.0.1"</span><span class="p">,</span><span class="x"> </span><span class="m">21601</span><span class="p">,</span><span class="x"> </span><span class="s">"Mindata123"</span><span class="p">,</span><span class="x"> </span><span class="m">0</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC11'><span class="x">	</span><span class="k">if</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">!=</span><span class="x"> </span><span class="no">nil</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC12'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Error</span><span class="p">(</span><span class="s">"TestAIRedis_Init failed, err "</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="o">.</span><span class="n">Error</span><span class="p">())</span>&#x000A;</div><div class='line' id='LC13'><span class="x">	</span><span class="p">}</span><span class="x"> </span><span class="k">else</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC14'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Log</span><span class="p">(</span><span class="s">"TestAIRedis_Init ok"</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC15'><span class="x">	</span><span class="p">}</span>&#x000A;</div><div class='line' id='LC16'><span class="x">	</span><span class="n">opt</span><span class="x"> </span><span class="o">:=</span><span class="x"> </span><span class="n">aiRds</span><span class="o">.</span><span class="n">Cli</span><span class="o">.</span><span class="n">Options</span><span class="p">()</span>&#x000A;</div><div class='line' id='LC17'><span class="x">	</span><span class="n">t</span><span class="o">.</span><span class="n">Log</span><span class="p">(</span><span class="n">opt</span><span class="o">.</span><span class="n">PoolSize</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC18'><span class="x">	</span><span class="n">t</span><span class="o">.</span><span class="n">Log</span><span class="p">(</span><span class="n">opt</span><span class="o">.</span><span class="n">PoolTimeout</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC19'><span class="p">}</span>&#x000A;</div><div class='line' id='LC20'>&#x000A;</div><div class='line' id='LC21'><span class="k">func</span><span class="x"> </span><span class="n">TestAIRedis_Rpush</span><span class="p">(</span><span class="n">t</span><span class="x"> </span><span class="o">*</span><span class="n">testing</span><span class="o">.</span><span class="n">T</span><span class="p">)</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC22'><span class="x">	</span><span class="n">aiRds</span><span class="x"> </span><span class="o">:=</span><span class="x"> </span><span class="o">&amp;</span><span class="n">AIRedis</span><span class="p">{}</span>&#x000A;</div><div class='line' id='LC23'><span class="x">	</span><span class="n">aiRds</span><span class="o">.</span><span class="n">Init</span><span class="p">(</span><span class="s">"127.0.1"</span><span class="p">,</span><span class="x"> </span><span class="m">21601</span><span class="p">,</span><span class="x"> </span><span class="s">"Mindata123"</span><span class="p">,</span><span class="x"> </span><span class="m">0</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC24'><span class="x">	</span><span class="n">aiRds</span><span class="o">.</span><span class="n">Rpush</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">,</span><span class="x"> </span><span class="s">"1"</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC25'><span class="x">	</span><span class="n">aiRds</span><span class="o">.</span><span class="n">Rpush</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">,</span><span class="x"> </span><span class="s">"1"</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC26'><span class="x">	</span><span class="n">aiRds</span><span class="o">.</span><span class="n">Rpush</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">,</span><span class="x"> </span><span class="s">"1"</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC27'><span class="x">	</span><span class="n">length</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">:=</span><span class="x"> </span><span class="n">aiRds</span><span class="o">.</span><span class="n">LLen</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC28'><span class="x">	</span><span class="k">if</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">!=</span><span class="x"> </span><span class="no">nil</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC29'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Error</span><span class="p">(</span><span class="s">"TestAIRedis_Rpush failed, err "</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="o">.</span><span class="n">Error</span><span class="p">())</span>&#x000A;</div><div class='line' id='LC30'><span class="x">	</span><span class="p">}</span><span class="x"> </span><span class="k">else</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC31'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Log</span><span class="p">(</span><span class="s">"queue length "</span><span class="p">,</span><span class="x"> </span><span class="n">length</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC32'><span class="x">	</span><span class="p">}</span>&#x000A;</div><div class='line' id='LC33'><span class="x">	</span><span class="n">aiRds</span><span class="o">.</span><span class="n">Lpop</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC34'><span class="x">	</span><span class="n">length</span><span class="p">,</span><span class="x"> </span><span class="n">_</span><span class="x"> </span><span class="o">=</span><span class="x"> </span><span class="n">aiRds</span><span class="o">.</span><span class="n">LLen</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC35'><span class="x">	</span><span class="n">t</span><span class="o">.</span><span class="n">Log</span><span class="p">(</span><span class="s">"queue length "</span><span class="p">,</span><span class="x"> </span><span class="n">length</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC36'><span class="p">}</span>&#x000A;</div><div class='line' id='LC37'>&#x000A;</div><div class='line' id='LC38'><span class="k">func</span><span class="x"> </span><span class="n">TestAIRedis_SetNX</span><span class="p">(</span><span class="n">t</span><span class="x"> </span><span class="o">*</span><span class="n">testing</span><span class="o">.</span><span class="n">T</span><span class="p">)</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC39'><span class="x">	</span><span class="n">aiRds</span><span class="x"> </span><span class="o">:=</span><span class="x"> </span><span class="o">&amp;</span><span class="n">AIRedis</span><span class="p">{}</span>&#x000A;</div><div class='line' id='LC40'><span class="x">	</span><span class="n">aiRds</span><span class="o">.</span><span class="n">Init</span><span class="p">(</span><span class="s">"127.0.1"</span><span class="p">,</span><span class="x"> </span><span class="m">21601</span><span class="p">,</span><span class="x"> </span><span class="s">"Mindata123"</span><span class="p">,</span><span class="x"> </span><span class="m">0</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC41'><span class="x">	</span><span class="n">res</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">:=</span><span class="x"> </span><span class="n">aiRds</span><span class="o">.</span><span class="n">SetNX</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">,</span><span class="x"> </span><span class="s">"1"</span><span class="p">,</span><span class="x"> </span><span class="m">0</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC42'><span class="x">	</span><span class="k">if</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">!=</span><span class="x"> </span><span class="no">nil</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC43'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Error</span><span class="p">(</span><span class="s">"TestAIRedis_Rpush failed, "</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="o">.</span><span class="n">Error</span><span class="p">())</span>&#x000A;</div><div class='line' id='LC44'><span class="x">	</span><span class="p">}</span><span class="x"> </span><span class="k">else</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC45'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Log</span><span class="p">(</span><span class="s">"TestAIRedis_Rpush res "</span><span class="p">,</span><span class="x"> </span><span class="n">res</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC46'><span class="x">	</span><span class="p">}</span>&#x000A;</div><div class='line' id='LC47'><span class="x">	</span><span class="n">aiRds</span><span class="o">.</span><span class="n">Delete</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC48'><span class="x">	</span><span class="n">res</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">=</span><span class="x"> </span><span class="n">aiRds</span><span class="o">.</span><span class="n">SetNX</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">,</span><span class="x"> </span><span class="s">"1"</span><span class="p">,</span><span class="x"> </span><span class="m">0</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC49'><span class="x">	</span><span class="k">if</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">!=</span><span class="x"> </span><span class="no">nil</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC50'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Error</span><span class="p">(</span><span class="s">"TestRedisSetX failed, "</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="o">.</span><span class="n">Error</span><span class="p">())</span>&#x000A;</div><div class='line' id='LC51'><span class="x">	</span><span class="p">}</span><span class="x"> </span><span class="k">else</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC52'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Log</span><span class="p">(</span><span class="s">"TestRedisSetX res "</span><span class="p">,</span><span class="x"> </span><span class="n">res</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC53'><span class="x">	</span><span class="p">}</span>&#x000A;</div><div class='line' id='LC54'><span class="x">	</span><span class="n">res</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">=</span><span class="x"> </span><span class="n">aiRds</span><span class="o">.</span><span class="n">SetNX</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">,</span><span class="x"> </span><span class="s">"1"</span><span class="p">,</span><span class="x"> </span><span class="m">0</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC55'><span class="x">	</span><span class="k">if</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">!=</span><span class="x"> </span><span class="no">nil</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC56'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Error</span><span class="p">(</span><span class="s">"TestRedisSetX failed, "</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="o">.</span><span class="n">Error</span><span class="p">())</span>&#x000A;</div><div class='line' id='LC57'><span class="x">	</span><span class="p">}</span><span class="x"> </span><span class="k">else</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC58'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Log</span><span class="p">(</span><span class="s">"TestRedisSetX res "</span><span class="p">,</span><span class="x"> </span><span class="n">res</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC59'><span class="x">	</span><span class="p">}</span>&#x000A;</div><div class='line' id='LC60'><span class="p">}</span>&#x000A;</div><div class='line' id='LC61'>&#x000A;</div><div class='line' id='LC62'><span class="k">func</span><span class="x"> </span><span class="n">TestAIRedis_Lpop</span><span class="p">(</span><span class="n">t</span><span class="x"> </span><span class="o">*</span><span class="n">testing</span><span class="o">.</span><span class="n">T</span><span class="p">)</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC63'><span class="x">	</span><span class="n">aiRds</span><span class="x"> </span><span class="o">:=</span><span class="x"> </span><span class="o">&amp;</span><span class="n">AIRedis</span><span class="p">{}</span>&#x000A;</div><div class='line' id='LC64'><span class="x">	</span><span class="n">aiRds</span><span class="o">.</span><span class="n">Init</span><span class="p">(</span><span class="s">"127.0.1"</span><span class="p">,</span><span class="x"> </span><span class="m">21601</span><span class="p">,</span><span class="x"> </span><span class="s">"Mindata123"</span><span class="p">,</span><span class="x"> </span><span class="m">0</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC65'><span class="x">	</span><span class="n">aiRds</span><span class="o">.</span><span class="n">Delete</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC66'><span class="x">	</span><span class="n">result</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">:=</span><span class="x"> </span><span class="n">aiRds</span><span class="o">.</span><span class="n">Lpop</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC67'><span class="x">	</span><span class="k">if</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">!=</span><span class="x"> </span><span class="no">nil</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC68'><span class="x">		</span><span class="k">if</span><span class="x"> </span><span class="n">err</span><span class="o">.</span><span class="n">Error</span><span class="p">()</span><span class="x"> </span><span class="o">!=</span><span class="x"> </span><span class="s">"redis: nil"</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC69'><span class="x">			</span><span class="n">t</span><span class="o">.</span><span class="n">Errorf</span><span class="p">(</span><span class="s">"TestAIRedis_Lpop failed, err [%s]"</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="o">.</span><span class="n">Error</span><span class="p">())</span>&#x000A;</div><div class='line' id='LC70'><span class="x">		</span><span class="p">}</span>&#x000A;</div><div class='line' id='LC71'><span class="x">	</span><span class="p">}</span><span class="x"> </span><span class="k">else</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC72'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Log</span><span class="p">(</span><span class="s">"TestAIRedis_Lpop res "</span><span class="p">,</span><span class="x"> </span><span class="n">result</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC73'><span class="x">	</span><span class="p">}</span>&#x000A;</div><div class='line' id='LC74'><span class="x">	</span><span class="n">aiRds</span><span class="o">.</span><span class="n">Rpush</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">,</span><span class="x"> </span><span class="s">"hello"</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC75'><span class="x">	</span><span class="n">result</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">=</span><span class="x"> </span><span class="n">aiRds</span><span class="o">.</span><span class="n">Lpop</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC76'><span class="x">	</span><span class="k">if</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">!=</span><span class="x"> </span><span class="no">nil</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC77'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Error</span><span class="p">(</span><span class="s">"TestAIRedis_Lpop failed, err "</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="o">.</span><span class="n">Error</span><span class="p">())</span>&#x000A;</div><div class='line' id='LC78'><span class="x">	</span><span class="p">}</span><span class="x"> </span><span class="k">else</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC79'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Log</span><span class="p">(</span><span class="s">"TestAIRedis_Lpop res "</span><span class="p">,</span><span class="x"> </span><span class="n">result</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC80'><span class="x">	</span><span class="p">}</span>&#x000A;</div><div class='line' id='LC81'><span class="p">}</span>&#x000A;</div><div class='line' id='LC82'>&#x000A;</div><div class='line' id='LC83'>&#x000A;</div><div class='line' id='LC84'><span class="k">func</span><span class="x"> </span><span class="n">TestAIRedis_LLen</span><span class="p">(</span><span class="n">t</span><span class="x"> </span><span class="o">*</span><span class="n">testing</span><span class="o">.</span><span class="n">T</span><span class="p">)</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC85'><span class="x">	</span><span class="n">aiRds</span><span class="x"> </span><span class="o">:=</span><span class="x"> </span><span class="o">&amp;</span><span class="n">AIRedis</span><span class="p">{}</span>&#x000A;</div><div class='line' id='LC86'><span class="x">	</span><span class="n">aiRds</span><span class="o">.</span><span class="n">Init</span><span class="p">(</span><span class="s">"127.0.1"</span><span class="p">,</span><span class="x"> </span><span class="m">21601</span><span class="p">,</span><span class="x"> </span><span class="s">"Mindata123"</span><span class="p">,</span><span class="x"> </span><span class="m">0</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC87'><span class="x">	</span><span class="n">aiRds</span><span class="o">.</span><span class="n">Delete</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC88'><span class="x">	</span><span class="n">length</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">:=</span><span class="x"> </span><span class="n">aiRds</span><span class="o">.</span><span class="n">LLen</span><span class="p">(</span><span class="s">"jim_test"</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC89'><span class="x">	</span><span class="k">if</span><span class="x"> </span><span class="n">err</span><span class="x"> </span><span class="o">!=</span><span class="x"> </span><span class="no">nil</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC90'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Error</span><span class="p">(</span><span class="s">"TestAIRedis_LLen failed, err "</span><span class="p">,</span><span class="x"> </span><span class="n">err</span><span class="o">.</span><span class="n">Error</span><span class="p">())</span>&#x000A;</div><div class='line' id='LC91'><span class="x">	</span><span class="p">}</span><span class="x"> </span><span class="k">else</span><span class="x"> </span><span class="p">{</span>&#x000A;</div><div class='line' id='LC92'><span class="x">		</span><span class="n">t</span><span class="o">.</span><span class="n">Log</span><span class="p">(</span><span class="s">"queue length "</span><span class="p">,</span><span class="x"> </span><span class="n">length</span><span class="p">)</span>&#x000A;</div><div class='line' id='LC93'><span class="x">	</span><span class="p">}</span>&#x000A;</div><div class='line' id='LC94'><span class="p">}</span>&#x000A;</div></pre></div></div>
</div>

</div>
</div>
<div class='tree_progress'></div>
<div class='ui small modal' id='modal-linejump'>
<div class='ui custom form content'>
<div class='field'>
<div class='ui right action input'>
<input placeholder='跳转至某一行...' type='number'>
<div class='ui orange button'>
跳转
</div>
</div>
</div>
</div>
</div>

<div class='row column inner-comment' id='blob-comment'>
<input id="comment_path" name="comment_path" type="hidden" value="ai_redis/redis_test.go" />
<div class='tree-comments'>
<h3 id='tree_comm_title'>
评论
(
<span class='comments-count'>
0
</span>
)
</h3>
<div class='ui threaded comments middle aligned' id='notes-list'></div>
<input id="ajax_add_note_id" name="ajax_add_note_id" type="hidden" />
<div class='text-center'>
<div class='tip-loading' style='display: none'>
<div class='ui active mini inline loader'></div>
正在加载...
</div>
</div>
</div>
<script>
  "use strict";
  (function(){
    var page = 1
    var commentsCount = 0
    var $container = $('.tree-comments')
    var $comments = $container.find('.ui.comments')
    var $tipLoading = $container.find('.tip-loading')
    var $btnLoad = $container.find('.btn-load-more')
  
    if (commentsCount < 1) {
      return;
    }
  
    var path;
    if ($('#comment_path').val() === '') {
      path = '/';
    } else {
      path = $('#comment_path').val();
    }
  
    function loadComments () {
      $btnLoad.hide();
      $tipLoading.show();
      $.ajax({
        url: '/realwrtoff/ai_redis/comment_list',
        data: {
          page: page,
          path: path
        },
        success: function(data) {
          var err;
          try {
            $tipLoading.hide();
            $btnLoad.show();
            if (data.status !== 0) {
              $btnLoad.text('无更多评论')
              return $btnLoad.prop('disabled', true).addClass('disabled');
            } else {
              TreeComment.CommentListHandler(data);
              page += 1;
              if (data.comments_count < 10) {
                $('#ajax_add_note_id').val('');
                $btnLoad.text('无更多评论')
                $btnLoad.prop('disabled', true).addClass('disabled');
              }
              // osctree can not load script
              $comments.find('.timeago').timeago();
              $comments.find('.commenter-role-label').popup();
              noteAnchorLocater.setup();
              toMathMlCode('', 'comments');
              return $('.markdown-body pre code').each(function(i, block) {
                return hljs.highlightBlock(block);
              });
            }
          } catch (error) {
            err = error;
            return console.log('loadComments error:' + err);
          }
        }
      });
    };
  
  
    function checkLoad () {
      var listTop, top;
      top = $(window).scrollTop();
      listTop = $container.offset().top;
      if (listTop >= top && listTop < top + $(window).height()) {
        $(window).off('scroll', checkLoad);
        return loadComments();
      }
    };
  
    $btnLoad.on('click', loadComments);
    loadComments()
  
    try {
      this.noteAnchorLocater.init({
        lastAnchorElm: '#tree_comm_title'
      });
    } catch (error) {
      console.log('noteAnchorLocater err:' + error);
    }
  })();
</script>

</div>
<div class='inner-comment-box' id='comment-box'>
<p>
你可以在<a href="/login">登录</a>后，发表评论
</p>

</div>

<script>
  "use strict";
  $('#checked_comments_with_star').checkbox('set unchecked')
</script>

</div>
</div>
</div>
</div>
<script>
  (function() {
    $(function() {
      Tree.init();
      return TreeCommentActions.init();
    });
  
  }).call(this);
</script>

<script>
  (function() {
    var donateModal;
  
    Gitee.modalHelper = new GiteeModalHelper({
      alertText: '提示',
      okText: '确定'
    });
  
    donateModal = new ProjectDonateModal({
      el: '#project-donate-modal',
      alipayUrl: '/realwrtoff/ai_redis/alipay',
      wepayUrl: '/realwrtoff/ai_redis/wepay',
      modalHelper: Gitee.modalHelper
    });
  
    if ("" === 'true') {
      donateModal.show();
    }
  
    $('#project-donate').on('click', function() {
      return donateModal.show();
    });
  
  }).call(this);
</script>
<script>
  Tree.initHighlightTheme('white')
</script>

<script>
  $(function() {
    GitLab.GfmAutoComplete.dataSource = "/realwrtoff/ai_redis/autocomplete_sources"
    GitLab.GfmAutoComplete.Emoji.assetBase = '/assets/emoji'
    GitLab.GfmAutoComplete.setup();
  });
</script>

<footer id='git-footer-main'>
<div class='ui container'>
<div class='ui two column grid'>
<div class='column'>
<p><a href="https://gitee.com/" target="_blank">© Gitee.com  </a></p>
<div class='ui three column grid' id='footer-left'>
<div class='column'>
<div class='ui link list'>
<div class='item'>
<a href="/about_us" class="item">关于我们</a>
</div>
<div class='item'>
<a href="/terms" class="item">使用条款</a>
</div>
<div class='item'>
<a href="/oschina/git-osc/issues" class="item">意见建议</a>
</div>
<div class='item'>
<a href="/links.html" class="item">合作伙伴</a>
</div>
</div>
</div>
<div class='column'>
<div class='ui link list'>
<div class='item'>
<a href="/api/v5/swagger" class="item">OpenAPI</a>
</div>
<div class='item'>
<a href="/all-about-git" class="item">Git 大全</a>
</div>
<div class='item'>
<a href="https://copycat.gitee.com/" class="item">代码克隆检测</a>
</div>
<div class='item'>
<a href="/appclient" class="item">APP与插件下载</a>
</div>
</div>
</div>
<div class='column'>
<div class='ui link list'>
<div class='item'>
<a href="https://gitee.com/help" class="item">帮助文档</a>
</div>
<div class='item'>
<a href="https://gitee.com/git-osc/" class="item">更新日志</a>
</div>
<div class='item'>
<a href="/gists" class="item">代码片段</a>
</div>
<div class='item'>
<a href="/gitee-stars" class="item">码云封面人物</a>
</div>
</div>
</div>
</div>
</div>
<div class='column right aligned followus'>
<div class='qrcode weixin'>
<img alt="Qrcode-weixin" src="/assets/qrcode-weixin-8ab7378f5545710bdb3ad5c9d17fedfe.jpg" />
<p class='weixin-text'>微信服务号</p>
</div>
<div class='phone-and-qq column'>
<div class='ui list official-support-container'>
<div class='item'>
<a href="//shang.qq.com/wpa/qunwpa?idkey=0d6c2fc0b5b71ac33405dd575bb490bf1a50e3c9a9f694e8a689cb59ee7dacc3" class="icon-popup" title="点击加入码云官方三群"><i class='iconfont icon-logo-qq'></i>
<span>
官方技术支持QQ群：655903986
</span>
</a></div>
<div class='item mail-and-zhihu'>
<a href="mailto: git@oschina.cn"><i class='iconfont icon-ic-mail'></i>
<span id='git-footer-email'>
git#oschina.cn
</span>
</a><a href="https://www.zhihu.com/org/ma-yun-osc/ " target="_blank"><i class='iconfont icon-zhihu'></i>
<span>
码云Gitee
</span>
</a></div>
<div class='item'>
企业版售前及售后使用咨询：400-606-0201
</div>
</div>
</div>
</div>
</div>
</div>
<div class='bottombar'>
<div class='ui container'>
<div class='ui grid'>
<div class='five wide column partner'>
本站带宽由 <a href="https://www.anchnet.com/" target="_blank" title="anchnet"><img alt="51idc" src="/51idc.png" /></a> 提供
</div>
<div class='eleven wide column right aligned'>
<div class='copyright'>
<a href="http://www.miitbeian.gov.cn/">粤ICP备12009483号-8</a>
深圳市奥思网络科技有限公司版权所有
</div>
<i class='icon world'></i>
<a href="/language/zh-CN">简 体
</a>/
<a href="/language/zh-TW">繁 體
</a>/
<a href="/language/en">English
</a></div>
</div>
</div>
</div>
</footer>
<script>
  Gitee.initFooter()
  $('#git-footer-main .icon-popup').popup({ position: 'bottom center' })
</script>


<div class='side-toolbar'>
<div class='button toolbar-help'>
<i class='iconfont icon-help'></i>
</div>
<div class='toolbar-help-dialog'>
<div class='toolbar-dialog-header'>
<h3 class='toolbar-dialog-title'>搜索帮助</h3>
<form accept-charset="UTF-8" action="/help/load_keywords_data" class="toolbar-help-search-form" method="get"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /></div>
<div class='ui icon input fluid toolbar-help-search'>
<input name='keywords' placeholder='请输入产品名称或问题' type='text'>
<i class='icon search'></i>
</div>
</form>

<i class='iconfont icon-close toolbar-dialog-close-icon'></i>
</div>
<div class='toolbar-dialog-content'>
<div class='toolbar-help-hot-search'>
<div class='toolbar-list'>
<div class='toolbar-list-item'>
<a href="/help/articles/4104">Git 是什么</a>
</div>
<div class='toolbar-list-item'>
<a href="/help/articles/4114">Git 仓库的基本操作</a>
</div>
<div class='toolbar-list-item'>
<a href="/help/articles/4166">企业版和个人版功能对比</a>
</div>
<div class='toolbar-list-item'>
<a href="/help/articles/4167">企业版和个人版服务对比</a>
</div>
<div class='toolbar-list-item'>
<a href="/help/articles/4194">如何处理代码冲突</a>
</div>
<div class='toolbar-list-item'>
<a href="/help/articles/4202">在小程序Web开发工具中使用Git做版本管理</a>
</div>
</div>
</div>
<div class='toolbar-help-search-reseult'>
<div class='toolbar-help-flex-column'>
<div class='ui centered inline loader toolbar-help-loader'></div>
<div class='toolbar-list'></div>
<div class='toolbar-help-link-to-help'>
<a href="" target="_blank">查看更多搜索结果</a>
</div>
</div>
</div>
</div>
</div>

<div class='button share-link'>
<i class='iconfont icon-share'></i>
</div>
<div class='ui popup'>
<div class='header'>
分享到
</div>
<div class='bdsharebuttonbox'>
<a class='item bds_weixin' data-cmd='weixin' title='分享到微信'>weixin</a>
<a class='item bds_tsina' data-cmd='tsina' title='分享到新浪微博'>sina</a>
<a class='item bds_sqq' data-cmd='sqq' title='分享到QQ好友'>qq</a>
<a class='item bds_qzone' data-cmd='qzone' title='分享到QQ空间'>qzone</a>
</div>
</div>
<div class='popup button' data-content='评论' id='home-comment'>
<i class='iconfont icon-comment'></i>
</div>
<div class='button gotop popup' data-content='回到顶部' id='gotop'>
<i class='iconfont icon-top'></i>
</div>
</div>
<script>
  Gitee.initSideToolbar({
    hasComment: true,
    commentUrl: '/realwrtoff/ai_redis#tree_comm_title'
  })
</script>
<script>
  window._bd_share_config = {
    "common": {
      "bdSnsKey": {},
      "bdText": document.title,
      "bdMini": "1",
      "bdMiniList": ["bdxc","tqf","douban","bdhome","sqq","thx","ibaidu","meilishuo","mogujie","diandian","huaban","duitang","hx","fx","youdao","sdo","qingbiji","people","xinhua","mail","isohu","yaolan","wealink","ty","iguba","fbook","twi","linkedin","h163","evernotecn","copy","print"],
      "bdPic": "",
      "bdStyle": "1",
      "bdSize": "32"
    },
    "share": {}
  }
</script>
<script src="/bd_share/static/api/js/share.js" type="text/javascript"></script>


<style>
  .float-left-box{display:none;position:fixed;left:0;bottom:0;z-index:99}.float-left-box .close-left{position:absolute;top:20px;left:25px;cursor:pointer}.float-left-box .float-people{width:160px;padding:10px}
</style>
<div class='float-left-box'>
<a href='/gitee-stars/12' target='_blank'>
<img alt="12_float_left_people" class="float-people" src="/assets/gitee_stars/12_float_left_people-c59aee86c0ba52d3b58fff9f73f16929.png" />
<img alt="12_float_left_close" class="close-left" src="/assets/gitee_stars/12_float_left_close-983603b0c61d3148acff0a5e1efb3646.png" />
</a>
</div>
<script>
  var giteeStarsWidget = $('.float-left-box')
  if ($.cookie('visit-gitee-12') == 1) {
    giteeStarsWidget.hide()
  } else {
    giteeStarsWidget.show()
  }
  $('.close-left').on('click', function(e) {
    e.preventDefault()
    $.cookie('visit-gitee-12', 1, { path: '/', expires: 30})
    giteeStarsWidget.hide()
  })
</script>

<script>
  (function() {
    this.__gac = {
      domain: 'www.oschina.net'
    };
  
  }).call(this);
</script>
<script defer src='//www.oschina.net/public/javascripts/cjl/ga.js?t=20160926' type='text/javascript'></script>

</body>
</html>
