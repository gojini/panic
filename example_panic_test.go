package panic_test

import (
	"errors"
	"fmt"

	"gojini.dev/panic"
)

func ExampleOnError() {
	err := errors.New("some error")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic:", r)
		}
	}()

	panic.OnError(err)

	// Output:
	// Panic: some error
}
