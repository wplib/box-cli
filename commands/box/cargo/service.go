package cargo

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wplib/box-cli/commands/box"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage 'Services' for WPLib Box",
	Long: "Manage 'Services' for WPLib Box which typically come in the form of a Docker container.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service called")
	},
	Hidden:true,
}

func init() {
	box.RootCmd.AddCommand(serviceCmd)
	annotations := make(map[string]string)
	annotations["cargo"] = "yes"
	serviceCmd.Annotations = annotations

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
