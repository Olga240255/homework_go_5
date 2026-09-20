package cliargs

import (
	"errors"
	"io"
	"slices"
	"strconv"
	"strings"
)

// ErrUsage marks invalid command-line arguments.
var ErrUsage = errors.New("usage: args <name> [repeat]")

// Options contains validated user arguments.
type Options struct {
	Name   string
	Repeat int
}

// UserArgs returns a new slice without os.Args[0].
func UserArgs(all []string) []string {
	// TODO: implement according to docs/task.md.
	if all == nil || len(all) < 2 {
		return []string{}
	}
	return slices.Clone(all[1:])
}

// Parse validates the arguments passed after os.Args[0].
func Parse(args []string) (Options, error) {
	// TODO: implement according to docs/task.md.
	if args == nil || len(args) < 1 || len(args) > 2 {
		return Options{}, ErrUsage
	}
	var O Options

	O.Name = strings.TrimSpace(args[0])
	if O.Name == "" {
		return Options{}, ErrUsage
	}

	if len(args) == 1 {
		O.Repeat = 1
	} else {
		repeat, err := strconv.Atoi(args[1])
		if err != nil || repeat <= 0 {
			return Options{}, ErrUsage
		}
		O.Repeat = repeat
	}
	return O, nil
}

// Run writes the command result to the supplied writer.
func Run(args []string, out io.Writer) error {
	// TODO: implement according to docs/task.md.
	var err error
	O, err := Parse(args)
	if err != nil {
		return err
	}
	for i := 0; i < O.Repeat; i++ {
		_, err = out.Write([]byte("hello, "))
		_, err = out.Write([]byte(O.Name))
		_, err = out.Write([]byte("\n"))
	}
	if err != nil {
		return err
	}
	return nil
}

// Example returns deterministic output used by cmd and the example test.
func Example() string {
	var out strings.Builder
	_ = Run([]string{"Maria", "2"}, &out)
	return out.String()
}
