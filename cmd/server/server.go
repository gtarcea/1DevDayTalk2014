package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gtarcea/1DevDayTalk2014/ws"
	"github.com/jessevdk/go-flags"
)

type serverOptions struct {
	Port uint   `long:"port" description:"The port the server listens on" default:"8081"`
	Bind string `long:"bind" description:"Address of local interface to listen on" default:"localhost"`
}

func main() {
	var opts serverOptions
	_, err := flags.Parse(&opts)

	if err != nil {
		os.Exit(1)
	}

	webserver(opts.Bind, opts.Port)
}

func webserver(bindTo string, port uint) {
	container := ws.NewRegisteredServicesContainer()
	http.Handle("/api/", http.StripPrefix("/api", container))
	webdir := os.Getenv("DEVDAY_WEBDIR")
	dir := http.Dir(webdir)
	http.Handle("/", http.FileServer(dir))
	addr := fmt.Sprintf("%s:%d", bindTo, port)
	fmt.Println(http.ListenAndServe(addr, nil))
}
