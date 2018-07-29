package box

import (
	"fmt"

	"github.com/spf13/cobra"
)

// DbCmd represents the database command
var DbCmd = &cobra.Command{
	Use:   "db",
	Aliases: []string{"databases"},
	Short: "Manage MySQL-type databases that WPLib Box provides for you",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("database called")
	},
}

func init() {
	RootCmd.AddCommand(DbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// databaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// databaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
