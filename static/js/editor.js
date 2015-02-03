$(function() {

	$("textarea").bind("keyup", function() {
		updatePreview(this);
	});

	updatePreview($("textarea"));
});

function updatePreview(elm) {
	$.ajax({
		type: "GET",
		url: "/api/markdown/",
		dataType: 'json',
		data: { text: $(elm).val() },
		success: function (json) {
			$(".preview").html(json["Markdown"]);
		}
	});

}
