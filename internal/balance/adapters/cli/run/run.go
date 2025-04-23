package run

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli/root"
	"github.com/spf13/cobra"
)

func init() {
	root.RootCmd.AddCommand(cmdRun)
}

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run the balance service",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
