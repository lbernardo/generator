package template_gen

import (
	"fmt"
	"os"
	"text/template"
)

func WriteByTemplate(templateContent string, fileResult string, variables map[string]string) error {
	file, err := os.OpenFile(fileResult, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("error to create file %v: %v", fileResult, err)
	}

	tpl, err := template.New("new-template").Parse(templateContent)
	if err != nil {
		return fmt.Errorf("error to load template: %v", err)
	}
	if err := tpl.Execute(file, variables); err != nil {
		return fmt.Errorf("error execute template: %v", err)
	}
	return nil
}
