package tests

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/sammygrey/DHKey.go/utils"
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

func BenchmarkNewBaseModulo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		utils.NewBaseModulo()
	}
}

func BenchmarkEncrypt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		utils.Encrypt(test_end, *big.NewInt(8), "You rock!")
	}
}

func BenchmarkDecrypt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		utils.Decrypt(test_end, *big.NewInt(8), "109,131,137,52,134,131,119,127,53")
	}
}
