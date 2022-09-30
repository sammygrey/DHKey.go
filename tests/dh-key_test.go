package tests

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/sammygrey/dh-key.go/utils"
)

var test_end utils.Endpoint = utils.NewEndpoint(*big.NewInt(5), *big.NewInt(23), *big.NewInt(15))
var test_end2 utils.Endpoint = utils.NewEndpoint(*big.NewInt(5), *big.NewInt(23), *big.NewInt(6))

func TestGenPartial(t *testing.T) {
	var tests = []struct {
		endpoint utils.Endpoint
		expected big.Int
	}{
		{test_end, *big.NewInt(19)},
		{test_end2, *big.NewInt(8)},
	}

	for _, test := range tests {
		partialKey := utils.GenPartial(test.endpoint)
		if partialKey.Cmp(&test.expected) != 0 {
			t.Error("TEST FAILED: " + fmt.Sprint(test.expected) + " expected, but received " + fmt.Sprint(partialKey))
		}
	}

}

func TestGenFull(t *testing.T) {
	var tests = []struct {
		endpoint   utils.Endpoint
		partialKey big.Int
		expected   big.Int
	}{
		{test_end, *big.NewInt(8), *big.NewInt(2)},
		{test_end2, *big.NewInt(19), *big.NewInt(2)},
	}

	for _, test := range tests {
		fullKey := utils.GenFull(test.endpoint, test.partialKey)
		if fullKey.Cmp(&test.expected) != 0 {
			t.Error("TEST FAILED: " + fmt.Sprint(test.expected) + " expected, but received " + fmt.Sprint(fullKey))
		}
	}
}

func BenchmarkNewBaseModulo16(b *testing.B) {
	for n := 0; n < b.N; n++ {
		utils.NewBaseModulo(16)
	}
}

func BenchmarkNewBaseModulo24(b *testing.B) {
	for n := 0; n < b.N; n++ {
		utils.NewBaseModulo(24)
	}
}

func BenchmarkNewBaseModulo32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		utils.NewBaseModulo(32)
	}
}

func BenchmarkNewPrivateKey16(b *testing.B){
	for n := 0; n < b.N; n++ {
		utils.NewPrivateKey(16)
	}
}

func BenchmarkNewPrivateKey24(b *testing.B){
	for n := 0; n < b.N; n++ {
		utils.NewPrivateKey(24)
	}
}

func BenchmarkNewPrivateKey32(b *testing.B){
	for n := 0; n < b.N; n++ {
		utils.NewPrivateKey(32)
	}
}

//figure out how to change these to better fit the new functions

var test_bm []big.int = utils.NewBaseModulo(32)
var test_pk big.int = utils.NewPrivateKey(32)
var test_pk2 big.int = utils.NewPrivateKey(32)
var test_end3 utils.Endpoint = utils.NewEndpoint(test_bm[0], test_bm[1], test_pk)
var test_end4 utils.Endpoint = utils.NewEndpoint(test_bm[0], test_bm[1], test_pk2)
var test_pbk big.int = utils.genPartial(test_end3)
var test_pbk2 big.int = utils.genPartial(test_end4)

func BenchmarkEncrypt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		utils.Encrypt(test_end3, test_pbk2, "You rock!")
	}
}

var test_cipherText string = utils.Encrypt(test_end3, test_pbk2, "You rock!")

func BenchmarkDecrypt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		utils.Decrypt(test_end4, test_pbk, test_cipherText)
	}
}
