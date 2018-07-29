package cargo

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wplib/box-cli/commands/box"
)

// executableCmd represents the executable command
var executableCmd = &cobra.Command{
	Use:   "executable",
	Short: "Manage 'Executables' for WPLib Box",
	Long: "Manage 'Executables' for WPLib Box which typically come in the form of a Docker container.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executable called")
	},
	Hidden:true,
}

func init() {
	box.RootCmd.AddCommand(executableCmd)
	annotations := make(map[string]string)
	annotations["cargo"] = "yes"
	executableCmd.Annotations = annotations

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// executableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// executableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
