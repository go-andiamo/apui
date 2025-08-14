function toggleDetails(evt) {
    if (evt.newState === "open") {
        document.querySelectorAll('header.navigation details')
            .forEach(el => {
                if (el !== evt.target) {
                    el.removeAttribute('open');
                }
            });
        const focusable = evt.target.querySelector("input,button.select");
        if (focusable) {
            focusable.focus();
        }
    }
}
