package util

import (
	"fmt"
	"os"
)

func HandleError(err error, msg string, args ...interface{}) {
	if msg == "" {
		msg = "Error"
	}

	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	fmt.Fprintf(os.Stderr, "%s: %v", msg, err)
	os.Exit(1)
}
