package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/sammygrey/DHKey.go/utils"
)

func inputBig(reader bufio.Reader, str string) big.Int {
	//create a new big int object and assign user input to it
	bigInput := new(big.Int)
	fmt.Println(str)
	strInput, _ := reader.ReadString('\n')
	strInput = strings.TrimSpace(strInput)
	bigInput, _ = bigInput.SetString(strInput, 10)
	return *bigInput
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("This program will walk you through the basic usage of a Diffie Hellman Key Exchange.")

	//take in user private key
	privateKey := inputBig(*reader, "First, type in your private key (integer), this can be any valid integer")

	//take in publically shared base and modulo
	bm := utils.NewBaseModulo()
	exampleBase, exampleModulo := bm[0], bm[1]
	fmt.Printf("Next, we need a public base and modulo for our encryption process, prime numbers that both parties have agreed upon.\n For an example base and modulo you can use: %d and %d\n", &exampleBase, &exampleModulo)
	base := inputBig(*reader, "Please enter your shared base:")
	modulo := inputBig(*reader, "Please enter your shared modulo:")

	//create new endpoint using parts provided above
	end := utils.NewEndpoint(base, modulo, privateKey)
	var publicKey big.Int = utils.GenPartial(end)
	fmt.Printf("Your new DHKey Endpoint has been generated!:\nPublic Base: %d\nPublic Modulo: %d\nPublic Key: %d\nPrivate Key: %d\nMake sure to only share the details marked public and not your private key!\n", &end.PublicBase, &end.PublicModulo, &publicKey, &end.PrivateKey)
	exampleEnd := utils.NewEndpoint(exampleBase, exampleModulo, *big.NewInt(6))

	//create example endpoint for single user testing
	var examplePublicKey big.Int = utils.GenPartial(exampleEnd)
	var thirdPartyKey big.Int
	//take in third party key from user, example public key is acceptable
	fmt.Printf("For this next part you need someone else's public key. You can have mine for example: %d\n", &examplePublicKey)
	fmt.Scanf("%d\n", &thirdPartyKey)

	//take in message from user and encrypt it
	fmt.Println("Now, just type a message for it to be encrypted!")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)
	exampleEncrypted := utils.Encrypt(exampleEnd, publicKey, "You rock!")
	encrypted := utils.Encrypt(end, publicKey, message)
	fmt.Printf("Your encrypted message is as follows: %s\n", encrypted)

	//take in outside encrypted message from user, example encrypted message is acceptable
	fmt.Printf("Now, you can decrypt the message of a friend using the same base and modulo as you, using their public key!\nEnter the message a friend has sent you or you can use my example here provided you used the details I gave above: %s\n", exampleEncrypted)
	outsideEncrypted, _ := reader.ReadString('\n')
	outsideEncrypted = strings.TrimSpace(outsideEncrypted)
	fmt.Printf("Using the details you entered above we can now decrypt this message to as follows: %s\n", utils.Decrypt(end, thirdPartyKey, outsideEncrypted))
}
