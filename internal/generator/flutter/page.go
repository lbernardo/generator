package flutter

import (
	"fmt"
	"gitlab.com/playup/generator/internal/templates/flutter"
	"gitlab.com/playup/generator/pkg/strings"
	"gitlab.com/playup/generator/pkg/template_gen"
	"os"
	"path"
)

func NewPage(params map[string]string, rootPath string) {
	var name = params["name"]
	var feature = params["feature"]
	var featurePath = path.Join(rootPath, "lib", "app", "features", feature)
	var filePage = path.Join(featurePath, fmt.Sprintf("%v_page.dart", name))

	if err := template_gen.WriteByTemplate(flutter.Page, filePage, map[string]string{
		"ClassName": strings.ToPascalCase(name),
		"name":      name,
	}); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("created page", filePage)
	NewController(params, rootPath)
}
