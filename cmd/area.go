package cmd

import (
	"fmt"
	"time"

	"github.com/gjimeneza/micro-ucr-hash/pkg/hash"
	"github.com/spf13/cobra"
)

var areaCmd = &cobra.Command{
	Use:   "area",
	Short: "Nonce generation focused on area",
	Long: `Nonce generation focused on area.
It takes a payload and target and generates a valid nonce.
Example usage:
	 ./micro-ucr-hash area -p 397d9f2f40ca9e6c6b1f3324 -t 10
	 ./micro-ucr-hash area -p ed18be0f984ae0e2e3128efe -t 10`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Entered payload: [%# x]\n", payload)
		fmt.Printf("Desired target: %v\n", target)

		ha := hash.HashArea{}

		p := hash.Payload{}
		copy(p[:], payload)

		start := time.Now()
		nonce, _ := ha.Sistema(true, target, p)
		duration := time.Since(start)

		fmt.Printf("Elapsed time: %v\n", duration)
		fmt.Printf("Generated HashOutput: [%# x]\n", hash.AreaHashOutput[:])
		fmt.Printf("Generated Nonce: [%# x]\n", nonce[:])
	},
}

func init() {
	rootCmd.AddCommand(areaCmd)

	// These flags can be used in terminal like:
	//./micro-ucr-hash area -p 397d9f2f40ca9e6c6b1f3324fd -t 10
	areaCmd.Flags().BytesHexVarP(&payload, "payload", "p", []byte{}, "Payload")
	areaCmd.Flags().Uint8VarP(&target, "target", "t", 10, "Target")
}
