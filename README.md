# DHKey.go

[![Go Report Card](https://goreportcard.com/badge/github.com/sammygrey/DHKey.go)](https://goreportcard.com/report/github.com/sammygrey/DHKey.go)

🔑🔁🔑 Go microservice that provides Diffie Hellman Key Exhange services

## Usage

### `main.go`

While Key-Exchanges are somewhat counter intuitive to perform personally, the example usage provided in main.go should give you a feeling of what the functionality in this program entails and how it can be used.
If you plan on using this program seriously I recommend automating the processes available in utils to encrypt/decrypt content.

### Importing Elsewhere

You can import the tools used for this project by getting "github.com/sammygrey/DHKey.go/utils" and adding it to the imports of your project.

### Features:

#### Random Base and Modulo Generation:

This program includes a function named `NewBaseModulo` to randomly generate a public base and modulo for use between two parties.
This can be used in place of user calculated variables or pre-determined bases and modulos.

#### Partial and Full Key Generation:

You can generate partial (public) and full (shared) keys for key exchange usage using the `GenPartial` and `GenFull` functions provided in utils.
These functions take in a preexisting endpoint to calculate these values.

#### Endpoint Creation

You can create Endpoint structures for class-like utility using this project.
You can create a new instance of the Endpoint struct using standard GoLang syntax or instantiate one using the `NewEndpoint` function
Below is an example Endpoint struct with an included public key

![example endpoint struct](https://user-images.githubusercontent.com/49354894/156836743-813ad31d-319e-4b57-b555-d971aeac484e.png)

#### Encryption and Decryption Using Keys

Using the keys and methods provided in the utilities named above you can encrypt strings and decrypt them using shared secrets.
If you want an in-depth analysis of the mathematics behind how this works I recommend [Computerphile's](https://www.youtube.com/watch?v=Yjrfm_oRO0w) and [Kahn Academy's](https://www.youtube.com/watch?v=M-0qt6tdHzk) videos on it.
