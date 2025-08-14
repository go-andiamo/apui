function removeQueryParam(evt) {
    const p = evt.target.parentElement.parentElement;
    document.getElementById('qps').deleteRow(p.rowIndex);
}
function addQueryParam() {
    const qpName = document.getElementById("qps-select").value;
    let table = document.getElementById('qps');
    let row = table.insertRow(-1);
    let cell1 = document.createElement('th');
    cell1.textContent = qpName;
    row.appendChild(cell1);
    let cell2 = row.insertCell(-1);
    cell2.innerHTML = '<input/>';
    let cell3 = row.insertCell(-1);
    cell3.innerHTML = '<button onclick="(e => removeQueryParam(e))(event)">-</button>';
    row.querySelector('input').focus();
}
function queryParamsGet(path) {
    let inps = document.querySelectorAll('table.qps input');
    let params = [];
    inps.forEach(el => {
        if (el.value !== '') {
            params.push(el.name+'='+encodeURIComponent(el.value));
        }
    })
    if (params.length > 0) {
        window.location.href = path+'?'+params.join('&');
    } else {
        window.location.href = path;
    }
}
