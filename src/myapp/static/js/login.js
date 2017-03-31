(function($) {
	if(getCookie("isAuthor") === "yes") {
		window.location.href="/";
	}
	$("#submit").click(function() {
		var data = {
			"username": $("input[name='username']").val(),
			"password": $("input[name='password']").val()
		};
		console.log(data);
		$.ajax({
			url: "/doLogin",
			type: "POST",
			data: data,
			success: function(msg) {
				if(msg === "登录成功!") {
					window.location.href="/";
				} else {
					alert(msg);
				}
			}
		});
	});
})(jQuery);
