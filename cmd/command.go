package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// commandCmd represents the 'cargo' command
var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Manage custom commands for WPLib Box's 'Box' tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cargo called")
	},
}

func init() {
	RootCmd.AddCommand(commandCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// containerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// containerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
