package flutter

import (
	"fmt"
	"github.com/lbernardo/generator/internal/templates/flutter"
	"github.com/lbernardo/generator/pkg/strings"
	"github.com/lbernardo/generator/pkg/template_gen"
	"os"
	"path"
)

func NewController(params map[string]string, rootPath string) {
	var name = params["name"]
	var feature = params["feature"]
	var featurePath = path.Join(rootPath, "lib", "app", "features", feature)
	var fileController = path.Join(featurePath, fmt.Sprintf("%v_controller.dart", name))
	if err := template_gen.WriteByTemplate(flutter.Controller, fileController, map[string]string{
		"ClassName": strings.ToPascalCase(name),
		"name":      name,
	}); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("create controller", fileController)
}
