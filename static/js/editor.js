$(function() {

	$("textarea").bind("keyup", function() {
		$.ajax({
			type: "GET",
			url: "/api/markdown/",
			dataType: 'json',
			data: { text: $(this).val() },
			success: function (json) {
				$(".preview").html(json["Markdown"]);
			}
		});
	});

});
