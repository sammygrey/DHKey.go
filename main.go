package main

import (
	"fmt"

	"github.com/sammygrey/DHKey.go/utils"
)

func Error(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fmt.Println("This program will walk you through the basic usage of a Diffie Hellman Key Exchange.")
	var privateKey int
	fmt.Println("First, type in your private key (integer), this can be any valid integer")
	fmt.Scanf("%d", &privateKey)
	var bm []int = utils.NewBaseModulo()
	base, modulo := bm[0], bm[1]
	end := utils.NewEndpoint(base, modulo, privateKey)
	fmt.Printf("Your new DHKey Endpoint:\nPublic Base: %d\nPublic Modulo: %d\nPrivate Key: %d\n Make sure to only share your base and modulo and not your private key!\n", end.PublicBase, end.PublicModulo, end.PrivateKey)
	var publicKey int = utils.GenPartial(end)
	fmt.Printf("The public key generated from your private key is: %d\n", publicKey)
	//fmt.Println("For this next part you need someone else's public key. You can have mine for example: ")

	//fmt.Scanf("%d", &publicKey)
	//var message string
	//fmt.Println("Now, just type a message for it to be encrypted:")
	//fmt.Scanln(&message)
}
