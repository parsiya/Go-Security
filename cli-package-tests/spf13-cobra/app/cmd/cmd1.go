package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cmd1Cmd represents the cmd1 command
var cmd1Cmd = &cobra.Command{
	Use:   "cmd1",
	Short: "Short description of cmd1",
	Long:  `Long description of cmd1`,
	// cmd1 aliases
	Aliases: []string{"cmd11", "cmd12"},
	Run:     cmd1Executor,
}

// cmd1Executor is run when cmd1 is called.
func cmd1Executor(cmd *cobra.Command, args []string) {
	// Here we can see what command was run and also any arguments.
	fmt.Println("cmd1 called")
	// Print arguments after cmd1
	fmt.Println("arguments", args)
}

func init() {
	rootCmd.AddCommand(cmd1Cmd)
}
