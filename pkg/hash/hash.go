package hash

type Bounty [3]byte
type Nonce [4]byte
type Payload [12]byte
type Bloque [16]byte

// HashArea is a struct that describes a hashing system designed for small circuit area
type HashArea struct{}

// HashArea is a struct that describes a hashing system designed for speed
type HashSpeed struct{}

// Hash interface can implement Sistema from any type that uses it
type Hash interface {
	Sistema(inicio bool, target byte, p Payload) (Nonce, bool)
}
