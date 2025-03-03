package set1_challenges

import "testing"

func TestIsValidMessage(t *testing.T) {
	messages := []string{
		"Hola mundo",
		"Cooking MC's like pound bacon.",
		"MICHAEL.",
	}

	if !isValidMessage(messages[0]) {
		t.Fatalf("El mensaje %s no es válido", messages[0])
	}
	if !isValidMessage(messages[1]) {
		t.Fatalf("El mensaje %s no es válido", messages[0])
	}
	if !isValidMessage(messages[2]) {
		t.Fatalf("El mensaje %s no es válido", messages[0])
	}

	wrongMsg := "<^shgDSA#%"
	if isValidMessage(wrongMsg) {
		t.Fatalf("El mensaje %s fue considerado válido cuando no lo es", wrongMsg)
	}

}
