package set1_challenges

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// challenge 3 - Single-byte XOR cipher
func TestSingleByteXORDecypher(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	decoded, err := hex.DecodeString(input)
	if err != nil {
		t.Fatal("Error on decode input")
	}

	resultado := DecypherText(decoded)
	if resultado == "" {
		t.Fatal("No se obtuvo el resultado")
	}
}

// challenge 3 - Single-byte XOR cipher
func TestSingleWithKey(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	key := []byte("X")

	decoded, err := hex.DecodeString(input)
	if err != nil {
		t.Fatalf("Error on decode input")
	}

	msg := SingleByteXORDecypher(decoded, key[0])
	fmt.Println(msg)
}

func TestBestKey(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	decoded, err := hex.DecodeString(input)
	if err != nil {
		t.Fatal("Error on decode input")
	}

	key := BestKey(decoded)

	resultado := SingleByteXORDecypher(decoded, key)
	if resultado == "" {
		t.Fatal("No se obtuvo el resultado")
	} else {
		fmt.Println("Res:", resultado)
	}
}
