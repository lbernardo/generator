package flutter

import (
	_ "embed"
)

//go:embed module.tpl
var Module string

//go:embed controller.tpl
var Controller string

//go:embed page.tpl
var Page string
