
package cargo

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wplib/box-cli/cmd"
)

// stackCmd represents the stack command
var stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Manage 'Stacks' in WPLib Box",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stack called")
	},
	Hidden:true,
}

func init() {
	cmd.RootCmd.AddCommand(stackCmd)
	annotations := make(map[string]string)
	annotations["cargo"] = "yes"
	stackCmd.Annotations = annotations

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
