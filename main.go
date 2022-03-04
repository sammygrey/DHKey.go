package main

import (
	"fmt"
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
	base, modulo := utils.newBaseModulo()
	end := utils.newEndpoint(base, modulo, privateKey)
	fmt.Printf("Your new DHKey Endpoint:\nPublic Base: %d\nPublic Modulo: %d\nPrivate Key: %d\n Make sure to only share your base and modulo and not your private key!\n", end.publicBase, end.publicModulo, end.privateKey)
	var publicKey int = utils.genPartial(end)
	fmt.Println(publicKey)
	//fmt.Println("For this next part you need someone else's public key. You can have mine for example: ")

	//fmt.Scanf("%d", &publicKey)
	//var message string
	//fmt.Println("Now, just type a message for it to be encrypted:")
	//fmt.Scanln(&message)
}
