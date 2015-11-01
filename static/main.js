$(document).ready(function() {
	$(".q-form").submit(function() {
		$("#prog").stop().slideDown();
		$.post("/submit", {
			"frage": $(".q-form .q-input").val()
		}, function(data) {
			if (data.indexOf("Vielen Dank") > -1) {
				$(".q-form .q-input").val("");
			}
			$(".msg-card strong").html(data);
			$(".msg-card").stop().slideDown();
			$("#prog").stop().slideUp();
		});
		return false;
	});
	var hidemsg = function() {
		$(".msg-card").stop().slideUp();
	}

	$(".q-form .q-input").change(hidemsg).keydown(hidemsg);
});
