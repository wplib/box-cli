
package passthru

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wplib/box-cli/commands/box"
)

// wpCmd represents the database command
var wpCmd = &cobra.Command{
	Use:   "wp",
	Short: "Directly run WP-CLI commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wp called")
	},
	Hidden: true,
}

func init() {
	box.RootCmd.AddCommand(wpCmd)
	annotations := make(map[string]string)
	annotations["passthru"] = "passthru"
	wpCmd.Annotations = annotations

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// databaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// databaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
