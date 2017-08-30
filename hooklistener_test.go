package main

import (
	"testing"
	"net/http"
	"bytes"
	"../hooklistener/util/provider"
	"errors"
)

func TestHandler(t *testing.T) {

	cases := []struct {
		headerName string
		headerValue string
		provider provider.Provider
		err error
	}{
		{"X-Event-Key", "repo:push", provider.Bitbucket, nil},
		{"X-GitHub-Event", "CommitCommentEvent", provider.GitHub, nil},
		{"", "", provider.Unknown, errors.New("Unknown provider")},
	}

	for _, c := range cases {
		req, _ := http.NewRequest("POST", "http://127.0.0.1:8080", bytes.NewBuffer([]byte("{}")))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set(c.headerName, c.headerValue)

		p, err := getHookProvider(req)

		if p != c.provider {
			t.Errorf("Expected: %q, Found %q", c.provider.String(), p)
		}

		if c.err != nil {
			if c.err.Error() != err.Error() {
				t.Errorf("Expected: %q, Found %q", c.err.Error(), err.Error())
			}
		}
	}
}

func TestProviderStringGithub(t *testing.T) {
	if provider.GitHub.String() != "GitHub" {
		t.Errorf("Expected: %q, Found %q", "GitHub", provider.GitHub.String())
	}
}

func TestProviderStringBitbucket(t *testing.T) {
	if provider.Bitbucket.String() != "Bitbucket" {
		t.Errorf("Expected: %q, Found %q", "GitHub", provider.GitHub.String())
	}
}