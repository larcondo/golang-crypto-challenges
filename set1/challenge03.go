package set1_challenges

import (
	"fmt"
	"sort"
)

// challenge 3 - Single-byte XOR cipher
func singleByteXORDecypher(cipherText []byte, key byte) string {
	var result []byte

	for i := range cipherText {
		result = append(result, cipherText[i]^key)
	}

	return string(result)
}

func rankByteAppearance(cipherText []byte) []byte {
	list := make(map[byte]int)

	// create list of character appareance
	for i := range cipherText {
		list[cipherText[i]] += 1
	}

	// sort
	keys := make([]byte, 0, len(list))
	for key := range list {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return list[keys[i]] > list[keys[j]]
	})

	return keys
}

func DecypherText(cipherText []byte) string {
	// probable := []byte("ETAOIN SHRDLU")
	probable := []byte("etaoin shrdlu")

	keys := rankByteAppearance(cipherText)

	for i := range probable {
		// keys[1] se repite 2 veces en los primeros valores
		clave := keys[0] ^ probable[i]
		output := singleByteXORDecypher(cipherText, clave)
		if isValidMessage(output) {
			fmt.Println("Letra:", string(probable[i]), "Clave:", clave, "(", string(clave), ")", " -> Mensaje:", output)
			return output
		}
	}

	return ""
}
