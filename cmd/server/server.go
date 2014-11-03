package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gtarcea/1DevDayTalk2014/ws"
	"github.com/jessevdk/go-flags"
)

type serverOptions struct {
	Port uint   `long:"server-port" description:"The port the server listens on" default:"8080"`
	Bind string `long:"bind" description:"Address of local interface to listen on" default:"localhost"`
}

func main() {
	var opts serverOptions
	_, err := flags.Parse(&opts)

	if err != nil {
		os.Exit(1)
	}

	webserver(opts.Port)
}

func webserver(port uint) {
	container := ws.NewRegisteredServicesContainer()
	http.Handle("/", container)
	webdir := os.Getenv("DEVDAY_WEBDIR")
	dir := http.Dir(webdir)
	http.Handle("/app/", http.StripPrefix("/app/", http.FileServer(dir)))
	addr := "localhost:8081"
	fmt.Println(http.ListenAndServe(addr, nil))
}
