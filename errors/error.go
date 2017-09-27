package errors

import (
	"fmt"
	"os"
)

func Fatal(err error, msg string, args ...interface{}) {
	Print(err, msg, args...)
	os.Exit(1)
}

func Print(err error, msg string, args ...interface{}) {
	if msg == "" {
		msg = "Error"
	}

	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	fmt.Fprintf(os.Stderr, "%s: %v", msg, err)
}
