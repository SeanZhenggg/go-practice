package main

import (
	"errors"
	"fmt"
	"golang.org/x/xerrors"
)

type errResp struct {
	Code    int
	Message string
	Raw     string
}

func (e errResp) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Raw: %s", e.Code, e.Message, e.Raw)
}

func main() {
	//err := errors.New("error 1")
	//err = errors.Wrap(err, "error 2")
	//fmt.Printf("%+v\n", err)
	//
	//err1 := xerrors.Errorf("error 3 : %w", errors.New("i am a error!!!"))
	//err2 := xerrors.Errorf("error 4 : %w\n", err1)
	//err3 := xerrors.Errorf("error 5 : %w\n", err2)
	//fmt.Printf("%+v\n", err3)

	//var tmp error
	//if errors.As(err1, &tmp) {
	//	fmt.Printf("temp : %+v\n", tmp)
	//}

	//err := errors.New("i am error")
	//wraped := gErr.Wrap(err, "second error")
	//fmt.Printf("%v", wraped)

	var (
	//ErrBase = errors.New("a new error")
	//ErrBase2 = xerrors.New("another new error")
	)

	//err := xerrors.Errorf("raiseError:%w", ErrBase)
	//fmt.Printf("%+v\n", err)
	//
	//err2 := xerrors.Errorf("raiseError2: %w", nil)
	//fmt.Print("%+v\n", err2)
	resp := &errResp{Code: 100, Message: "i am error", Raw: `i am raw error response`}
	err1 := xerrors.Errorf("this is error 1 : %w", resp)
	err2 := xerrors.Errorf("this is error 2 : %w", err1)
	var e = &errResp{}
	if ok := errors.As(err2, &e); ok {
		fmt.Printf("is resp err: %v", e)
	} else {
		fmt.Printf("is not resp err: %v", e)
	}

	//err3 := errors.Wrap(ErrBase, "failed error")
	//fmt.Printf("%+v\n", err3)

	//err4 := errors.Wrap(nil, "failed error")
	//fmt.Print("%+v\n", err4)
	//fmt.Println(errors.Is(err, ErrBase2))

}
