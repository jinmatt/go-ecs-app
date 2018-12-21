package main

import (
	"fmt"
	"net/http"

	logger "github.com/sirupsen/logrus"
	goji "goji.io"
	"goji.io/pat"
)

func hello(w http.ResponseWriter, r *http.Request) {
	logger.WithField("handler", "hello").Info("hello handler called")
	fmt.Fprint(w, "Hello, ECS from master with wip!")
}

func helloTo(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	logger.WithField("handler", "hello").Infof("helloTo handler called with name %", name)
	fmt.Fprintf(w, "Hello, %s!", name)
}

func health(w http.ResponseWriter, r *http.Request) {
	logger.WithField("handler", "health").Info("health handler called")
	fmt.Fprint(w, "healthy")
}

func main() {
	logger.WithField("action", "setup").Info("Setting up ecs app")
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/"), hello)
	mux.HandleFunc(pat.Get("/hello/:name"), helloTo)
	mux.HandleFunc(pat.Get("/health"), health)

	logger.WithField("action", "startup").Infof("Listening on port: %s", "8000")
	http.ListenAndServe(":8000", mux)
}
