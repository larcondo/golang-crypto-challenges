package set2_challenges

func PKCS7_padding(plaintext []byte, blocksize int) []byte {
	rest := len(plaintext) % blocksize

	if rest == 0 {
		return plaintext
	}

	paddNumber := blocksize - rest
	origLength := len(plaintext)

	paddedText := make([]byte, origLength+paddNumber)

	copy(paddedText, plaintext)

	for i := range paddNumber {
		paddedText[i+origLength] = byte(paddNumber)
	}

	return paddedText
}
