package flutter

import (
	"fmt"
	"github.com/lbernardo/generator/internal/templates/flutter"
	"github.com/lbernardo/generator/pkg/execute"
	"github.com/lbernardo/generator/pkg/os_gen"
	"github.com/lbernardo/generator/pkg/template_gen"
	"os"
	"path"
)

func NewProject(params map[string]string, rootPath string) {
	var name = params["name"]
	var org = params["org"]
	cm := execute.Execute("flutter", "create", name, "--project-name", name, "--org", org, "--platforms", "android,ios")
	if err := cm.Run(); err != nil {
		fmt.Println("error create flutter project:", err)
		os.Exit(1)
	}
	cm = execute.Execute("flutter", "pub", "add", "flutter_modular", "mobx", "flutter_mobx", "mobx_codegen", "icons_plus", "google_fonts")
	cm.Dir = path.Join(rootPath, name)
	if err := cm.Run(); err != nil {
		fmt.Println("error get depedencies:", err)
		os.Exit(1)
	}

	dirs := []string{
		"lib/app",
		"lib/app/core",
	}

	for _, d := range dirs {
		os_gen.Mkdir(path.Join(name, d))
	}

	files := map[string]string{
		"lib/main.dart":           flutter.Main,
		"lib/app/app_module.dart": flutter.AppModule,
		"lib/app/app_widget.dart": flutter.AppWidget,
	}

	os.Remove(path.Join(name, "lib/main.dart"))

	for filename, content := range files {
		fileresult := path.Join(name, filename)
		if err := template_gen.WriteByTemplate(content, fileresult, map[string]string{
			"name": name,
			"org":  org,
		}); err != nil {
			fmt.Println("error create file", fileresult, err)
			continue
		}
		fmt.Println("created", fileresult)
	}

}
