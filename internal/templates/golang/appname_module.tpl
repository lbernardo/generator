package {{.name}}

import (
	"{{.pkg}}/internal/app/{{.name}}/http"
	"{{.pkg}}/internal/app/{{.name}}/service"
	"go.uber.org/fx"
)

var Module = fx.Module("{{.name}}", fx.Options(
	http.Module,
	fx.Provide(
		service.New{{.className}}Service,
	),
))
