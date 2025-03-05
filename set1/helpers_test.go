package set1_challenges

import "testing"

func TestIsValidMessage(t *testing.T) {
	tests := []struct {
		text string
		want bool
	}{
		{text: "Hola mundo", want: true},
		{text: "Cooking MC's like pound bacon.", want: true},
		{text: "MICHAEL.", want: true},
		{text: "<^shgDSA#%", want: false},
		{text: "marshal01", want: false},
		{text: "Hola, mundo!?", want: true},
	}

	for _, v := range tests {
		res := isValidMessage(v.text)
		if res != v.want {
			t.Fatalf("isVM(%v) = %v, want %v\n", v.text, res, v.want)
		}
	}
}
