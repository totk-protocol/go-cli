package totkcli

import (
	"crypto/ed25519"
	"encoding/binary"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/curve25519"
)

func Totk(priv []byte, pub []byte) (string, error) {
	secret, err := curve25519.X25519(priv, pub)
	if err != nil {
		return "", err
	}

	now := getTimestamp()
	privkey := make([]byte, 64)
	copy(privkey, secret)
	// copy(privkey[32:], secret)
	sig := ed25519.Sign(privkey, now)
	off := sig[len(sig)-1] & 0xf
	n := binary.BigEndian.Uint32(sig[off:off+4]) % 1000000
	s := strconv.Itoa(int(n))
	if len(s) == 6 {
		return s, nil
	} else {
		return strings.Repeat("0", 6-len(s)) + s, nil
	}
}

func getTimestamp() []byte {
	now := time.Now().Unix() / 30
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(now))
	return buf
}
