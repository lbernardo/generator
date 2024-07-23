package golang

import (
	_ "embed"
)

//go:embed main.tpl
var Main string

//go:embed app_module.tpl
var AppModule string

//go:embed config.yaml.tpl
var ConfigYaml string

//go:embed http_module.tpl
var HttpModule string

//go:embed http_server.tpl
var HttpServer string

//go:embed logger_module.tpl
var LoggerModuler string

//go:embed config_module.tpl
var ConfigModule string

//go:embed app_http_module.tpl
var AppHttpModule string

//go:embed app_http_route.tpl
var AppHttpRoute string

//go:embed app_service.tpl
var AppService string

//go:embed appname_module.tpl
var AppNameModule string

//go:embed repository.tpl
var Repository string

//go:embed entity.tpl
var Entity string

//go:embed database_connection.tpl
var DatabaseConnection string

//go:embed database_module.tpl
var DatabaseModule string
