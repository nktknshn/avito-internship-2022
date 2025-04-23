package root

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "balance",
	Args: cobra.MinimumNArgs(1),
}

var (
	flagConfigPath = "config.yaml"
)

func init() {
	RootCmd.PersistentFlags().StringVar(&flagConfigPath, "cfg-path", flagConfigPath, "config path")
}

func GetConfigPath() string {
	return flagConfigPath
}
