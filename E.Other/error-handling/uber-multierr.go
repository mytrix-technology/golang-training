package error_handling

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"io"
)

func ExampleCombine() {
	err := multierr.Combine(
		errors.New("call 1 failed"),
		nil, // successful request
		errors.New("call 3 failed"),
		nil, // successful request
		errors.New("call 5 failed"),
	)
	fmt.Printf("%+v", err)
	// Output:
	// the following errors occurred:
	//  -  call 1 failed
	//  -  call 3 failed
	//  -  call 5 failed
}

func ExampleAppend() {
	var err error
	err = multierr.Append(err, errors.New("call 1 failed"))
	err = multierr.Append(err, errors.New("call 2 failed"))
	fmt.Println(err)
	// Output:
	// call 1 failed; call 2 failed
}

func ExampleErrors() {
	err := multierr.Combine(
		nil, // successful request
		errors.New("call 2 failed"),
		errors.New("call 3 failed"),
	)
	err = multierr.Append(err, nil) // successful request
	err = multierr.Append(err, errors.New("call 5 failed"))

	errors := multierr.Errors(err)
	for _, err := range errors {
		fmt.Println(err)
	}
	// Output:
	// call 2 failed
	// call 3 failed
	// call 5 failed
}

func ExampleAppendInto() {
	var err error

	if multierr.AppendInto(&err, errors.New("foo")) {
		fmt.Println("call 1 failed")
	}

	if multierr.AppendInto(&err, nil) {
		fmt.Println("call 2 failed")
	}

	if multierr.AppendInto(&err, errors.New("baz")) {
		fmt.Println("call 3 failed")
	}

	fmt.Println(err)
	// Output:
	// call 1 failed
	// call 3 failed
	// foo; baz
}

type fakeCloser func() error

func (f fakeCloser) Close() error {
	return f()
}

func FakeCloser(err error) io.Closer {
	return fakeCloser(func() error {
		return err
	})
}

func ExampleClose() {
	var err error

	closer := FakeCloser(errors.New("foo"))

	defer func() {
		fmt.Println(err)
	}()
	defer multierr.AppendInvoke(&err, multierr.Close(closer))

	fmt.Println("Hello, World")

	// Output:
	// Hello, World
	// foo
}
