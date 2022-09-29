package utils

import (
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"math/big"
	"errors"
	"io"
)

// Endpoint is ...
type Endpoint struct {
	PublicBase   big.Int //- the public base known to both people
	PublicModulo big.Int // - the public modulus known to both people
	PrivateKey   big.Int // - the private key known only to each person
}

// NewBaseModulo is ...
func NewBaseModulo(bytes uint8) ([]big.Int, error) {
	//ideally g**q = 1 mod p, where q is a random prime integer, but all prime numbers will work
	bits := int(bytes)*8
	if (bytes == 16 || bytes == 24 || bytes == 32){
		p, _ := rand.Prime(rand.Reader, bits)
		g, _ := rand.Prime(rand.Reader, bits)
		base, modulo := p.Abs(p), g.Abs(g)
		return []big.Int{*base, *modulo}, nil
	}
	return nil, errors.New("base and modulo may only be 16/24/32 bytes")

}

// NewEndpoint is ...
func NewEndpoint(publicBase, publicModulo, privateKey big.Int) Endpoint {
	//creates a endpoint struct using an oop style function
	return Endpoint{publicBase, publicModulo, privateKey}
}

func NewPrivateKey(bytes uint8) (big.Int, error) {
	privateKey := new(big.Int)
	if (bytes == 16 || bytes == 24 || bytes == 32){
		pkBytes := make([]byte, bytes)
		_, err := rand.Read(pkBytes)
		privateKey.SetBytes(pkBytes)
		return *privateKey, err
	}
	return *privateKey, errors.New("key may only be 16/24/32 bytes")
}

// GenPartial is ...
func GenPartial(end Endpoint) big.Int {
	//generate public key using private key and public parts to hand over to other party
	//this is safe to directly hand over
	partial := new(big.Int)
	partial.Exp(&end.PublicBase, &end.PrivateKey, &end.PublicModulo)
	return *partial
}

// GenFull is ...
func GenFull(end Endpoint, partialKey big.Int) big.Int {
	//generate full shared secret using the other parties' public key and our personal endpoint
	//this should not be shared directly
	full := new(big.Int)
	full.Exp(&partialKey, &end.PrivateKey, &end.PublicModulo)
	return *full
}

// Encrypt is ...
func Encrypt(end Endpoint, partialKey big.Int, message string) ([]byte, error) {
	plainText := []byte(message)
	fullKey := GenFull(end, partialKey)
	block, err := aes.NewCipher((&fullKey).Bytes())
	if err != nil {
        return nil, err
    }
	b := base64.StdEncoding.EncodeToString(plainText)
    cipherText := make([]byte, aes.BlockSize+len(b))
    initialVector := cipherText[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, initialVector); err != nil {
        return nil, err
    }
    cipherFeedback := cipher.NewCFBEncrypter(block, initialVector)
    cipherFeedback.XORKeyStream(cipherText[aes.BlockSize:], []byte(b))
	return cipherText, nil
}

// Decrypt is ...
func Decrypt(end Endpoint, partialKey big.Int,encrypted string) ([]byte, error) {
	fullKey := GenFull(end, partialKey)
	block, err := aes.NewCipher((&fullKey).Bytes())
	cipherText := []byte(encrypted)
	if err != nil {
        return nil, err
    }
    if len(cipherText) < aes.BlockSize {
        return nil, errors.New("cipher text is too short")
    }
    initialVector := cipherText[:aes.BlockSize]
    plainText := cipherText[aes.BlockSize:]
    cipherFeedback := cipher.NewCFBDecrypter(block, initialVector)
    cipherFeedback.XORKeyStream(plainText, plainText)
    data, err := base64.StdEncoding.DecodeString(string(plainText))
    if err != nil {
        return nil, err
    }
	return data, nil
}