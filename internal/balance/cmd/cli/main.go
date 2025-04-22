package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	"github.com/nktknshn/avito-internship-2022/internal/balance/service"
	"github.com/nktknshn/avito-internship-2022/pkg/cli_cobra"
	"github.com/spf13/cobra"
)

var (
	flagConfigPath = "config.yaml"
	flagUsername   = ""
	flagPassword   = ""
	flagRole       = ""
)

var cmdSignUp = &cobra.Command{
	Use:   "signup",
	Short: "Sign up",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		cfg := config.Config{}
		err := cleanenv.ReadConfig(flagConfigPath, &cfg)

		if err != nil {
			return err
		}

		app, err := service.NewApplication(ctx, &cfg)

		if err != nil {
			return err
		}

		cliAdapter := cli.NewCliAdapter(app)

		err = cliAdapter.SignUp(ctx, args[0], args[1], args[2])

		if err != nil {
			return err
		}

		return nil
	},
}

var cliCobra = cli_cobra.NewCliCobra()

func init() {
	cliCobra.RootCmd.PersistentFlags().StringVar(&flagConfigPath, "cfg-path", flagConfigPath, "config path")
	cliCobra.RootCmd.AddCommand(cmdSignUp)

	// cmdSignUp.Flags().StringVar(&flagUsername, "username", flagUsername, "username")
	// cmdSignUp.Flags().StringVar(&flagPassword, "password", flagPassword, "password")
	// cmdSignUp.Flags().StringVar(&flagRole, "role", flagRole, "role")
}

func main() {

	err := cliCobra.RootCmd.Execute()

	if err != nil {
		panic(err)
	}
}
