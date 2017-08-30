package bitbucket

import (
	"net/http"
	"../util/provider"
)

const BitbucketEventHeader = "X-Event-Key"

type Webhook struct {
	Event string
}

func ParsePayload(r *http.Request) *Webhook {
	return &Webhook{
		Event: r.Header.Get(BitbucketEventHeader)}
}

func (w Webhook) Provider() provider.Provider {
	return provider.Bitbucket
}
