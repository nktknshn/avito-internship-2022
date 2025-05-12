package signin

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli/deps"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli/root"
)

func init() {
	root.RootCmd.AddCommand(cmdSignIn)
}

var cmdSignIn = &cobra.Command{
	Use:   "signin",
	Short: "Sign in",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		cliAdapter, err := deps.GetCliAdapter(ctx)
		if err != nil {
			return err
		}
		token, err := cliAdapter.SignIn(ctx, args[0], args[1])
		if err != nil {
			return err
		}

		//nolint:forbidigo // expected
		fmt.Println(token)
		return nil
	},
}
