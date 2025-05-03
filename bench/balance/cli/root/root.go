package root

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "balance-bench-cli",
	Short: "Balance is a CLI tool for benchmarking balance",
	Long:  `Balance is a CLI tool for benchmarking balance`,
	Args:  cobra.ExactArgs(1),
}
