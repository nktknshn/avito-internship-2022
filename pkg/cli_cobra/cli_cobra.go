package cli_cobra

import (
	"github.com/spf13/cobra"
)

type CliCobra struct {
	RootCmd *cobra.Command
}

func NewCliCobra() *CliCobra {
	return &CliCobra{
		RootCmd: &cobra.Command{},
	}
}
