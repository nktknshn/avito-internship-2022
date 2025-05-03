package http

import (
	"github.com/nktknshn/avito-internship-2022-bench/cli/root"
	"github.com/nktknshn/avito-internship-2022-bench/logger"

	"github.com/spf13/cobra"
)

func init() {
	root.RootCmd.AddCommand(cmd)
	cmd.AddCommand(cmdBench)
}

var cmd = &cobra.Command{
	Use:   "http",
	Short: "HTTP benchmark",
	Long:  `HTTP benchmark`,
	Args:  cobra.ExactArgs(1),
}

var cmdBench = &cobra.Command{
	Use:   "bench",
	Short: "HTTP benchmark",
	Long:  `HTTP benchmark`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.GetLogger().Info("HTTP benchmark")

		client, err := getClientAuthorized("http://localhost:5050")
		if err != nil {
			logger.GetLogger().Error("Failed to get client", "error", err)
			return
		}

		b := Bench{client}
		b.Bench(cmd.Context())

		// resp, err := client.GetBalanceWithResponse(context.Background(), 1, client_http.GetBalanceJSONRequestBody{})

		// if err != nil {
		// 	logger.GetLogger().Error("Failed to get balance", "error", err)
		// 	return
		// }

		// logger.GetLogger().Info("Balance", "balance", string(resp.Body))

	},
}
