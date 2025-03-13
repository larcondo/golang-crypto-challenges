package set1_challenges

import (
	"encoding/base64"
	"sort"
)

type op struct {
	keysize  int
	distance float32
}

func lookupTable() []int {
	table := make([]int, 256)
	table[0] = 0
	for i := range table {
		table[i] = (i & 1) + table[int(i/2)]
	}

	return table
}

func DecodeBase64(s string) []byte {
	output, err := base64.StdEncoding.DecodeString(s)
	check(err)
	return output
}

func HammingDistance(text1 string, text2 string) int {
	distance := 0

	lt := lookupTable()

	for i := range len(text1) {
		xor := text1[i] ^ text2[i]
		distance += lt[xor]
	}

	return distance
}

func GetBestKeySize(text []byte) (int, error) {
	// Options of KEYSIZE: 2 to 40
	options := make([]op, 39)

	initialDistance := float32(len(text))
	for i := range options {
		options[i] = op{
			keysize:  i + 2,
			distance: initialDistance,
		}
	}

	for i := range options {
		keysize := options[i].keysize
		var sum int
		pares := len(text)/(keysize*2) - 1

		if pares <= 2 {
			continue
		}

		for j := range pares {
			s1 := text[j*2*keysize : (j*2+1)*keysize]
			s2 := text[(j*2+1)*keysize : (j*2+2)*keysize]
			sum += HammingDistance(string(s1), string(s2))
		}

		distance := float32(sum) / float32(keysize)
		prom := distance / float32(pares)
		// fmt.Printf("Keysize: %v, Sum: %v, Prom.: %v\n", options[i].keysize, sum, prom)
		options[i].distance = prom
	}

	sort.Slice(options, func(i, j int) bool {
		return options[i].distance < options[j].distance
	})

	return options[0].keysize, nil
}

func GetCompleteKey(text []byte, keysize int) []byte {
	var totalBlocks int
	resto := len(text) % keysize

	if resto > 0 {
		totalBlocks = len(text)/keysize + 1
	} else {
		totalBlocks = len(text) / keysize
	}

	var cipherTextBlocks [][]byte

	for i := range totalBlocks {
		if i < totalBlocks-1 {
			cipherTextBlocks = append(cipherTextBlocks, text[i*keysize:(i+1)*keysize])
		} else {
			cipherTextBlocks = append(cipherTextBlocks, text[i*keysize:])
		}
	}

	trasposedBlocks := make([][]byte, keysize)

	for i := range keysize {
		var block []byte
		for j := range cipherTextBlocks {
			if i < len(cipherTextBlocks[j]) {
				block = append(block, cipherTextBlocks[j][i])
			}
		}
		trasposedBlocks[i] = block
	}

	var completeKey []byte

	for i := range trasposedBlocks {
		bk := BestKey(trasposedBlocks[i])
		completeKey = append(completeKey, bk)
		// btext := singleByteXORDecypher(trasposedBlocks[i], bk)
		// fmt.Printf("Index: %v, Key: %v, Text: %v\n", i, bk, btext)
	}

	return completeKey
}
