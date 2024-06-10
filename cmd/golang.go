package cmd

import (
	"github.com/lbernardo/generator/internal/generator/golang"
	"github.com/lbernardo/generator/pkg/entity"
)

var GolangCmd = &entity.Module{
	Name: "golang",
	Generators: []*entity.Generator{
		&entity.Generator{
			Name:        "project",
			Description: "Create a new project",
			Params: []entity.Param{
				{
					Name:        "name",
					Description: "project name",
					Required:    true,
				},
				{
					Name:        "pkg",
					Description: "package module",
					Required:    true,
				},
			},
			Run: golang.NewProject,
		},
		&entity.Generator{
			Name: "app",
			Params: []entity.Param{
				{
					Name:        "name",
					Description: "app name",
					Required:    true,
				},
			},
			Run: golang.NewApp,
		},
	},
}

func init() {
	AddCommand(GolangCmd)
}
