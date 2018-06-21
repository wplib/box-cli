package shorthand

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wplib/box-cli/cmd"
)

// webCmd represents the database command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "[Experimental] Manage websites that WPLib Box serves for you",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("web called")
	},
	Hidden: true,
}

func init() {
	cmd.RootCmd.AddCommand(webCmd)
	annotations := make(map[string]string)
	annotations["shorthand"] = "yes"
	webCmd.Annotations = annotations

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// databaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// databaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
