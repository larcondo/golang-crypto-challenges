package set1_challenges

import (
	"strings"
)

func isValidMessage(msg string) bool {
	mayusc := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	minusc := "abcdefghijklmnopqrstuvwxyz"
	symbol := " !?.,'"

	for i := range len(msg) {
		cond1 := strings.Contains(mayusc, string(msg[i]))
		cond2 := strings.Contains(minusc, string(msg[i]))
		cond3 := strings.Contains(symbol, string(msg[i]))

		if !(cond1 || cond2 || cond3) {
			return false
		}
	}

	return true
}
