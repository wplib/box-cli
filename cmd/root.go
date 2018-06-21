package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "box",
	Short: "Box is the command-line tool for WPLib Box that you run in your terminal.",
	Long: `Box is the command-line tool for WPLib Box. 
Box is designed to run both on your host Mac, Windows or Linux computer but also inside the WPLib Box virtual machine (VM).
To run "box" inside the VM, first type "vagrant ssh" on your command line.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init()  {
}