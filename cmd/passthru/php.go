package passthru

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wplib/box-cli/cmd"
)

// phpCmd represents the database command
var phpCmd = &cobra.Command{
	Use:   "php",
	Short: "[Experimental] Manage PHP versions that WPLib Box serves for you",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("php called")
	},
	Hidden:true,
}

func init() {
	cmd.RootCmd.AddCommand(phpCmd)
	annotations := make(map[string]string)
	annotations["passthru"] = "passthru"
	phpCmd.Annotations = annotations

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// databaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// databaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
