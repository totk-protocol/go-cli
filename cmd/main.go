package main

import (
	"fmt"
	"os"
	"path"
	"totkcli/internal"
)

func main() {
	opt, err := parseArg()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	dir := path.Dir(opt.Store)
	err = os.MkdirAll(dir, 0755)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	err = internal.Execute(opt)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
