package hash

import (
	"encoding/binary"
)

// NextNonce takes current Nonce, converts it in its uint32 representation (4
// bytes), increments it by one and returns the number's Nonce form.
func (hs *HashSpeed) NextNonce(nonce Nonce) Nonce {
	nonceU32 := binary.BigEndian.Uint32(nonce[:])
	nonceU32++
	binary.BigEndian.PutUint32(nonce[:], nonceU32)

	return nonce
}

// Concatenate takes a payload (12 bytes) and a Nonce (4 bytes) and concatenates them into
// a Bloque (16 bytes)
func (hs *HashSpeed) Concatenate(p Payload, n Nonce) Bloque {
	bloque := Bloque{}

	copy(bloque[:], p[:])
	copy(bloque[len(p):], n[:])

	return bloque
}

// MicroHashUcr is the main hashing function, it takes a Bloque (16 bytes),
// makes some predefined bitwise operations and returns a Bounty (3 bytes)
func (hs *HashSpeed) MicroHashUcr(bloque Bloque) Bounty {
	w := make([]byte, 32)

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

	return Bounty(h)
}

// CheckBounty takes a Bounty (3 bytes) a target (1 byte) and returns true if the
// first two bytes (Little Endian) are below the Target.
func (hs *HashSpeed) CheckBounty(bounty Bounty, target byte) bool {
	if (bounty[0] < target) && (bounty[1] < target) {
		return true
	}

	return false
}

// Sistema is the main system that encompasses nonce generation, hash creation
// and bounty checking. It receives a signal to start (inicio), a target (1
// byte) and a Payload (12 bytes) and returns the first Nonce that meets the
// target requirements according to the Bounty returned by the hashing function.
// It is implemented concurrently using goroutines to simulate parallelization
// in a real integrated circuit.
//
// The system divides the total possible Nonces in n goroutines that each start
// in a different Nonce and starts checking Bounties from there. The first
// goroutine that obtains a valid Bounty, returns the Nonce through a Go channel
// which in turn closes it for the rest of the goroutines. Finally, this Nonce
// is returned.
func (hs *HashSpeed) Sistema(inicio bool, target byte, p Payload) (Nonce, bool) {
	if !inicio {
		return Nonce{}, false
	}

	n := 6
	channel := make(chan Nonce)

	for i := 0; i < n; i++ {
		nonceU32 := (4294967296 / n) * i
		nonce := Uint32ToNonce(uint32(nonceU32))
		go hs.sistema(target, nonce, p, channel)
	}

	res := <-channel

	return res, true
}

// sistema is the function called by the goroutines in System, it encompasses
// nonce generation, hash creation and bounty checking. When a Nonce that meets
// the bounty target is found it is returned through a go channel.
func (hs *HashSpeed) sistema(target byte, initNonce Nonce, p Payload, c chan Nonce) {
	nonce := initNonce
	bloque := hs.Concatenate(p, nonce)
	bounty := hs.MicroHashUcr(bloque)
	terminado := hs.CheckBounty(bounty, target)

	for !terminado {
		nonce = hs.NextNonce(nonce)
		bloque = hs.Concatenate(p, nonce)
		bounty = hs.MicroHashUcr(bloque)
		terminado = hs.CheckBounty(bounty, target)
	}

	//fmt.Println("Bounty:", bounty)
	//fmt.Println("Nonce:", nonce)

	c <- nonce
}
