$(function() {

	// a-zA-Z0-9, 文字数など
	$(".name").bind("keyup", function() {
		if ($(this).val().length < 1) {
			$(this).parent("tr").append("<td>名前を入力してください。</td>")
		}
	});

	$(".address").bind("keyup", function() {
		if (!$(this).val().match(/^([a-zA-Z0-9])+([a-zA-Z0-9\._-])*@([a-zA-Z0-9_-])+([a-zA-Z0-9\._-]+)+$/)){
			$(this).parent("tr").append("<td>アドレスが正しくありません。</td>");
		}
	});

	$(".password").bind("keyup", function() {
		if (($(this).val().length < 1) || ($(this).val().length > 50)) {
			$(this).parent("tr").append("<td>パスワードは1文字以上、50文字以内にしてください。</td>")
		}
	});

});
