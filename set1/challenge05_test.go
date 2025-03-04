package set1_challenges

import (
	"encoding/hex"
	"testing"
)

type test5value struct {
	key  string
	text string
	want string
}

func TestRepeatingKeyXOR(t *testing.T) {

	test5 := test5value{
		key: "ICE",
		text: `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`,
		want: `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`,
	}

	encrypted := repeatingKeyXOR(test5.text, test5.key)

	encoded := hex.EncodeToString(encrypted)

	if encoded != test5.want {
		t.Fatalf("Result: %v\nWant: %v", encoded, test5.want)
	}

}
