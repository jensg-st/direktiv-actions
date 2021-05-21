package main

import (
	"fmt"

	"github.com/sethvargo/go-githubactions"
)

type args struct {
	name  string
	value string
}

const (
	serverIdx   = iota
	workflowIdx = iota
	tokenIdx    = iota
	dataIdx     = iota
)

func main() {

	in := []args{
		args{
			name: "server",
		},
		args{
			name: "workflow",
		},
		args{
			name: "token",
		},
		args{
			name: "data",
		},
	}

	for _, i := range in {
		getValue(&i.value, i.name)
	}

	fmt.Printf("using server: %v", in[serverIdx].value)
	fmt.Printf("executing workflow: %v", in[workflowIdx].value)

	// fmt.Printf("ARGS %v\n", os.Args)
	// fmt.Printf("ENVS %v\n", os.Environ())
}

func getValue(val *string, key string) {
	*val = githubactions.GetInput(key)
}
