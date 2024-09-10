package panic

import (
	"fmt"
	"os"
)

var ExitFunc = os.Exit

func OnError(err error) {
	if err != nil {
		panic(err)
	}
}

func ExitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		ExitFunc(1)
	}
}
