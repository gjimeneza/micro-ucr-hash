package hash_test

import (
	"fmt"
	"testing"

	"github.com/gjimeneza/micro-ucr-hash/pkg/hash"
)

func TestHashArea(t *testing.T) {
	testPayload := hash.Payload{0x39, 0x7d, 0x9f, 0x2f, 0x40, 0xca, 0x9e, 0x6c, 0x6b, 0x1f, 0x33, 0x24} // datos de entrada
	testNonce := hash.Nonce{0x00, 0x00, 0x00, 0x00}                                                     // nonce de entrada, se inicializa en 0,0,0,0

	var bounty hash.Bounty
	terminado := 0
	ha := hash.HashArea{}

	i := 0

	for (terminado == 0) && (i <= 1024) {

		fmt.Printf("Entrada: [%# x]\n", testPayload[:]) // imprime el bounty obtenido
		fmt.Printf("Nounce: [%# x]\n", testNonce[:])    // imprime el nounce utilizado
		fmt.Printf("\n")                                // nueva linea

		bounty = ha.MicroHashUcr(testPayload, testNonce) // obtiene el bounty
		terminado = ha.ValidateBounty(10, bounty)        // verifica que sea un bounty correcto

		if terminado == 0 {
			ha.Next(&testNonce) // si no se ha terminado, calcular proximo nounce
		}

		i = i + 1

	}

	fmt.Printf("El proceso ha terminado.\n")
	fmt.Printf("El bounty obtenido es: [%# x]\n", bounty[:]) // imprime el bounty obtenido

}
