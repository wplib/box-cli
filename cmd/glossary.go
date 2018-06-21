
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// glossaryCmd represents the 'glossary' command
var glossaryCmd = &cobra.Command{
	Use:   "glossary",
	Short: "Display a Glossary of WPLib Box terms to your terminal.",
	Long: "Output a glossary of terms to help you understand the meaning of all the names WPLib Box gives to its various components.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("glossary called")
	},
}

func init() {
	RootCmd.AddCommand(glossaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// containerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// containerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
