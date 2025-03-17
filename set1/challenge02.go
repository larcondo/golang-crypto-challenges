package set1_challenges

import "errors"

// challenge 2 - Fixed XOR
func FixedXOR(in1 []byte, in2 []byte) ([]byte, error) {
	var result []byte

	if len(in1) != len(in2) {
		return result, errors.New("buffer con diferentes longitudes")
	}

	if len(in1) < 1 {
		return result, errors.New("buffer vacio")
	}

	for i := range in1 {
		result = append(result, in1[i]^in2[i])
	}

	return result, nil
}
