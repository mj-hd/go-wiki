$(function() {

	$("textarea").bind("keyup", function() {
		updatePreview();
	});

	updatePreview();
});

function updatePreview() {
	$.ajax({
		type: "GET",
		url: "/api/markdown/",
		dataType: 'json',
		data: { text: $(this).val() },
		success: function (json) {
			$(".preview").html(json["Markdown"]);
		}
	});

}
