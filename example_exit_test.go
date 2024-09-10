package panic_test

import (
	"errors"
	"fmt"

	"gojini.dev/panic"
)

func ExampleExitOnError() {
	err := errors.New("some error")

	panic.ExitFunc = func(code int) {
		fmt.Println("Exit:", code)
	}

	panic.ExitOnError(err)

	// Output:
	// Exit: 1
}
