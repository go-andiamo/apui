package styling

import (
	"github.com/go-andiamo/aitch/html"
)

const baseCss = `
html, body {
	height: 100%;
	margin: 0;
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
	padding: 0.25em;
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
	font-family: var(--nav-font-family,sans-serif),sans-serif;
	font-size: var(--nav-font-size);
	color: var(--nav-text-color);
	background-color: var(--nav-bg-color);
	border-bottom: 1px solid var(--nav-border-color);
}
footer {
	font-family: var(--ftr-font-family,sans-serif),sans-serif;
	font-size: var(--ftr-font-size);
	color: var(--ftr-text-color);
	background-color: var(--ftr-bg-color);
	border-top: 1px solid var(--ftr-border-color);
}
main {
	flex: 1 1 auto;
	overflow-y: auto;
	padding: 0.25em;
	font-family: var(--main-font-family,sans-serif),sans-serif;
	font-size: var(--main-font-size);
	color: var(--main-text-color);
	background-color: var(--main-bg-color);
}
`

var BaseCssNode = html.StyleElement([]byte(baseCss))
