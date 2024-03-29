package golib

import (
	"os"
	"fmt"
)

func CheckErr(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}