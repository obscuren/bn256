package bn256

import (
	"testing"

	"crypto/rand"
)

func TestRandomG1Marshal(t *testing.T) {
	for i := 0; i < 10; i++ {
		n, g1, err := RandomG1(rand.Reader)
		if err != nil {
			t.Error(err)
			continue
		}
		t.Logf("%d: %x\n", n, g1.Marshal())
	}
}
