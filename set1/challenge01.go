package set1_challenges

import (
	b64 "encoding/base64"
	"encoding/hex"
	"log"
)

// challenge 1 - Convert hex to base64
func HexToBase64(input string) string {
	decodedHex, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}
	encoded := b64.StdEncoding.EncodeToString(decodedHex)
	return encoded
}
