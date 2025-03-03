package set1_challenges

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

// challenge 4
func readFile(filepath string) (string, error) {
	if filepath == "" {
		return "", errors.New("empty file path")
	}

	data, err := os.ReadFile("./4.txt")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func mostRepeatedChar(msg []byte) []byte {
	list := make(map[byte]int)
	for i := range msg {
		list[msg[i]] += 1
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

func detectSingleCharXOR() []string {
	probable := []byte("etaoin shrdlu")
	// probable := []byte("ETAOIN SHRDLU")
	fileContent, err := readFile("./4.txt")
	check(err)

	lines := strings.Split(fileContent, "\n")

	fmt.Println("Lines:", len(lines))

	var messages []string

	for j := range len(lines) {
		// fmt.Printf("Linea %d: %s\n", j, lines[j])
		decode, err := hex.DecodeString(lines[j])
		check(err)

		charOrder := mostRepeatedChar(decode)

		for i := range probable {
			clave := charOrder[0] ^ probable[i]
			output := singleByteXORDecypher(decode, clave)
			msg := strings.ReplaceAll(output, "\n", "")
			if isValidMessage(msg) {
				fmt.Println("Mensaje encontrado:", msg)
				fmt.Println("Linea:", j, "String:", lines[j])
				fmt.Println("Clave:", string(clave))
				fmt.Println("Letra más repetida:", string(probable[i]))
				messages = append(messages, msg)
			}
		}
	}
	return messages
}
