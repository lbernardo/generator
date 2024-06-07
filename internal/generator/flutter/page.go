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

func NewPage(params map[string]string, rootPath string) {
	var name = params["name"]
	var pathFromLibApp = params["path"]
	var featurePath = path.Join(rootPath, "lib", "app", pathFromLibApp, "pages")
	os_gen.Mkdir(featurePath)
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
