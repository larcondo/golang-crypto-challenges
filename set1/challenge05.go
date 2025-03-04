package set1_challenges

func repeatingKeyXOR(text string, key string) []byte {

	var j int
	var resultado []byte
	count := 1

	for i := range len(text) {
		if i >= (count * len(key)) {
			count++
			j = 0
		}
		resultado = append(resultado, text[i]^key[j])
		j++
	}

	return resultado
}
