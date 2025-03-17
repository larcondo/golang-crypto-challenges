package set1_challenges

import (
	"fmt"
	"testing"
)

// challenge 1 - Convert hex to base64
func TestHexTobase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	encoded := HexToBase64(input)

	fmt.Println("Result  :", encoded)
	fmt.Println("Expected:", expected)

	if encoded != expected {
		t.Fatalf("Result string %s doesn't match with expected: %s", encoded, expected)
	}
}
