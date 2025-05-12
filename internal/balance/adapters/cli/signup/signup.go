package signup

import (
	"github.com/spf13/cobra"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli/deps"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli/root"
)

func init() {
	root.RootCmd.AddCommand(cmdSignUp)
}

var cmdSignUp = &cobra.Command{
	Use:   "signup",
	Short: "Sign up",
	//nolint:mnd // ...
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		cliAdapter, err := deps.GetCliAdapter(ctx)
		if err != nil {
			return err
		}
		err = cliAdapter.SignUp(ctx, args[0], args[1], args[2])
		if err != nil {
			return err
		}
		return nil
	},
}
