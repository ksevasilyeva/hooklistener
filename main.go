package main

import (
	"net/http"
	"errors"
	"log"
	"../hooklistener/github"
	"../hooklistener/bitbucket"
	"../hooklistener/util/provider"
)

func main() {
	RunDefautServer()
}

func RunDefautServer() {
	server := http.Server{
		Addr:    ":8080",
		Handler: &hookHandler{},
	}
	server.ListenAndServe()
	log.Println("INFO", "Start listening :8080")
}

func (*hookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p, err := getHookProvider(r)
	if err != nil {
		log.Println("ERROR", err)
	}
	if p != 0 {
		log.Println("INFO", p)
	}
}

type hookHandler struct{}

type Webhooker interface {
	Provider() provider.Provider
}

func getHookProvider(r *http.Request) (provider.Provider, error) {
	if r.Header.Get(github.GithubEventHeader) != "" {
		return github.ParsePayload(r).Provider(), nil
	}
	if r.Header.Get(bitbucket.BitbucketEventHeader) != "" {
		return bitbucket.ParsePayload(r).Provider(), nil
	}
	return provider.Unknown, errors.New("Unknown provider")
}