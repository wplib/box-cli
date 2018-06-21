package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var dbExportCmd = &cobra.Command{
	Use:   "export",
	Aliases: []string{"backup"},
	Short: "Export/Backup a database file",
	Long: `Export/Backup from the database named 'wordpress' in WPLib Box to a database dump named backup.sql.
If a backup.sql already exists it will increment by 1, e.g. backup1.sql, backup2.sql, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("export called")
	},
}

func init() {
	dbCmd.AddCommand(dbExportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
