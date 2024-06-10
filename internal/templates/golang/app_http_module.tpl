package http

import (
	"go.uber.org/fx"
)

var Module = fx.Module("{{.name}}_http",
	fx.Provide(
	    // Handlers
	),
	fx.Invoke(
		RegisterRoutes,
	))
