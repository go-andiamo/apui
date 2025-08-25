window.addEventListener("load", function () {
   let t = document.querySelector("header.navigation details.ams select");
   if (t) {
       t.dispatchEvent(new Event('change'));
   }
});
function methodSelect(evt) {
    const selected = evt.target.value;
    let i = 0;
    document.querySelectorAll("header.navigation details.ams .ams-method").forEach(el => {
        if (i == selected) {
            el.classList.add("selected");
        } else {
            el.classList.remove("selected");
        }
        i++;
    });
    methodReset();
}
function methodReset() {
    const methodDiv = document.querySelector("header.navigation details.ams div.ams-method.selected");
    methodDiv.classList.remove("fetching")
    methodDiv.classList.remove("response")
    methodDiv.classList.add("send");
}
function methodExec(method, path) {
    const methodDiv = document.querySelector("header.navigation details.ams div.ams-method.selected");
    const body = methodDiv.querySelector("pre");
    const status = methodDiv.querySelector(".status");
    const statusText = methodDiv.querySelector(".status-text");
    status.classList.remove("x1xx", "x2xx", "x3xx", "x4xx", "x5xx");
    methodDiv.classList.remove("send");
    methodDiv.classList.add("fetching")
    const req = new XMLHttpRequest();
    req.open(method, path, false);
    req.setRequestHeader("Content-Type", "application/json");
    req.onreadystatechange = () => {
        if (req.readyState === XMLHttpRequest.DONE) {
            status.innerText = ""+req.status;
            statusText.innerText = ""+req.statusText;
            methodDiv.classList.remove("fetching");
            methodDiv.classList.add("response");
            status.innerText = "" + req.status;
            let sClass = (req.status - (req.status % 100)) / 100;
            if (sClass > 0 && sClass <= 5) {
                status.classList.add("x"+sClass+"xx");
            }
            if (req.status < 200 || req.status >= 300) {
                console.error(method, path, req.status, req.responseText);
            }
        }
    }
    let payload = null;
    if (body) {
        payload = body.innerText;
    }
    req.send(payload);
}