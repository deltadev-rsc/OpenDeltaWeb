function scroll() {
	document.addEventListener('DOMContentLoaded', function() {
		const anchors = document.querySelectorAll('a[href^="#"]');

		anchors.forEach(anchor => {
			anchor.addEventListener('click', function(e) {
				e.preventDefault();

				const targetId = this.getAttribute('href');
				if (targetId === '#') {
					return;
				}

				const targetElement = document.querySelector(targetId);
				if (targetElement) {
					targetElement.scrollIntoView({
						behavior: 'smooth',
						block: 'start'
					});
				}
			});
		});
	});
}

scroll();
