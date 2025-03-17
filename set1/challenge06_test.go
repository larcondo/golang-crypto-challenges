package set1_challenges

import (
	"encoding/hex"
	"fmt"
	"testing"
)

type tValue struct {
	x    byte
	want int
}

// OK
func TestDecodeBase64(t *testing.T) {
	// Verificado con valores del challenge 1 (proceso inverso)
	input := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	want := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	output := DecodeBase64(input)

	wantBytes, err := hex.DecodeString(want)
	ResolveErr(t, err)

	// hexa := hex.EncodeToString(output)

	if string(output) != string(wantBytes) {
		t.Fatalf("base64toHex(%v) = %v, want %v", input, string(output), string(wantBytes))
	}
}

// OK
func TestLookup(t *testing.T) {
	lt := lookupTable()

	tests := []tValue{
		{x: byte(1), want: 1},
		{x: byte(2), want: 1},
		{x: byte(3), want: 2},
		{x: byte(4), want: 1},
		{x: byte(5), want: 2},
		{x: byte(6), want: 2},
		{x: byte(7), want: 3},
		{x: byte(10), want: 2},
		{x: byte(11), want: 3},
		{x: byte(200), want: 3},
		{x: byte(210), want: 4},
		{x: byte(250), want: 6},
		{x: byte(251), want: 7},
		{x: byte(255), want: 8},
	}

	for _, tt := range tests {
		if lt[tt.x] != tt.want {
			t.Fatalf("lookup[%v] = %v, want %v", tt.x, lt[tt.x], tt.want)
		}
	}
}

// OK
func TestHammingDistance(t *testing.T) {
	texts := []string{
		"this is a test",
		"wokka wokka!!!",
	}

	want := 37
	res := HammingDistance(texts[0], texts[1])

	if res != want {
		t.Fatalf("HammingDistance(%v, %v) = %v, want %v", texts[0], texts[1], res, want)
	}
}

// OK
func TestGetBestKeySize(t *testing.T) {
	KEY := "ICE"
	STANZA := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	encrypted := RepeatingKeyXOR(STANZA, KEY)

	keysize, err := GetBestKeySize(encrypted)
	if err != nil {
		t.Errorf("Error inesperado %v", err)
	}

	if keysize != 3 {
		t.Errorf("Valor esperado de keysize: %v, obtenido %v", 3, keysize)
	}
}

func TestGetCompleteKey(t *testing.T) {
	testKey := "ICE"
	testText := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	encrypted := RepeatingKeyXOR(testText, testKey)

	keysize, err := GetBestKeySize(encrypted)
	ResolveErr(t, err)

	if keysize != 3 {
		t.Fatalf("Keysize: %v, Want: %v\n", keysize, 3)
	}

	key := GetCompleteKey(encrypted, keysize)

	if string(key) != testKey {
		t.Fatalf("Key: %v, Want: %v\n", string(key), testKey)
	} else {
		fmt.Printf("Keysize obtenido: %v\nKey obtenida: %v", keysize, string(key))
	}
}

func TestBreakRepeatingKeyXOR(t *testing.T) {
	content, err := readFile("./6.txt")
	ResolveErr(t, err)
	ciphertext := DecodeBase64(content)

	keysize, err := GetBestKeySize(ciphertext)
	ResolveErr(t, err)

	key := GetCompleteKey(ciphertext, keysize)

	decrypted := RepeatingKeyXOR(string(ciphertext), string(key))

	fmt.Println("Best Key Size:", keysize)
	fmt.Println(string(key))
	fmt.Println(string(decrypted))
}
