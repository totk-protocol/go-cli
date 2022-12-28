package internal

import (
	"errors"
	"totkcli"

	"github.com/tidwall/buntdb"
)

type Store struct {
	db *buntdb.DB
}

func Open(path string) (*Store, error) {
	db, err := buntdb.Open(path)

	if err != nil {
		return nil, err
	}

	return &Store{db: db}, nil
}

func (s *Store) Get(id string) (totkcli.PrivateKey, error) {
	if len(id) != 64 {
		return nil, errors.New("Invalid id length")
	}

	var key totkcli.PrivateKey
	err := s.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get("key:" + id)

		if err != nil {
			return err
		}

		key, err = totkcli.KeyFromHex(val)

		return err
	})

	return key, err
}

func (s *Store) Add(key totkcli.PrivateKey) error {
	if len(key) != 32 {
		return errors.New("Invalid key length")
	}

	return s.db.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Get("key:" + key.ID())

		if errors.Is(err, buntdb.ErrNotFound) {
			_, _, err := tx.Set("key:"+key.ID(), key.String(), nil)

			if err != nil {
				return err
			}
		} else {
			return err
		}

		return nil
	})
}

func (s *Store) Del(id string) error {
	if len(id) != 64 {
		return errors.New("Invalid key length")
	}

	return s.db.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete("key:" + id)
		return err
	})
}

func (s *Store) List() ([]totkcli.PrivateKey, error) {
	keys := make([]totkcli.PrivateKey, 0)

	err := s.db.View(func(tx *buntdb.Tx) error {
		var err error = nil

		tx.AscendKeys("key:*", func(key string, value string) bool {
			k, e := totkcli.KeyFromHex(value)

			if e != nil {
				err = e
				return false
			}

			keys = append(keys, k)
			return true
		})

		return err
	})

	return keys, err
}
