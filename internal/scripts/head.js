function focusing(evt) {
    if (!evt.target.classList.contains("remove")) {
        let opened = document.querySelectorAll("details[open]:not(.keep)");
        let details = evt.target.closest("details:not(.keep)");
        if (!details && opened.length) {
            opened.forEach(el => {
                el.removeAttribute("open");
            });
        }

        /*
                let opened = document.querySelectorAll("details[open]:not(.keep)");
                if (opened.length) {
                    console.log("focusing", evt, opened);
                }
                let details = evt.target.closest("details");
                if (!details) {
                    document.querySelectorAll("details[open]:not(.keep)").forEach(el => {
                        el.removeAttribute("open");
                    });
                }
         */
    }
}
function toggleDetails(evt) {
    if (evt.newState === "open" && !evt.target.classList.contains("keep")) {
        document.querySelectorAll('details[open]:not(.keep)')
            .forEach(el => {
                if (el !== evt.target) {
                    el.removeAttribute('open');
                }
            });
        const focusable = evt.target.querySelector("input,button,select");
        if (focusable) {
            focusable.focus();
        }
    }
}
