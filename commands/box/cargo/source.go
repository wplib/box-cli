package cargo

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wplib/box-cli/commands/box"
)

// sourceCmd represents the source command
var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Manage (website) 'Source' (code) for WPLib Box",
	Long: "Manage website 'Source' code for WPLib Box which typically comes in the form of .PHP, .HTML, .CSS, .JS, or related; or packaged in a .ZIP file.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("source called")
	},
	Hidden:true,
}

func init() {
	box.RootCmd.AddCommand(sourceCmd)
	annotations := make(map[string]string)
	annotations["cargo"] = "yes"
	sourceCmd.Annotations = annotations

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sourceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sourceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
