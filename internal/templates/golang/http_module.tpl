package http

import (
	"go.uber.org/fx"
)

var Module = fx.Module("http_server",
	fx.Provide(
		NewServer,
	),
	fx.Invoke(Start),
)
