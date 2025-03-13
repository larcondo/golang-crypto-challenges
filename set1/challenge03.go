package set1_challenges

import (
	"fmt"
	"regexp"
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

// var letterFrequencies = map[string]float64{
// 	"A": 8.167,
// 	"B": 1.492,
// 	"C": 2.782,
// 	"D": 4.253,
// 	"E": 12.702,
// 	"F": 2.228,
// 	"G": 2.015,
// 	"H": 6.094,
// 	"I": 6.966,
// 	"J": 0.153,
// 	"K": 0.772,
// 	"L": 4.025,
// 	"M": 2.406,
// 	"N": 6.749,
// 	"O": 7.507,
// 	"P": 1.929,
// 	"Q": 0.095,
// 	"R": 5.987,
// 	"S": 6.327,
// 	"T": 9.056,
// 	"U": 2.758,
// 	"V": 0.978,
// 	"W": 2.360,
// 	"X": 0.150,
// 	"Y": 1.974,
// 	"Z": 0.074,
// }

func isValidChar(s string) bool {
	return regexp.MustCompile("^[a-zA-Z ]$").MatchString(s)
}

func rankString(text []byte) float64 {
	points := 0.0

	for i := range text {
		char := string(text[i])
		if isValidChar(char) {
			// points += letterFrequencies[strings.ToUpper(char)]
			points += 1.0
			// } else {
			// 	points -= 10.0 // si hay un char no valido reinicio los points
		}
	}

	return points
}

type keyInfo struct {
	key    byte
	points float64
}

func BestKey(cipherText []byte) byte {
	possibleKeys := make([]byte, 256)
	for i := range possibleKeys {
		possibleKeys[i] = byte(i)
	}
	best := keyInfo{key: byte(0), points: 0.0}

	for i := range possibleKeys {
		dText := singleByteXORDecypher(cipherText, possibleKeys[i])
		currentPoints := rankString([]byte(dText))
		if currentPoints > best.points {
			best.key = possibleKeys[i]
			best.points = currentPoints
		}
	}

	return best.key
}
