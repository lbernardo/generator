package flutter

import (
	"fmt"
	"github.com/lbernardo/generator/pkg/execute"
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
	cm = execute.Execute("flutter", "pub", "get", "flutter_modular", "mobx", "flutter_mobx", "mobx_codegen", "icons_plus", "google_fonts")
	cm.Dir = path.Join(rootPath, name)
	if err := cm.Run(); err != nil {
		fmt.Println("error get depedencies:", err)
		os.Exit(1)
	}
}
