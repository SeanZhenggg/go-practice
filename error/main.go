package main

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/xerrors"
)

func main() {
	err := errors.New("error 1")
	err = errors.Wrap(err, "error 2")
	fmt.Printf("%+v\n", err)

	err1 := xerrors.Errorf("error 3 : ")
	err2 := xerrors.Errorf("error 4 : %w\n", err1)
	err3 := xerrors.Errorf("error 5 : %w\n", err2)
	fmt.Printf("%+v\n", err3)
}
