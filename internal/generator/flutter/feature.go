package flutter

import (
	"fmt"
	"gitlab.com/playup/generator/internal/templates/flutter"
	"gitlab.com/playup/generator/pkg/os_gen"
	"gitlab.com/playup/generator/pkg/strings"
	"gitlab.com/playup/generator/pkg/template_gen"
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
}
