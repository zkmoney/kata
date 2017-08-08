package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type withStack struct {
	err   error
	stack int
}

func (w *withStack) Error() string {
	return w.err.Error()
}

func wrap(err error) error {
	return &withStack{
		err:   err,
		stack: 1,
	}
}

func main() {
	err := foo()
	fmt.Println(err)
	// fmt.Printf("String verb: %s\n", err)

	fmt.Printf("Verbose verb: %n\n", err)
}

func foo() error {
	if err := bar(); err != nil {
		return errors.Wrap(err, "call bar, got an error")
	}
	return nil
}

func bar() error {
	err := errors.New("error ahhhhhhhhhhh!!!!")
	return err
}
