package main

import (
	"fmt"
	"os"
	"github.com/warmbly/warmbly/internal/pkg/argon2"
)

func main() {
	h, err := argon2.Hash(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Print(h)
}
