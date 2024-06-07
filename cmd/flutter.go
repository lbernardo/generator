package cmd

import (
	"github.com/lbernardo/generator/internal/generator/flutter"
	"github.com/lbernardo/generator/pkg/entity"
)

var FlutterCmd = &entity.Module{
	Name: "flutter",
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
					Name:        "org",
					Description: "project organization",
				},
			},
			Run: flutter.NewProject,
		},
		&entity.Generator{
			Name: "feature",
			Params: []entity.Param{
				{
					Name:        "name",
					Description: "feature name",
					Required:    true,
				},
				{
					Name:        "child",
					Description: "if your feature is a child of other child",
				},
			},
			Run: flutter.NewFeature,
		},
		&entity.Generator{
			Name:        "page",
			Description: "Create a new page with controller",
			Params: []entity.Param{
				{
					Name:        "name",
					Description: "page name",
					Required:    true,
				},
				{
					Name:        "path",
					Description: "path from lib/app",
				},
			},
			Run: flutter.NewPage,
		},
		&entity.Generator{
			Name: "controller",
			Params: []entity.Param{
				{
					Name:        "name",
					Description: "controller name",
					Required:    true,
				},
				{
					Name:        "path",
					Description: "path from lib/app",
				},
			},
			Run: flutter.NewController,
		},
	},
}

func init() {
	AddCommand(FlutterCmd)
}
