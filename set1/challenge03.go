package set1_challenges

// challenge 3 - Single-byte XOR cipher
func singleByteXORDecypher(encrypted []byte, key byte) string {
	var result []byte

	for i := range encrypted {
		result = append(result, encrypted[i]^key)
	}

	return string(result)
}
