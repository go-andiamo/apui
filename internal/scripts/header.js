function themeSelect(evt) {
    document.body.classList.forEach(cls => {
        if (cls.startsWith("theme-")) {
            document.body.classList.remove(cls);
        }
    });
    document.body.classList.add(evt.target.value);
}