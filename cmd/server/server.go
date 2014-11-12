package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"code.google.com/p/go.net/websocket"

	"github.com/gtarcea/1DevDayTalk2014/ws"
	"github.com/gtarcea/1DevDayTalk2014/ws/events"
	"github.com/jessevdk/go-flags"
)

// serverOptions describes the command line arguments that the server command accepts.
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

	server(opts.Bind, opts.Port)
}

// server sets up all the services. This includes:
//   a. Serving the website
//   b. The websocket service
//   c. The users rest service
func server(bindTo string, port uint) {
	privateKey := readKeyFile("app.rsa")
	publicKey := readKeyFile("app.rsa.pub")
	container := ws.NewRegisteredServicesContainer(privateKey, publicKey)
	http.Handle("/api/", http.StripPrefix("/api", container))

	webdir := os.Getenv("DEVDAY_WEBDIR")
	dir := http.Dir(webdir)
	http.Handle("/", http.FileServer(dir))

	hub := events.NewHub(events.NewEventConnections())
	hub.Start()
	s := events.NewServer(hub)
	http.Handle("/ws", websocket.Handler(s.OnConnection))

	addr := fmt.Sprintf("%s:%d", bindTo, port)
	fmt.Println(http.ListenAndServe(addr, nil))
}

// readKeyFile reads the public and private key files. It assumes
// the files are stored in the directory pointed at by the
// DEVDAY_CONFIG environment variable. It panics if it cannot read
// the given file.
func readKeyFile(filename string) []byte {
	configdir := os.Getenv("DEVDAY_CONFIG")
	if configdir == "" {
		panic("DEVDAY_CONFIG is not set")
	}
	path := filepath.Join(configdir, filename)
	key, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Unable to read file: %s", path))
	}

	return key
}
