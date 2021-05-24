package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
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
	waitIdx     = iota
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
		args{
			name: "wait",
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

	u := &url.URL{}
	u.Scheme = in[protocolIdx].value
	u.Host = in[serverIdx].value
	u.Path = fmt.Sprintf("/api/namespaces/%s/workflows/%s/execute", wf[0], wf[1])

	fmt.Printf("ENS %v\n", os.Environ())
	if in[waitIdx].value == "true" {
		q := u.Query()
		q.Set("wait", "true")
		u.RawQuery = q.Encode()
	}

	githubactions.Infof("direktiv url %v\n", u.String())

	req, err := http.NewRequest("POST", u.String(),
		strings.NewReader(in[dataIdx].value))
	if err != nil {
		githubactions.Fatalf("")
	}
	req.Header.Set("Content-Type", "application/json")

	// r, err := http.Post(u.String(), "application/json", strings.NewReader(in[dataIdx].value))

}

func getValue(val *string, key string) {
	*val = githubactions.GetInput(key)
}
