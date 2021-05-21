package main

import (
	"fmt"
	"net/url"
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
	protocolIdx = iota
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
		args{
			name: "protocol",
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

	githubactions.Infof("render url\n")

	u := &url.URL{}
	u.Scheme = in[protocolIdx].value
	u.Host = in[serverIdx].value
	u.Path = fmt.Sprintf("/api/namespaces/%s/workflows/%s/execute", wf[0], wf[1])

	githubactions.Infof("direktiv url %v\n", u.String())

	// r, err := http.Post(u.String(), "application/json", bytes.NewBuffer(""))

}

func getValue(val *string, key string) {
	*val = githubactions.GetInput(key)
}
