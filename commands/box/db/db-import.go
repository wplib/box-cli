package db

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wplib/box-cli/commands/box"
)

// dbImportCmd represents the import command
var dbImportCmd = &cobra.Command{
	Use:   "import",
	Aliases: []string{"restore"},
	Short: "Import/Restore a database file",
	Long: `Import/Restore a database dump (.SQL file) to the database named 'wordpress' in WPLib Box`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("import called")
	},
}

func init() {
	box.DbCmd.AddCommand(dbImportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
