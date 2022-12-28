package totkcli

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/curve25519"
)

type PrivateKey []byte

func (k PrivateKey) String() string {
	return hex.EncodeToString(k)
}

func (k PrivateKey) ID() string {
	pub, _ := curve25519.X25519(curve25519.Basepoint, k)
	return hex.EncodeToString(pub)
}

func (k PrivateKey) Totk(pub []byte) (string, error) {
	return Totk(k, pub)
}

func NewKey() PrivateKey {
	k := make([]byte, 32)
	rand.Read(k)
	return k
}

func KeyFromHex(s string) (PrivateKey, error) {
	if len(s) != 64 {
		return nil, errors.New("Private key hex must be 64")
	}

	k, err := hex.DecodeString(s)

	if err != nil {
		return nil, err
	}

	return k, nil
}
