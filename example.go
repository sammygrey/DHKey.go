package main

import (
	"fmt"
	"math/big"

	"github.com/sammygrey/dh-key.go/utils"
)

func err(err error){
	if err != nil{
		fmt.Println(err)
	}
}

func main() {

	bm, err1 := utils.NewBaseModulo(32)
	err(err1)
	pk1, err2 := utils.NewPrivateKey(32)
	err(err2)
	pk2, err3 := utils.NewPrivateKey(32)
	err(err3)
	base, modulo := bm[0], bm[1]
	end1 := utils.NewEndpoint(base, modulo, pk1)
	end2 := utils.NewEndpoint(base, modulo, pk2)
	var pbk1 big.Int = utils.GenPartial(end1)
	var pbk2 big.Int = utils.GenPartial(end2)
	fullKey := utils.GenFull(end1, pbk2)
	fmt.Println(&fullKey)
	message := "You Rock!"
	fmt.Println(message)
	cipherText, _ := utils.Encrypt(end1, pbk2, message)
	fmt.Println(cipherText)
	encrypted := string(cipherText)
	fmt.Println(encrypted)
	decrypted, _ := utils.Decrypt(end2, pbk1, encrypted)
	fmt.Println(decrypted)

}