package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dbAuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage database authentication",
}

func init() {
	dbAuthCmd.AddCommand(dbAuthTokenCmd)
}

func dbAuthTokenArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) == 0 {
		return getDatabaseNames(createTursoClient()), cobra.ShellCompDirectiveNoFileComp
	}
	return []string{}, cobra.ShellCompDirectiveNoFileComp
}

var dbAuthTokenCmd = &cobra.Command{
	Use:               "token database_name",
	Short:             "Creates a bearer token to authenticate requests to the database",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: dbAuthTokenArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		turso := createTursoClient()
		token, err := turso.Databases.Token(args[0])
		if err != nil {
			return fmt.Errorf("your database does not support token generation")
		}
		fmt.Println(token)
		return nil
	},
}