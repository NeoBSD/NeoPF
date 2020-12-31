package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/neoBSD/neoPF/freebsd"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	machine, err := freebsd.GetSystemInfo()
	if err != nil {
		return err
	}

	err = printJSON(machine, false)
	return nil
}

func printJSON(obj interface{}, htmlEscape bool) error {
	if htmlEscape {
		js, err := json.Marshal(obj)
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, string(js))
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	err := enc.Encode(obj)
	if err != nil {
		return err
	}
	return nil
}
