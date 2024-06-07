package flutter

import (
	"fmt"
	"github.com/lbernardo/generator/internal/templates/flutter"
	"github.com/lbernardo/generator/pkg/execute"
	"github.com/lbernardo/generator/pkg/os_gen"
	"github.com/lbernardo/generator/pkg/strings"
	"github.com/lbernardo/generator/pkg/template_gen"
	"os"
	"path"
)

func NewController(params map[string]string, rootPath string) {
	var name = params["name"]
	var pathFromLibApp = params["path"]
	var featurePath = path.Join(rootPath, "lib", "app", pathFromLibApp, "controllers")
	os_gen.Mkdir(featurePath)
	var fileController = path.Join(featurePath, fmt.Sprintf("%v_controller.dart", name))
	if err := template_gen.WriteByTemplate(flutter.Controller, fileController, map[string]string{
		"ClassName": strings.ToPascalCase(name),
		"name":      name,
	}); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	cm := execute.Execute("flutter", "pub", "run", "build_runner", "build")
	cm.Dir = rootPath
	if err := cm.Run(); err != nil {
		fmt.Println("warn: error to run build_runner")
	}
	fmt.Println("create controller", fileController)
}
