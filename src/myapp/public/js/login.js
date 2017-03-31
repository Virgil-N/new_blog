!function($){"yes"===getCookie("isAuthor")&&(window.location.href="/"),$("#submit").click(function(){var data={username:$("input[name='username']").val(),password:$("input[name='password']").val()};console.log(data),$.ajax({url:"/doLogin",type:"POST",data:data,success:function(msg){"登录成功!"===msg?window.location.href="/":alert(msg)}})})}(jQuery);
//# sourceMappingURL=login.js.map
