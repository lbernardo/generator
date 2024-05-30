package entity

import (
	"github.com/spf13/cobra"
)

type Module struct {
	Name       string
	Generators []*Generator
}

func (m *Module) Command() *cobra.Command {
	var cmd = &cobra.Command{
		Use: m.Name,
	}
	for _, gen := range m.Generators {
		gen.addComand(cmd)
	}
	return cmd
}
