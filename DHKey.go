package DHKey

type Endpoint struct {
	publicBase   int //int64 - the public base known to both people
	publicModulo int // - the public modulus known to both people
	privateKey   int // - the private key known only to each person
}

func newEndpoint(publicBase, publicKey2, privateKey int) Endpoint {
	return Endpoint{publicBase, publicKey2, privateKey}
}

func genPartial(end Endpoint) int {
	//This partial key is also known as the personal public key in terms of this algorithm
	partial := Exp(end.publicBase, end.privateKey)
	return partial % end.publicModulo
}

func genFull(end Endpoint, partialKey int) int {
	return Exp(partialKey, end.privateKey) % end.publicModulo
}

func encrypt(end Endpoint, partialKey int, message string) string {
	var encrypted []rune //int32
	for i, char := range message {
		encrypted[i] = rune(int(char) + genFull(end, partialKey))
	}
	return string(encrypted)

}

func decrypt(end Endpoint, partialKey int, encrypted string) string {
	var message []rune
	for i, char := range encrypted {
		message[i] = rune(int(char) - genFull(end, partialKey))
	}
	return string(message)
}

// fast binary operation for calculating power of integers, obviously only works to int64 max
func Exp(x, y int) int {
	power := 1
	for y > 0 {
		if y&1 != 0 {
			power *= x
		}
		y >>= 1
		x *= x
	}
	return power
}
