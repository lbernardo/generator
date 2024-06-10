package golang

import (
	"fmt"
	"github.com/lbernardo/generator/internal/templates/golang"
	"github.com/lbernardo/generator/pkg/modfile"
	"github.com/lbernardo/generator/pkg/os_gen"
	"github.com/lbernardo/generator/pkg/strings"
	"github.com/lbernardo/generator/pkg/template_gen"
	"os"
	"path"
)

func NewApp(params map[string]string, rootPath string) {
	var name = params["name"]

	module, err := modfile.GetModule(path.Join(rootPath, "go.mod"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	appRootPath := path.Join(rootPath, "internal", "app", name)

	dirs := []string{
		"entity",
		"http/handlers",
		"repository",
		"service",
	}

	for _, d := range dirs {
		name := appRootPath + "/" + d
		os_gen.Mkdir(name)
		fmt.Println("created", name)
	}

	files := map[string]string{
		fmt.Sprintf("%v_module.go", name): golang.AppNameModule,
		"http/route.go":                   golang.AppHttpRoute,
		"http/http_module.go":             golang.AppHttpModule,
		"service/service.go":              golang.AppService,
	}

	for file, content := range files {
		var filename = path.Join(appRootPath, file)
		if err := template_gen.WriteByTemplate(content, filename, map[string]string{
			"name":      name,
			"className": strings.ToPascalCase(name),
			"pkg":       module,
		}); err != nil {
			fmt.Println("error to create", filename)
			continue
		}
		fmt.Println("created", filename)
	}
}
