package main

import (
	"fmt"
	"os"

	"github.com/wojtekw92/go-playground/pkg/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}