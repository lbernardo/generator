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

//go:embed app_module.tpl
var AppModule string

//go:embed app_widget.tpl
var AppWidget string

//go:embed main.tpl
var Main string
