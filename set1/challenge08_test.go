package set1_challenges

import (
	"strings"
	"testing"
)

func TestIs_AES_in_ECB(t *testing.T) {
	text := []byte("YELLOW SUBMARINEYELLOW SUBMARINE")
	KEY := []byte("YELLOW SUBMARINE")

	ciphertext, err := Encrypt_AES128_ECB(text, KEY)
	if err != nil {
		t.Fatal("Error on encrypt.")
	}
	result := Is_AES_in_ECB(ciphertext, len(KEY))

	if !result {
		t.Fatal("Text hasn't been encrypted in ECB.")
	}
	t.Log("Text has been encrypted in ECB")
}

func TestDetect_AES_in_ECB(t *testing.T) {
	file := "./8.txt"
	keySize := 16

	content, err1 := readFile(file)
	if err1 != nil {
		t.Fatalf("Error en leer archivo %v", file)
	}

	lines := strings.Split(content, "\n")

	var textsInECB []string

	for i := range lines {
		decode := DecodeBase64(lines[i])
		if Is_AES_in_ECB(decode, keySize) {
			textsInECB = append(textsInECB, lines[i])
		}
	}

	if len(textsInECB) < 1 {
		t.Fatal("No texts in ECB Mode.")
	}
	if len(textsInECB) > 1 {
		t.Fatal("More than 1 text in ECB Mode.")
	}

	t.Log(textsInECB[0])
}
