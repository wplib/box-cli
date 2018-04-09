package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// installableCmd represents the installable command
var installableCmd = &cobra.Command{
	Use:   "installable",
	Short: "Manage Installables in WPLib Box",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("installable called")
	},
}

func init() {
	rootCmd.AddCommand(installableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
