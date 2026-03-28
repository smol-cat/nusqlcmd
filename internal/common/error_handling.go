package common

import (
	"fmt"
	"os"
)

func PanicOnErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ExitOnErr(err error, code int) {
	ExitOnErrFunc(err, code, func(err error) {
		fmt.Println("Error: " + err.Error())
	})
}

func ExitOnErrFunc(err error, code int, action func(error)) {
	if err != nil {
		action(err)
		os.Exit(code)
	}
}

func WarnUnrecognizedType(typeName string) {
	fmt.Fprintf(os.Stderr, "Warning: Unrecognized type '%s', defaulting to string\n", typeName)
}
