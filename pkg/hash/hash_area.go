package hash

import (
	"encoding/binary"
	"fmt"
)

var AreaHashOutput HashOutput

func (ha *HashArea) NextNonce(nonce Nonce) Nonce {

	nonceU32 := binary.BigEndian.Uint32(nonce[:])

	nonceU32++

	binary.BigEndian.PutUint32(nonce[:], nonceU32)

	return nonce
}

func (h *HashArea) Concatenador(p Payload, n Nonce) Bloque {
	bloque := Bloque{}

	// Concatenacion del nonce y el array de uint32s de entrada
	copy(bloque[:], p[:])       // de la posicion 0 - 12 los bytes de entrada
	copy(bloque[len(p):], n[:]) // de la posicion 12-16 el nonce

	return bloque
}

func (ha *HashArea) MicroHashUcr(bloque Bloque) HashOutput {
	w := make([]byte, 32)

	// Proceso principal
	for i := 0; i <= 15; i++ {
		w[i] = bloque[i]
	}
	for i := 16; i <= 31; i++ {
		w[i] = w[i-3] | (w[i-9] ^ w[i-14])
	}

	h := [3]byte{0x01, 0x89, 0xfe}
	a := h[0]
	b := h[1]
	c := h[2]

	for i := 0; i < 32; i++ {
		var (
			k byte
			x byte
		)

		if i <= 16 {
			k = 0x99
			x = a ^ b
		} else {
			k = 0xa1
			x = a | b
		}

		a = b ^ c
		b = c << 4
		c = x + k + w[i]
	}

	h[0] = h[0] + a
	h[1] = h[1] + b
	h[2] = h[2] + c

	return HashOutput(h)

}

func (h *HashArea) ValidateOutput(target byte, hashOutput HashOutput) bool {

	// Esta funcion valida si el hashOutput calculado esta dentro del target especificado

	if (hashOutput[0] < target) && (hashOutput[1] < target) {
		return true
	}

	return false

}

func (hs *HashArea) Sistema(inicio bool, target byte, p Payload) (Nonce, bool) {

	if !inicio {
		return Nonce{}, false
	}

	// default Nonce{0x00, 0x00, 0x00, 0x00}
	// bloque 1 Nonce{0xfd, 0xed, 0x87, 0x3c}
	// bloque 2 Nonce{0x0f, 0xa2, 0x34, 0x91}

	return hs.sistemaIntern(target, Nonce{0x00, 0x00, 0x00, 0x00}, p)
}

func (hs *HashArea) sistemaIntern(target byte, initNonce Nonce, p Payload) (Nonce, bool) {

	nonce := initNonce
	bloque := hs.Concatenador(p, nonce)
	hashOutput := hs.MicroHashUcr(bloque)
	fmt.Printf("hashOutput: [%# x]\n", hashOutput[:]) // para probar hashOutputs
	terminado := hs.ValidateOutput(target, hashOutput)

	for !terminado {

		nonce = hs.NextNonce(nonce)
		//fmt.Printf("Nonce: [%# x]\n", nonce[:])
		bloque = hs.Concatenador(p, nonce)
		hashOutput = hs.MicroHashUcr(bloque)
		terminado = hs.ValidateOutput(target, hashOutput)

	}

	AreaHashOutput = hashOutput

	return nonce, true
}
