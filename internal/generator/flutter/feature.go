package flutter

import (
	"fmt"
	"github.com/lbernardo/generator/internal/templates/flutter"
	"github.com/lbernardo/generator/pkg/os_gen"
	"github.com/lbernardo/generator/pkg/strings"
	"github.com/lbernardo/generator/pkg/template_gen"
	"os"
	"path"
)

func NewFeature(params map[string]string, rootPath string) {
	var name = params["name"]
	var child = params["child"]
	var featurePath = path.Join(rootPath, "lib", "app", "features", name)
	if child != "" {
		featurePath = path.Join(rootPath, "lib", "app", "features", child, "features", name)
	}
	os_gen.Mkdir(featurePath)
	if err := template_gen.WriteByTemplate(flutter.Module, path.Join(featurePath, fmt.Sprintf("%v_module.dart", name)), map[string]string{
		"ClassName": strings.ToPascalCase(name),
	}); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Created new module")

	var paramsToPageAndController = map[string]string{
		"name":    name,
		"feature": name,
	}
	if child != "" {
		paramsToPageAndController["feature"] = path.Join(child, "features", name)
	}

	NewPage(paramsToPageAndController, rootPath)
}
