package cmd

import (
	"fmt"
	"time"

	"github.com/gjimeneza/micro-ucr-hash/pkg/hash"
	"github.com/spf13/cobra"
)

var speedCmd = &cobra.Command{
	Use:   "speed",
	Short: "Nonce generation focused on speed",
	Long: `Nonce generation focused on speed.
It takes a payload and target and generates a valid nonce.

It uses goroutines to simulate parallelization by dividing
all possible Nonces in goroutines that each start in a
different Nonce and starts checking Bounties from there.

Example usage:
	 ./micro-ucr-hash speed -p 397d9f2f40ca9e6c6b1f3324fd -t 10`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Entered payload: [%# x]\n", payload)
		fmt.Printf("Desired target: %v\n", target)

		hs := hash.HashSpeed{}

		p := hash.Payload{}
		copy(p[:], payload)

		start := time.Now()
		nonce, _ := hs.Sistema(true, target, p)
		duration := time.Since(start)

		fmt.Printf("Elapsed time: %v\n", duration)
		fmt.Printf("Generated HashOutput: [%# x]\n", hash.SpeedHashOutput[:])
		fmt.Printf("Generated Nonce: [%# x]\n", nonce[:])
	},
}

func init() {
	rootCmd.AddCommand(speedCmd)

	// These flags can be used in terminal like:
	//./micro-ucr-hash speed -p 397d9f2f40ca9e6c6b1f3324fd -t 10
	speedCmd.Flags().BytesHexVarP(&payload, "payload", "p", []byte{}, "Payload")
	speedCmd.Flags().Uint8VarP(&target, "target", "t", 10, "Target")
}
