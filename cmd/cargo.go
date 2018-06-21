package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cargoCmd represents the 'cargo' command
var cargoCmd = &cobra.Command{
	Use:   "cargo",
	Short: "Manage your various 'Cargo' (WPLib Box's extensibility model)",
	Long: "Manage 'Cargo; for WPLib Box where Cargo represents WPLib Box's extensibility model and may come form of Services, Executables, (website) Source (code) and/or Stacks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cargo called")
	},
}

func init() {
	RootCmd.AddCommand(cargoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// containerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// containerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
