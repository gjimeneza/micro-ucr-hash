package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "micro-ucr-hash",
	Short: "micro-ucr-hash is a functional description of a Nonce generator following the Micro Hash Ucr hashing function.",
	Long: `micro-ucr-hash provides a testbench to functionally
test a system that that encompasses nonce generation, hash creation
and bounty checking.

It provides two systems with different priorities,
one with a priority on less amount of modules and
the other on speed.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Call help on command
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
