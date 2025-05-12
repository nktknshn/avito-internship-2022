package list_users

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli/deps"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli/root"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_list_users"
)

var flagPrintHeader = true

func init() {
	root.RootCmd.AddCommand(cmdListUsers)
	cmdListUsers.Flags().BoolVar(&flagPrintHeader, "print-header", true, "Print header")
}

var cmdListUsers = &cobra.Command{
	Use:   "list-users",
	Short: "List users",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		cliAdapter, err := deps.GetCliAdapter(ctx)
		if err != nil {
			return err
		}

		users, err := cliAdapter.ListUsers(ctx)
		if err != nil {
			return err
		}

		var username string

		if len(args) > 0 {
			username = args[0]
		}

		printUsers(users, username)

		return nil
	},
}

func printUsers(users []auth_list_users.OutUser, username string) {
	tabwriter := tabwriter.NewWriter(
		os.Stdout,
		0,
		8,
		2,
		'\t',
		0,
	)
	defer tabwriter.Flush()

	if flagPrintHeader {
		fmt.Fprintln(tabwriter, "ID\tUsername\tRole")
	}

	for _, user := range users {
		if username != "" && user.Username.Value() != username {
			continue
		}
		fmt.Fprintf(tabwriter, "%s\t%s\t%s\n", user.ID.String(), user.Username.Value(), user.Role.Value())
	}
}
