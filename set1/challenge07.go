package set1_challenges

import "crypto/aes"

func Decrypt_AES128_ECB(ciphertext []byte, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	deciphertext := make([]byte, len(ciphertext))

	blockSize := len(key)

	for i := 0; i < len(ciphertext); i += blockSize {
		cipher.Decrypt(deciphertext[i:(i+blockSize)], ciphertext[i:(i+blockSize)])
	}

	return deciphertext, nil
}

func Encrypt_AES128_ECB(text []byte, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, len(text))
	blockSize := len(key)

	for i := 0; i < len(text); i += blockSize {
		cipher.Encrypt(ciphertext[i:(i+blockSize)], text[i:(i+blockSize)])
	}

	return ciphertext, nil
}
