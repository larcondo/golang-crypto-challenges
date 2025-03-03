package set1_challenges

import (
	"encoding/hex"
	"fmt"
	"sort"
	"testing"
)

// challenge 3 - Single-byte XOR cipher
func TestSingleByteXORDecypher(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	// probable := []byte("ETAOIN SHRDLU")
	probable := []byte("etaoin shrdlu")

	decoded, err := hex.DecodeString(input)
	if err != nil {
		t.Fatal("Error on decode input")
	}

	list := make(map[byte]int)

	// create list of character appareance
	for i := range decoded {
		list[decoded[i]] += 1
	}
	// fmt.Println(list)

	// sort
	keys := make([]byte, 0, len(list))
	for key := range list {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return list[keys[i]] > list[keys[j]]
	})
	// fmt.Println(keys)

	fmt.Println(decoded)

	for i := range probable {
		// keys[1] se repite 2 veces en los primeros valores
		clave := keys[1] ^ probable[i]
		output := singleByteXORDecypher(decoded, clave)
		isValid := isValidMessage(output)
		if isValid {
			fmt.Println("Letra:", string(probable[i]), "Clave:", clave, "(", string(clave), ")", " -> Mensaje:", output)
		}
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

	msg := singleByteXORDecypher(decoded, key[0])
	fmt.Println(msg)
}
