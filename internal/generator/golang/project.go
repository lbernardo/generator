package golang

import (
	"fmt"
	"github.com/lbernardo/generator/internal/templates/golang"
	"github.com/lbernardo/generator/pkg/execute"
	"github.com/lbernardo/generator/pkg/os_gen"
	"github.com/lbernardo/generator/pkg/template_gen"
	"path"
)

func NewProject(params map[string]string, rootPath string) {
	var name = params["name"]
	var pkg = params["pkg"]

	dirs := []string{
		"cmd/server",
		"config",
		"internal/app",
		"internal/database",
		"internal/configs",
		"internal/http/middleware",
		"internal/logger",
	}

	for _, d := range dirs {
		os_gen.Mkdir(path.Join(name, d))
	}

	cm := execute.Execute("go", "mod", "init", pkg)
	cm.Dir = name
	cm.Run()

	packages := []string{
		"github.com/spf13/viper",
		"go.uber.org/zap",
		"go.uber.org/fx",
		"go.mongodb.org/mongo-driver",
		"github.com/gin-gonic/gin",
		"github.com/gin-contrib/cors",
	}

	for _, p := range packages {
		cm := execute.Execute("go", "get", p)
		cm.Dir = name
		cm.Run()
	}

	files := map[string]string{
		"cmd/server/main.go":                golang.Main,
		"internal/app_module.go":            golang.AppModule,
		"config/config.yaml":                golang.ConfigYaml,
		"internal/http/module.go":           golang.HttpModule,
		"internal/http/server.go":           golang.HttpServer,
		"internal/logger/module.go":         golang.LoggerModuler,
		"internal/configs/config_module.go": golang.ConfigModule,
		"internal/database/connection.go":   golang.DatabaseConnection,
		"internal/database/module.go":       golang.DatabaseModule,
	}

	for file, content := range files {
		var filename = path.Join(name, file)
		if err := template_gen.WriteByTemplate(content, filename, map[string]string{
			"name":    name,
			"package": pkg,
		}); err != nil {
			fmt.Println("error to create", filename)
			continue
		}
		fmt.Println("created", filename)
	}

}
