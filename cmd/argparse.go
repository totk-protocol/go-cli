package main

import (
	"flag"
	"os"
	"path"
	"totkcli/internal"
)

func parseArg() (*internal.RunnerOption, error) {
	var isList, isAdd, isDel bool
	var key, target, store string

	flag.BoolVar(&isList, "l", false, "Show key list")
	flag.BoolVar(&isList, "list", false, "Show key list")
	flag.BoolVar(&isAdd, "g", false, "Show key list")
	flag.BoolVar(&isAdd, "gen", false, "Show key list")
	flag.BoolVar(&isDel, "d", false, "Show key list")
	flag.BoolVar(&isDel, "del", false, "Show key list")
	flag.StringVar(&key, "k", "", "Key id")
	flag.StringVar(&key, "key", "", "Key id")
	flag.StringVar(&store, "s", "", "Database file")
	flag.StringVar(&store, "store", "", "Database file")

	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		target = args[0]
	}

	if len(store) == 0 {
		dir, err := os.UserConfigDir()

		if err != nil {
			return nil, err
		}

		store = path.Join(dir, "totk", "keypair.db")
	}

	if isList {
		return &internal.RunnerOption{
			Action: internal.ACT_LIST,
			Store:  store,
		}, nil
	}

	if isAdd {
		return &internal.RunnerOption{
			Action: internal.ACT_GENERATE,
			Store:  store,
		}, nil
	}

	if isDel {
		return &internal.RunnerOption{
			Action: internal.ACT_DELETE,
			Store:  store,
			Target: target,
		}, nil
	}

	return &internal.RunnerOption{
		Action: internal.ACT_CALC,
		Id:     key,
		Target: target,
		Store:  store,
	}, nil
}
