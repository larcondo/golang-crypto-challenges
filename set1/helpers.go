package set1_challenges

import (
	"regexp"
)

func isValidMessage(msg string) bool {
	pattern := regexp.MustCompile("[a-zA-Z .!?,']")
	bb := pattern.FindAllString(msg, -1)
	return len(bb) == len(msg)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
