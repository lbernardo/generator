package os_gen

import (
	"fmt"
	"os"
)

func Mkdir(path string) {
	os.MkdirAll(path, 0755)
	fmt.Println("created", path)
}
