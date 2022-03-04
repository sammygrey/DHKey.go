package utils

import (
	"crypto/rand"
	"math/big"
	"strings"
)

type Endpoint struct {
	PublicBase   big.Int //upto int64 - the public base known to both people
	PublicModulo big.Int // - the public modulus known to both people
	PrivateKey   big.Int // - the private key known only to each person
}

func NewBaseModulo() []big.Int {
	//ideally g**q = 1 mod p, where q is a random prime integer, but all prime numbers should work
	p, _ := rand.Prime(rand.Reader, 64)
	g, _ := rand.Prime(rand.Reader, 64)
	base, modulo := p.Abs(p), g.Abs(g)
	return []big.Int{*base, *modulo}
}

func NewEndpoint(publicBase, publicModulo, privateKey big.Int) Endpoint {
	//creates a endpoint struct using an oop style function
	return Endpoint{publicBase, publicModulo, privateKey}
}

func GenPartial(end Endpoint) big.Int {
	//generate public key using private key and public parts to hand over to other party
	//this is safe to directly hand over
	partial := new(big.Int)
	partial.Exp(&end.PublicBase, &end.PrivateKey, &end.PublicModulo)
	return *partial
}

func GenFull(end Endpoint, partialKey big.Int) big.Int {
	//generate full shared secret using the other parties' public key and our personal endpoint
	//this should not be shared directly
	full := big.NewInt(0)
	full.Exp(&partialKey, &end.PrivateKey, &end.PublicModulo)
	return *full
}

func Encrypt(end Endpoint, partialKey big.Int, message string) string {
	var encrypted string
	//encode each character to an integer, add the resultant int value of the secret to further encrypt it
	//uses the shared modulo, plus the partial key + your private key
	fullKey := GenFull(end, partialKey)
	for _, char := range message {
		character := big.NewInt(int64(char))
		character.Add(character, &fullKey)
		encrypted += character.String() + ","
	}
	return encrypted

}

func Decrypt(end Endpoint, partialKey big.Int, encrypted string) string {
	var message []rune
	strSlice := strings.Split(encrypted, ",")
	fullKey := GenFull(end, partialKey)
	//undo the encryption process using the shared modulo, plus the partial key + your private key
	for _, strInt := range strSlice {
		num := new(big.Int)
		num.SetString(strInt, 10)
		num.Sub(num, &fullKey)
		message = append(message, rune(num.Int64()))
	}
	return string(message)
}
