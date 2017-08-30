package github

import (
	"net/http"
	"../util/provider"
)

const GithubEventHeader = "X-GitHub-Event"

type Webhook struct {
	Event string
}

func ParsePayload(r *http.Request) *Webhook {
	return &Webhook{
		Event: r.Header.Get(GithubEventHeader)}
}

func (w Webhook) Provider() provider.Provider {
	return provider.GitHub
}
