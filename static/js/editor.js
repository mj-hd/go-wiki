$(function() {

	var draft = $("textarea");
	var preview = $(".preview");

	draft.bind("keyup", function() {
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
			$(".preview").html(replaceImagePad(json["Markdown"]));
			applyDropZoneToImagePad();
		}
	});

}

function applyDropZoneToImagePad() {
	$(".imagePad").filedrop({
		paramname: "file",

		maxfiles: 1,
		maxfilesize: 1,
		url: "/api/fileUpload/",

		data: {
			page: encodeURIComponent($("input[name=OldTitle]").val())
		},

		uploadFinished: function(i, file, response) {
			$("textarea").val($("textarea").val().replace(/!\[(.*)\]\(\)/, '![$1](/api/fileView/'+encodeURIComponent($("input[name=OldTitle]").val())+'/'+encodeURIComponent(file.name)+')'));
			updatePreview($("textarea"));
		},

		error: function(err, file) {
			var message = $(".message");
			switch(err) {
				case "BrowserNotSupported":
					message.text("非対応のブラウザです。");
					break;
				case "TooManyFiles":
					message.text("ファイルは一つだけドロップしてください。");
					break;
				case "FileTooLarge":
					message.text("ファイルサイズが大きすぎます。(1MBまで)");
					break;
				default:
					break;
			}
		},

		beforeEach: function(file) {
			if (!file.type.match(/^image\//)) {
				$(".message").text("画像ファイルではありません。");

				return false;
			}
		},

		uploadStarted: function(i, file, len) {
			applyImage(this, file);
		},

		progressUpdated: function(i, file, progress) {
			$(".message").text("アップロード中…"+progress+"%");
		}
	});

}

function applyImage(pad, file) {

	//TODO
}

function replaceImagePad(html) {
	return html.replace(
			/!\[.*\]\(\)/, 
			'<div class="imagePad"><span class="message">ここに画像をドロップしてください。</span></div>'
			)
}
