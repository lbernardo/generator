package internal

import (
	"{{.package}}/internal/configs"
	"{{.package}}/internal/database"
	"{{.package}}/internal/http"
	"{{.package}}/internal/logger"
	"go.uber.org/fx"
)

var (
	ApplicationModule = fx.Options(
		// Infra
		logger.Module,
		configs.Module,
		http.Module,
		database.Module,
		// Application
	)
)
