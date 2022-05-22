package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

var ghCredentials = flag.String("ghcreds", "", "github basic auth credentials")

func reqIssues(owner, repo, ghCredentials string) error {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	getIssuesURL := url.URL{
		Scheme: "https",
		Host:   "api.github.com",
		Path:   fmt.Sprintf("/repos/%s/%s/issues", owner, repo),
	}
	req := http.Request{
		Method: http.MethodGet,
		URL:    &getIssuesURL,
		Header: map[string][]string{
			"Accept": {"application/vnd.github.v3+json"},
		},
	}

	resp, err := client.Do(&req)
	if err != nil {
		log.Fatalf("Failed to request: %v", err)
	}

	respString, err := io.ReadAll(resp.Body)
	log.Printf("%s", respString)
	return nil
}

func main() {
	flag.Parse()
	_ = reqIssues("hashicorp", "vault", *ghCredentials)
}
