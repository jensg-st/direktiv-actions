package main

import (
	"fmt"
	"strings"

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

	for i := range in {
		getValue(&in[i].value, in[i].name)
	}

	fmt.Printf("using server: %v\n", in[serverIdx].value)

	if in[serverIdx].value == "" || in[workflowIdx].value == "" {
		githubactions.Fatalf("server and workflow values are required\n")
	}

	doRequest(in)
}

func doRequest(in []args) {

	wf := strings.SplitN(in[workflowIdx].value, "/", 2)
	if len(wf) != 2 {
		githubactions.Fatalf("namespace/workflow is wroing format: %v\n",
			in[workflowIdx].value)
	}

	githubactions.Infof("executing workflow %s in %s\n", wf[0], wf[1])

	// set token if provided
	if len(in[tokenIdx].value) > 0 {
		githubactions.Infof("using token authentication\n")
	}

	// /api/namespaces/{namespace}/workflows/{workflow}/execute

	// resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

}

func getValue(val *string, key string) {
	*val = githubactions.GetInput(key)
}
