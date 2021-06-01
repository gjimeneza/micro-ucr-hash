package hash

// Next trata de evitar un overflow
// si se llega a 255 que se mantenga ahi
// y siga contando con la siguiente posicion
func (ha *HashArea) Next(nonce *Nonce) *Nonce {

	// Comentario para Leo: nonce es un puntero de tipo Nonce
	// En Go utilice punteros como si fueran datos normales, Go ahi ve como hace
	// super facil!

	if nonce[3] < 255 {
		nonce[3] = nonce[3] + 1

	} else if nonce[2] < 255 {
		nonce[2] = nonce[2] + 1

	} else if nonce[1] < 255 {
		nonce[1] = nonce[1] + 1

	} else if nonce[0] < 255 {
		nonce[0] = nonce[0] + 1

	}
	return nonce
}

func (ha *HashArea) MicroHashUcr(p Payload, n Nonce) Bounty {
	bloc := Bloque{}
	w := make([]byte, 32)

	// Concatenacion del nonce y el array de bytes de entrada
	copy(bloc[:], p[:])       // de la posicion 0 - 12 los bytes de entrada
	copy(bloc[len(p):], n[:]) // de la posicion 12-16 el nonce

	for i := 0; i <= 15; i++ {
		w[i] = bloc[i]
	}
	for i := 16; i <= 31; i++ {
		w[i] = w[i-3] | (w[i-9] ^ w[i-14])
	}

	h := [3]byte{0x01, 0x89, 0xfe}

	var k byte
	var x byte

	var a byte = h[0]
	var b byte = h[1]
	var c byte = h[2]

	for i := 0; i <= 31; i++ {

		// var a byte = h[0]
		// var b byte = h[1]
		// var c byte = h[2]

		if i <= 0 && i <= 16 {
			k = 0x99
			x = (a ^ b) & 0xFF
		} else if i <= 17 && i <= 31 {
			k = 0xa1
			x = (a | b) & 0xFF
		}

		a = (b ^ c) & 0xFF
		b = (c << 4) & 0xFF
		c = (x + k + w[i]) & 0xFF

		// if i == 31 {
		// 	h[0] = h[0] + a
		// 	h[1] = h[1] + b
		// 	h[2] = h[2] + c
		// }

	}

	h[0] = (h[0] + a) & 0xFF
	h[1] = (h[1] + b) & 0xFF
	h[2] = (h[2] + c) & 0xFF

	// Casteo a tipo de bounty
	return Bounty(h)

}

func (ha *HashArea) Sistema(inicio bool, target byte, p Payload) (Nonce, bool) {
	return Nonce{}, false
}

func (h *HashArea) ValidateBounty(target byte) {

	// Esta funcion valida si el bounty calculado esta dentro del target especificado

	var valid int  // define el valido
	var valido int // retorno
	valid = 0

	if bounty[0] < target && bounty[1] < target {
		valid = 1

	} else {
		valid = 0
	}

	valido = valid

	return valido

}
