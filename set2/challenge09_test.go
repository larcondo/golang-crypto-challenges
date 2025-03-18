package set2_challenges

import (
	"slices"
	"testing"
)

func TestPKCS7_padding(t *testing.T) {
	tests := []struct {
		text        []byte
		want        []byte
		blockLength int
	}{
		{
			text:        []byte("YELLOW SUBMARINE"),
			want:        []byte("YELLOW SUBMARINE\x04\x04\x04\x04"),
			blockLength: 20,
		},
		{
			text:        []byte("YELLOW SUBMARI"),
			want:        []byte("YELLOW SUBMARI\x02\x02"),
			blockLength: 16,
		},
		{
			text:        []byte("YELLOW SUBMARINEYELLOW SUB"),
			want:        []byte("YELLOW SUBMARINEYELLOW SUB\x06\x06\x06\x06\x06\x06"),
			blockLength: 16,
		},
	}

	for i := range tests {
		output := PKCS7_padding(tests[i].text, tests[i].blockLength)

		if slices.Compare(output, tests[i].want) != 0 {
			t.Fatalf("\nPKCS7_padding of %v with block length = %v: %v, wanted %v", string(tests[i].text), tests[i].blockLength, string(output), string(tests[i].want))
		}
	}

}
