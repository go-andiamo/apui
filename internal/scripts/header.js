(function() {
    const theme = localStorage.getItem("theme");
    if (theme) {
        document.addEventListener("DOMContentLoaded", function () {
            const select = document.getElementById("theme-select");
            if (select) {
                select.value = theme;
            }
        });
    }
})();

function themeSelect(evt) {
    setTheme(evt.target.value);
    localStorage.setItem("theme", evt.target.value);
}