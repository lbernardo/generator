package entity

import (
	"github.com/spf13/cobra"
)

type Generator struct {
	Name        string
	Description string
	Params      []Param
	Run         func(params map[string]string, rootPath string)
}

type Param struct {
	Name        string
	Value       string
	Description string
	Required    bool
}

func (g *Generator) addComand(cmd *cobra.Command) {
	var subCmd = &cobra.Command{
		Use:   g.Name,
		Short: g.Description,
		Run: func(cmd *cobra.Command, args []string) {
			var values = map[string]string{}
			for _, param := range g.Params {
				v, _ := cmd.Flags().GetString(param.Name)
				values[param.Name] = v
			}
			rootPath, _ := cmd.Flags().GetString("root")
			g.Run(values, rootPath)
		},
	}
	for _, param := range g.Params {
		subCmd.Flags().String(param.Name, param.Value, param.Description)
		if param.Required {
			subCmd.MarkFlagRequired(param.Name)
		}
	}
	cmd.AddCommand(subCmd)
}
