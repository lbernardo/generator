package modfile

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"os"
)

func GetModule(filePath string) (string, error) {

	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error to read go.mod: %v", err)
	}

	modFile, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return "", fmt.Errorf("error to parse module file: %v\n", err)
	}

	return modFile.Module.Mod.Path, nil
}
