$(function() {


	$("form").submit(function() {

		var hasErrorOccured = false;

		$(".error").remove()

		// a-zA-Z0-9, 文字数など
		if ($(".name").val().length < 1) {
			$(".name").parent().append("<td class='error'>名前を入力してください。</td>");
			hasErrorOccured = true;
		}	
		if (!$(".address").val().match(/^([a-zA-Z0-9])+([a-zA-Z0-9\._-])*@([a-zA-Z0-9_-])+([a-zA-Z0-9\._-]+)+$/)){
			$(".address").parent().append("<td class='error'>アドレスが正しくありません。</td>");
			hasErrorOccured = true;
		}	
		if (($(".password").val().length < 1) || ($(this).val().length > 50)) {
			$(".password").parent().append("<td class='error'>パスワードは1文字以上、50文字以内にしてください。</td>");
			hasErrorOccured = true;
		}
		return !hasErrorOccured;
	});

});
