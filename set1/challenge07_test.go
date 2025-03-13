package set1_challenges

import "testing"

func TestDecrypt_AES128_ECB(t *testing.T) {
	file := "./7.txt"

	content, err1 := readFile(file)
	if err1 != nil {
		t.Fatalf("Error en leer archivo %v", file)
	}

	decoded := DecodeBase64(content)

	KEY := []byte("YELLOW SUBMARINE")

	result, err2 := Decrypt_AES128_ECB(decoded, KEY)
	if err2 != nil {
		t.Fatalf("Error en Decrypt_AES128_ECB: %v", err2)
	}

	t.Logf("Resultado:\n%v", string(result))

}
