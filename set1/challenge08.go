package set1_challenges

import "slices"

func Is_AES_in_ECB(ciphertext []byte, keySize int) bool {
	var blocks []string

	for i := 0; i < len(ciphertext); i += keySize {
		block := ciphertext[i : i+keySize]
		sBlock := string(block)
		if slices.Contains(blocks, sBlock) {
			return true
		}
		blocks = append(blocks, sBlock)
	}

	// si no se detecta ninguna repetición
	return false
}
