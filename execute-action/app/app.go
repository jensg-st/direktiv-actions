package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("ARGS %v\n", os.Args)
	fmt.Printf("ENVS %v\n", os.Environ())
}
