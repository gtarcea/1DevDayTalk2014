package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"code.google.com/p/go.net/websocket"

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
	channelTest()
}

func webserver(bindTo string, port uint) {
	//http.HandleFunc("/ws", serveWebsocket)
	container := ws.NewRegisteredServicesContainer()
	http.Handle("/api/", http.StripPrefix("/api", container))
	webdir := os.Getenv("DEVDAY_WEBDIR")
	dir := http.Dir(webdir)
	http.Handle("/", http.FileServer(dir))
	http.Handle("/ws", websocket.Handler(EchoServer))
	addr := fmt.Sprintf("%s:%d", bindTo, port)
	fmt.Println(http.ListenAndServe(addr, nil))
}

func EchoServer(ws *websocket.Conn) {
	var msg string
	websocket.Message.Receive(ws, &msg)
	fmt.Println(msg)
}

type waiter struct {
	newValue  string
	observers chan bool
}

func channelTest() {
	w := &waiter{
		newValue:  "Hello",
		observers: make(chan bool),
	}

	for i := 0; i < 5; i++ {
		go startWaiter(w, i)
	}

	time.Sleep(5 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case w.observers <- true:
			fmt.Println("sent true")
		case <-time.After(2 * time.Second):
			fmt.Println("timeout")
		}
	}
	time.Sleep(2 * time.Second)
}

func startWaiter(w *waiter, i int) {
	fmt.Println("waiter", i, " waiting")
	for {
		select {
		case <-w.observers:
			fmt.Println("waiter", i, " got value", w.newValue)
		}
	}

}
