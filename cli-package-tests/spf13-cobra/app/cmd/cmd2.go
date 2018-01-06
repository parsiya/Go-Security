package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cmd2Cmd represents the cmd2 command
var cmd2Cmd = &cobra.Command{
	Use:   "cmd2",
	Short: "Short description of cmd2",
	Long:  `Short description of cmd2`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmd2 called")
	},
}

func init() {
	cmd1Cmd.AddCommand(cmd2Cmd)
}
