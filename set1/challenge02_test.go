package set1_challenges

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// challenge 2 - Fixed XOR
func TestFixedXOR(t *testing.T) {
	input1_CH2 := "1c0111001f010100061a024b53535009181c"
	input2_CH2 := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	input1, err := hex.DecodeString(input1_CH2)
	if err != nil {
		t.Fatalf("Error on decoding input1")
	}
	input2, err := hex.DecodeString(input2_CH2)
	if err != nil {
		t.Fatalf("Error on decoding input2")
	}

	output, err := FixedXOR(input1, input2)
	if err != nil {
		t.Fatalf("Error on calculate XOR: %s", err)
	}

	output_CH2 := hex.EncodeToString(output)

	fmt.Println("input1:", input1_CH2)
	fmt.Println("input2:", input2_CH2)
	fmt.Println("expected:", expected)
	fmt.Println("result:", output_CH2)

	if output_CH2 != expected {
		t.Fatalf("Result different as expected:")
	}

}
