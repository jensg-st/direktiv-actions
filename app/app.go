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

	fmt.Printf(">> %v\n", in)

	for i := range in {
		getValue(&in[i].value, in[i].name)
	}

	fmt.Printf(">> %v\n", in)

	fmt.Printf("using server: %v %v\n", in[serverIdx].value, in[serverIdx].name)
	fmt.Printf("executing workflow: %v\n", in[workflowIdx].value)

	fmt.Printf(">> %v\n", in)

	// fmt.Printf("ARGS %v\n", os.Args)
	// fmt.Printf("ENVS %v\n", os.Environ())
}

func getValue(val *string, key string) {
	githubactions.Infof("getting key for %v\n", key)
	*val = githubactions.GetInput(key)
	githubactions.Infof("value for %v\n", *val)
}
