package testdata

import (
	"embed"
)

//go:embed html/*.html
var HtmlFs embed.FS
