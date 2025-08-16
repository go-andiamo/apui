function focusing(evt) {
    let details = evt.target.closest("details");
    if (!details) {
        document.querySelectorAll("details[open]:not(.keep)").forEach(el => {
            el.removeAttribute("open");
        });
    }
}