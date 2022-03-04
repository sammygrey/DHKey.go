package utils

import (
	"crypto/rand"
)

type Endpoint struct {
	PublicBase   int //upto int64 - the public base known to both people
	PublicModulo int // - the public modulus known to both people
	PrivateKey   int // - the private key known only to each person
}

func NewBaseModulo() []int {
	//ideally g**q = 1 mod p, where q is a random prime integer, but all prime numbers should work
	p, _ := rand.Prime(rand.Reader, 64)
	g, _ := rand.Prime(rand.Reader, 64)
	base, modulo := uint32(p.Int64()), uint32(g.Int64()) //we don't want any negative numbers here, converting to uint and back solves this
	return []int{int(base), int(modulo)}
}

func NewEndpoint(publicBase, publicModulo, privateKey int) Endpoint {
	//creates a endpoint struct using an oop style function
	return Endpoint{publicBase, publicModulo, privateKey}
}

func GenPartial(end Endpoint) int {
	//generate public key using private key and public parts to hand over to other party
	//this is safe to directly hand over
	partial := exp(end.PublicBase, end.PrivateKey)
	return partial % end.PublicModulo
}

func GenFull(end Endpoint, partialKey int) int {
	//generate full shared secret using the other parties' public key and our personal endpoint
	//this should not be shared directly
	return exp(partialKey, end.PrivateKey) % end.PublicModulo
}

func Encrypt(end Endpoint, partialKey int, message string) string {
	var encrypted []rune //int32
	//encode each character to an integer, add the resultant int value of the secret to encode it
	for i, char := range message {
		encrypted[i] = rune(int(char) + GenFull(end, partialKey))
	}
	return string(encrypted)

}

func Decrypt(end Endpoint, partialKey int, encrypted string) string {
	var message []rune
	//perform the opposite of what we did to encrypt, to decrypt it
	for i, char := range encrypted {
		message[i] = rune(int(char) - GenFull(end, partialKey))
	}
	return string(message)
}

func exp(x, y int) int {
	// fast binary operation for calculating power of integers, only works to int64 max
	// this also isn't set up for negative integers for root style expressions
	power := 1
	for y > 0 {
		// if y bit is 1, you could also do y%2 != 0 here
		if y&1 != 0 {
			power *= x
		}
		//eventually this will make y == 0
		y >>= 1 //binary right shift -> y = y/2 once, you could also just do y /= 2
		x *= x
	}
	return power
}
