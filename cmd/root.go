package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "box",
	Short: "Box is the terminal command-line tool for WPLib Box.",
	Long: `Box is the terminal command-line tool for WPLib Box. 
Box is designed to run inside the WPLib Box virtual machine.
To run "box" first type "vagrant ssh" on your command line.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init()  {
}