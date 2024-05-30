package execute

import (
	"fmt"
	"os"
	"os/exec"
)

func Execute(name string, args ...string) *exec.Cmd {

	str := name
	for _, v := range args {
		str = str + " " + v
	}
	fmt.Println(str)

	cm := exec.Command(name, args...)
	cm.Stdout = os.Stdout
	cm.Stderr = os.Stderr
	return cm
}
