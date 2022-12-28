package totkcli

import (
	"crypto/rand"
	"os"
	"testing"

	"golang.org/x/crypto/curve25519"
)

func TestTotk(t *testing.T) {
	p1 := make([]byte, 32)
	rand.Read(p1)
	p2, err := curve25519.X25519(curve25519.Basepoint, p1)

	if err != nil {
		t.Fail()
	}

	res, err := Totk(p1, p2)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)

	dir, err := os.UserConfigDir()

	if err != nil {
		t.Fatal(err)
	}
	t.Log(dir)
}
