package {{.name}}

import (
	"go.uber.org/fx"
	"{{.pkg}}/internal/app/{{.name}}/http"
	"{{.pkg}}/internal/app/{{.name}}/repository"
	"{{.pkg}}/internal/app/{{.name}}/service"
)

var Module = fx.Module("{{.name}}", fx.Options(
	http.Module,
	fx.Provide(
		service.New{{.className}}Service,
		repository.New{{.className}}Repository,
	),
))
