package hash

import "encoding/binary"

// Uint32ToNonce converts a uint32 primitive number to a Nonce value
func Uint32ToNonce(u uint32) Nonce {
	nonce := Nonce{}
	binary.BigEndian.PutUint32(nonce[:], u)
	return nonce
}
