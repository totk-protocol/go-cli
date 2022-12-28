package internal

import (
	"os"
	"testing"
	"totkcli"
)

func TestStore(t *testing.T) {
	s, err := Open("./test.db")

	defer func() {
		os.Remove("./test.db")
	}()

	if err != nil {
		t.Fatal(err)
	}

	k := totkcli.NewKey()
	err = s.Add(k)

	if err != nil {
		t.Fatal(err)
	}

	keys, err := s.List()

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v\n", keys)

	if len(keys) != 1 {
		t.Fail()
	}

	err = s.Del(keys[0].ID())

	if err != nil {
		t.Fatal(err)
	}

	keys, err = s.List()

	if err != nil {
		t.Fatal(err)
	}

	if len(keys) != 0 {
		t.Fail()
	}
}
