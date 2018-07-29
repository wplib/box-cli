package box

import (
	"fmt"

	"github.com/spf13/cobra"
)

// componentCmd represents the 'component' command
var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Manage your various 'Component' (WPLib Box's extensibility model)",
	Long: "Manage 'Component; for WPLib Box where Component represents WPLib Box's extensibility model and may come form of Services, Executables, (website) Source (code) and/or Stacks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("component called")
	},
}

func init() {
	RootCmd.AddCommand(componentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// containerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// containerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
