package main

import (
	"{{.package}}/internal"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		internal.ApplicationModule,
	)
	app.Run()
}
