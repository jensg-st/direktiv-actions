package main

import (
	"fmt"

	"github.com/sethvargo/go-githubactions"
)

func main() {

	var server string
	getValue(&server, "server")

	fmt.Printf("using server: %v", server)

	// fmt.Printf("ARGS %v\n", os.Args)
	// fmt.Printf("ENVS %v\n", os.Environ())
}

func getValue(val *string, key string) {

	*val = githubactions.GetInput(key)
	if *val == "" {
		githubactions.Fatalf("missing '%s'", key)
	}

}
