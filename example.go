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

	bm, err1 := utils.NewBaseModulo(16)
	err(err1)
	pk, err2 := utils.NewPrivateKey(16)
	err(err2)
	pk2, err3 := utils.NewPrivateKey(16)
	err(err3)
	base, modulo := bm[0], bm[1]
	end := utils.NewEndpoint(base, modulo, pk)
	end2 := utils.NewEndpoint(base, modulo, pk2)
	var pbk big.Int = utils.GenPartial(end)
	var pbk2 big.Int = utils.GenPartial(end2)
	fmt.Println(pbk.String())
	//fullKey := utils.GenFull(end, pbk2)
	//fmt.Println(fullKey.String()) //this should not be visible to anyone but you
	plainText := "You Rock!"
	fmt.Println(plainText)
	cipherBytes, _ := utils.Encrypt(end, pbk2, plainText)
	fmt.Println(cipherBytes)
	cipherText := string(cipherBytes)
	fmt.Println(cipherText)
	plainText2, _ := utils.Decrypt(end2, pbk, cipherText)
	fmt.Println(string(plainText2))

}