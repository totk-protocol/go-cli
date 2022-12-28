package internal

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"totkcli"
)

type RunnerOption struct {
	Store  string
	Action int
	Id     string
	Target string
}

const (
	ACT_CALC = iota
	ACT_GENERATE
	ACT_LIST
	ACT_DELETE
)

func Execute(opt *RunnerOption) error {
	s, err := Open(opt.Store)

	if err != nil {
		return err
	}

	switch opt.Action {
	case ACT_CALC:
		return execCalc(s, opt.Id, opt.Target)
	case ACT_GENERATE:
		return execGenerate(s)
	case ACT_LIST:
		return execList(s)
	case ACT_DELETE:
		return execDel(s, opt.Target)
	default:
		return errors.New("Unknown action")
	}
}

func execCalc(s *Store, id string, target string) error {
	// get len
	var key totkcli.PrivateKey
	// Select default key
	if len(id) == 0 {
		lst, err := s.List()

		if err != nil {
			return err
		}

		if len(lst) == 0 {
			return errors.New("No Keyparis, Please add one")
		}

		key = lst[0]
	} else if len(id) == 64 {
		k, err := s.Get(id)

		if err != nil {
			return err
		}

		key = k
	} else {
		return errors.New("Invalid ID length")
	}

	if len(target) != 64 {
		return errors.New("Invalid public key length")
	}

	pub, err := hex.DecodeString(target)

	if err != nil {
		return err
	}

	secret, err := key.Totk(pub)

	if err != nil {
		return err
	}

	fmt.Printf("%v\n", secret)
	return nil
}

func execGenerate(s *Store) error {
	key := totkcli.NewKey()

	err := s.Add(key)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Key ID:\n")
	fmt.Printf(key.ID())
	return nil
}

func execList(s *Store) error {
	keys, err := s.List()

	if err != nil {
		return err
	}

	for _, k := range keys {
		fmt.Printf("%v\n", k.ID())
	}

	return nil
}

func execDel(s *Store, target string) error {
	if len(target) != 64 {
		return errors.New("Invalid ID length")
	}

	return s.Del(target)
}
