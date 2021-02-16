package main

import (
	"github.com/bborbe/http/requestbuilder"
	"github.com/bborbe/http/client_builder"
	"github.com/golang/glog"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	hostPtr   = flag.String("host", "", "Jira URL")
	userPtr   = flag.String("user", "", "Jira User")
	passPtr   = flag.String("pass", "", "Jira Password")
	schemaPtr = flag.String("schema", "", "Insight Objectschema Key")
)

func main() {
	flag.Parse()

	var rb requestbuilder.HttpRequestBuilder = requestbuilder.NewHTTPRequestBuilder(fmt.Sprintf("https://%s/rest/insight/1.0/objectschema/list", *hostPtr))
	rb = rb.AddBasicAuth(*userPtr, *passPtr)
	req, err := rb.Build()
	if err != nil {
		glog.Exitf("build request failed: %v", err)
	}
	httpClientBuilder := client_builder.New()
	httpClient := httpClientBuilder.Build()
	resp, err := httpClient.Do(req)
	if err != nil {
		glog.Exitf("do request failed: %v", err)
	}
	if resp.StatusCode/100 != 2 {
		glog.Exitf("statusCode %d != 2xx", resp.StatusCode)
	}
	io.Copy(os.Stdout, resp.Body)
}
