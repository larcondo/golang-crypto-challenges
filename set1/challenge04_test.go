package set1_challenges

import (
	"fmt"
	"testing"
)

func TestDetectSingleCharXOR(t *testing.T) {

	msgs := detectSingleCharXOR()

	fmt.Println(msgs)

	if len(msgs) < 1 {
		fmt.Println("No se detectaron mensajes válidos.")
	}
}
