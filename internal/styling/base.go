package styling

import (
	"github.com/go-andiamo/aitch/html"
)

const baseCss = `
html, body {
	height: 100%;
	margin: 0;
}
table {
	border-collapse: collapse;
}
.lr {
    margin-top: 4px;
    float: left;
    width: 100%;
    text-align: right;
}
body {
	display: flex;
	flex-direction: column;
	font-family: var(--body-font-family,sans-serif),sans-serif;
	font-size: var(--body-font-size);
	color: var(--body-text-color);
	background-color: var(--body-bg-color);
}
header, footer {
	padding: 0 4px 0 4px;
	flex: 0 0 auto;
}
header.header {
	font-family: var(--hdr-font-family,sans-serif),sans-serif;
	font-size: var(--hdr-font-size);
	color: var(--hdr-text-color);
	background-color: var(--hdr-bg-color);
	border-bottom: 1px solid var(--hdr-border-color);
}
header.navigation {
	padding: 4px 4px 4px 4px;
	font-family: var(--nav-font-family,sans-serif),sans-serif;
	font-size: var(--nav-font-size);
	color: var(--nav-text-color);
	background-color: var(--nav-bg-color);
	border-bottom: 1px solid var(--nav-border-color);
	position: relative;
	overflow: visible;
}
header.navigation details {
	background-color: inherit;
    border: 1px solid var(--nav-border-color);
    border-radius: 4px;
    padding: 2px 4px 2px 4px;
}
header.navigation details[open] {
	color: var(--nav-opened-text-color, var(--nav-text-color));
	background-color: var(--nav-opened-bg-color, var(--nav-bg-color));
    border: 1px solid black;
    border-radius: 4px 4px 0 0;
    z-index: 1000;
}
header.navigation details .content {
    position: absolute;
	border-radius: 0 4px 4px 4px;
    border: 1px solid var(--nav-border-color);
    background-color: var(--nav-bg-color);
	color: var(--nav-text-color);
    margin-left: -5px;
	margin-top: 2px;
	padding: 4px;
}
details summary {
	cursor: pointer;
}
footer {
	font-family: var(--ftr-font-family,sans-serif),sans-serif;
	font-size: var(--ftr-font-size);
	color: var(--ftr-text-color);
	background-color: var(--ftr-bg-color);
	border-top: 1px solid var(--ftr-border-color);
	padding: 4px 4px 4px 4px;
}
main {
	flex: 1 1 auto;
	overflow: auto;
	padding: 4px 4px 4px 4px;
	font-family: var(--main-font-family,sans-serif),sans-serif;
	font-size: var(--main-font-size);
	color: var(--main-text-color);
	background-color: var(--main-bg-color);
}
.small {
	font-size: small;
}
.github {
	background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' height='32' aria-hidden='true' viewBox='0 0 24 24' version='1.1' width='32'%3E%3Cpath d='M12 1C5.923 1 1 5.923 1 12c0 4.867 3.149 8.979 7.521 10.436.55.096.756-.233.756-.522 0-.262-.013-1.128-.013-2.049-2.764.509-3.479-.674-3.699-1.292-.124-.317-.66-1.293-1.127-1.554-.385-.207-.936-.715-.014-.729.866-.014 1.485.797 1.691 1.128.99 1.663 2.571 1.196 3.204.907.096-.715.385-1.196.701-1.471-2.448-.275-5.005-1.224-5.005-5.432 0-1.196.426-2.186 1.128-2.956-.111-.275-.496-1.402.11-2.915 0 0 .921-.288 3.024 1.128a10.193 10.193 0 0 1 2.75-.371c.936 0 1.871.123 2.75.371 2.104-1.43 3.025-1.128 3.025-1.128.605 1.513.221 2.64.111 2.915.701.77 1.127 1.747 1.127 2.956 0 4.222-2.571 5.157-5.019 5.432.399.344.743 1.004.743 2.035 0 1.471-.014 2.654-.014 3.025 0 .289.206.632.756.522C19.851 20.979 23 16.854 23 12c0-6.077-4.922-11-11-11Z'%3E%3C/path%3E%3C/svg%3E");
	background-repeat: no-repeat;
	background-size: contain;
	width: 1.1em;
	height: 1.1em;
	display: inline-block;
	vertical-align: middle;
}
.method {
	color: var(--methods-text-color);
	background-color: var(--methods-bg-color);
	border: 1px solid var(--methods-border-color);
	font-family: var(--methods-font-family,sans-serif),sans-serif;
	padding: 1px 4px;
	border-radius: 4px;
	margin-right: 4px;
}
button.method {
	margin-right: 0;
}
.method.get {
	color: var(--methods-get-text-color);
	background-color: var(--methods-get-bg-color);
	border: 1px solid var(--methods-get-border-color);
}
.method.delete {
	color: var(--methods-delete-text-color);
	background-color: var(--methods-delete-bg-color);
	border: 1px solid var(--methods-delete-border-color);
}
.method.put {
	color: var(--methods-put-text-color);
	background-color: var(--methods-put-bg-color);
	border: 1px solid var(--methods-put-border-color);
}
.method.post {
	color: var(--methods-post-text-color);
	background-color: var(--methods-post-bg-color);
	border: 1px solid var(--methods-post-border-color);
}
.method.patch {
	color: var(--methods-patch-text-color);
	background-color: var(--methods-patch-bg-color);
	border: 1px solid var(--methods-patch-border-color);
}
.method.options {
	color: var(--methods-options-text-color);
	background-color: var(--methods-options-bg-color);
	border: 1px solid var(--methods-options-border-color);
}
h1,h2,h3,h4,h5,h6 {
	margin-block-start: 0.5rem;
	margin-block-end: 0.5rem;
}
.description {
	font-style: italic;
	margin-left: 1em;
    display: inline-block;
    max-width: 15em;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    vertical-align: text-bottom;
}
.inline-dropdown {
	margin-left: 1em;
	display: inline-flex;
}
` + queryParamsCss

const queryParamsCss = `
details.qps input, details.qps button, details.qps select {
	vertical-align: middle;
}
table.qps th, table.qps td {
    border: 1px solid var(--nav-border-color);
    padding: 2px;
    text-wrap: nowrap;
}
table.qps th {
	text-align: right;
}
table.qps td input {
    width: 20em;
    border: none;
    font-size: 100%;
    min-height: 1.5em;
    padding: 0 3px;
}
.navigation details.qps select {
    font-weight: bold;
    font-size: 100%;
}
`

var BaseCssNode = html.StyleElement([]byte(baseCss))
