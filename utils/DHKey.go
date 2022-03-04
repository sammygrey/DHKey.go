package DHKey

type Endpoint struct {
	publicBase   int //int64 - the public base known to both people
	publicModulo int // - the public modulus known to both people
	privateKey   int // - the private key known only to each person
}

func newEndpoint(publicBase, publicModulo, privateKey int) Endpoint {
	//creates a endpoint struct using an oop style function
	return Endpoint{publicBase, publicModulo, privateKey}
}

func genPartial(end Endpoint) int {
	//generate public key using private key and public parts to hand over to other party
	//this is safe to directly hand over
	partial := Exp(end.publicBase, end.privateKey)
	return partial % end.publicModulo
}

func genFull(end Endpoint, partialKey int) int {
	//generate full shared secret using the other parties' public key and our personal endpoint
	//this should not be shared directly
	return Exp(partialKey, end.privateKey) % end.publicModulo
}

func encrypt(end Endpoint, partialKey int, message string) string {
	var encrypted []rune //int32
	//encode each character to an integer, add the resultant int value of the secret to encode it
	for i, char := range message {
		encrypted[i] = rune(int(char) + genFull(end, partialKey))
	}
	return string(encrypted)

}

func decrypt(end Endpoint, partialKey int, encrypted string) string {
	var message []rune
	//perform the opposite of what we did to encrypt, to decrypt it
	for i, char := range encrypted {
		message[i] = rune(int(char) - genFull(end, partialKey))
	}
	return string(message)
}

func Exp(x, y int) int {
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
