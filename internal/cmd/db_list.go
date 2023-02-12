package cmd

import (
	"github.com/chiselstrike/iku-turso-cli/internal/settings"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:               "list",
	Short:             "List databases.",
	Args:              cobra.NoArgs,
	ValidArgsFunction: noFilesArg,
	RunE: func(cmd *cobra.Command, args []string) error {
		settings, err := settings.ReadSettings()
		if err != nil {
			return err
		}
		databases, err := getDatabases(createTursoClient())
		if err != nil {
			return err
		}
		data := [][]string{}
		for _, database := range databases {
			url := getDatabaseUrl(settings, database)
			regions := getDatabaseRegions(database)
			data = append(data, []string{database.Name, database.Type, regions, url})
		}
		printTable([]string{"name", "type", "regions", "url"}, data)
		settings.SetDbNamesCache(extractDatabaseNames(databases))
		return nil
	},
}