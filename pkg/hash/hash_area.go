package hash

import (
	"encoding/binary"
)

var AreaHashOutput HashOutput // to return final HashOutput

// NextNonce takes current Nonce, converts it in its uint32 representation (4
// bytes), increments it by one and returns the number's Nonce form.
func (ha *HashArea) NextNonce(nonce Nonce) Nonce {

	nonceU32 := binary.BigEndian.Uint32(nonce[:])

	nonceU32++

	binary.BigEndian.PutUint32(nonce[:], nonceU32)

	return nonce
}

// Concatenate takes a payload (12 bytes) and a Nonce (4 bytes) and concatenates them into
// a Bloque (16 bytes)
func (h *HashArea) Concatenador(p Payload, n Nonce) Bloque {
	bloque := Bloque{}

	// Concatenacion del nonce y el array de uint32s de entrada
	copy(bloque[:], p[:])       // de la posicion 0 - 12 los bytes de entrada
	copy(bloque[len(p):], n[:]) // de la posicion 12-16 el nonce

	return bloque
}

// MicroHashUcr is the main hashing function, it takes a Bloque (16 bytes),
// makes some predefined bitwise operations and returns a HashOutput (3 bytes)
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

// ValidateOutput takes a HashOutput (3 bytes) a target (1 byte) and returns true if the
// first two bytes (Little Endian) are below the Target.
func (h *HashArea) ValidateOutput(target byte, hashOutput HashOutput) bool {

	// Esta funcion valida si el hashOutput calculado esta dentro del target especificado

	if (hashOutput[0] < target) && (hashOutput[1] < target) {
		return true
	}

	return false

}

// Sistema is the main system that encompasses nonce generation, hash creation
// and HashOutput checking. It receives a signal to start (inicio), a target (1
// byte) and a Payload (12 bytes) and returns the first Nonce that meets the
// target requirements according to the HashOutput returned by the hashing function.
//
// It is implemented in a way to use the least amount of modules in order to reduce
// the area needed to produce an integrated circuit
func (hs *HashArea) Sistema(inicio bool, target byte, p Payload) (Nonce, bool) {

	if !inicio {
		return Nonce{}, false
	}

	return hs.sistemaIntern(target, Nonce{0x00, 0x00, 0x00, 0x00}, p)
}

// sistemaIntern is the function called by the goroutines in System, it encompasses
// nonce generation, hash creation and HashOutput checking. When a Nonce that meets
// the HashOutput target is found it is returned through a go channel along with its
// HashOutput for further inspection.
func (hs *HashArea) sistemaIntern(target byte, initNonce Nonce, p Payload) (Nonce, bool) {

	nonce := initNonce
	bloque := hs.Concatenador(p, nonce)
	hashOutput := hs.MicroHashUcr(bloque)
	terminado := hs.ValidateOutput(target, hashOutput)

	for !terminado {

		nonce = hs.NextNonce(nonce)
		bloque = hs.Concatenador(p, nonce)
		hashOutput = hs.MicroHashUcr(bloque)
		terminado = hs.ValidateOutput(target, hashOutput)

	}

	AreaHashOutput = hashOutput

	return nonce, true
}
