package hash

// Main hashing modules will go here

var nonce = [4]byte{0, 0, 0, 0} // variable global el nonce, asi puede ser manipulada por las demas funciones

func next() [4]byte {

	// la logica trata de evitar un overflow
	// si se llega a 255 que se mantenga ahi
	// y siga contando con la siguiente posicion

	if nonce[3] < 255 {
		nonce[3] = nonce[3] + 1

	} else if nonce[2] < 255 {
		nonce[2] = nonce[2] + 1

	} else if nonce[1] < 255 {
		nonce[1] = nonce[1] + 1

	} else if nonce[0] < 255 {
		nonce[0] = nonce[0] + 1

	}

	next_nonce := nonce

	return next_nonce

}

func micro_hash_ucr(arr_in [12]byte, nonce_in [4]byte) [3]byte {

	var bloque [16]byte // guarda los 16 bits
	var bounty [3]byte  // retorna el bounty calculado
	var w [32]byte

	// concatenacion del nonce y el array de bytes de entrada
	copy(bloque[:], arr_in[:])              // de la posicion 0 - 12 los bytes de entrada
	copy(bloque[len(arr_in):], nonce_in[:]) // de la posicion 12-16 el nonce

	for i := 0; i <= 15; i++ {
		w[i] = bloque[i]
	}

	for i := 16; i <= 31; i++ {
		w[i] = w[i-3] | w[i-9] ^ w[i-14]
	}

	h := [3]byte{0x01, 0x89, 0xfe}

	var k byte
	var x byte
	for i := 0; i <= 32; i++ {

		var a byte = h[0]
		var b byte = h[1]
		var c byte = h[2]

		if i <= 0 && i <= 16 {
			k = 0x99
			x = a ^ b
		} else if i <= 17 && i <= 31 {
			k = 0xa1
			x = a | b
		}

		a = b ^ c
		b = c << 4
		c = x + k + w[i]

		if i == 32 {
			h[0] = h[0] + a
			h[1] = h[1] + b
			h[2] = h[2] + c
		}

	}

	bounty = h

	return bounty // se retorna el bounty calculado

}

func validate_bounty(target byte) {

	// esta funcion valida si el bounty calculado esta dentro del target especificado

}
