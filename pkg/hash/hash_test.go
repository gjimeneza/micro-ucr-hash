package hash_test

import (
	"fmt"
	"testing"

	"github.com/gjimeneza/micro-ucr-hash/pkg/hash"
)

func TestHashArea(t *testing.T) {
	testPayload := hash.Payload{0x39, 0x7d, 0x9f, 0x2f, 0x40, 0xca, 0x9e, 0x6c, 0x6b, 0x1f, 0x33, 0x24}
	testNonce := hash.Nonce{0xfd, 0xed, 0x87, 0x3c}

	ha := hash.HashArea{}
	bounty := ha.MicroHashUcr(testPayload, testNonce)

	fmt.Println("El bounty es:", bounty)

	nextNonce := ha.Next(&testNonce)
	fmt.Println("El nuevo nonce es:", *nextNonce)
}
