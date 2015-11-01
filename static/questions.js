$(document).ready(function() {
	var lines = [];
	window.setInterval(function() {
		$.get("/questions/raw", function(data) {
			$("#prog2").slideUp();
			var new_lines = data.split("\n");
			for(var x = 0; x < new_lines.length; x++) {
				var line = new_lines[x];
				if (line.length > 0) {
					if (lines.indexOf(line) == -1) {
						var $newcard = $('<div class="q-card mdl-card mdl-shadow--2dp"><div class="mdl-card__supporting-text">'
								+ line + '</div></div>')
						$newcard.appendTo($(".cards"));
						lines.push(line);
					}
				}
			}
		});
	}, 500);
});
