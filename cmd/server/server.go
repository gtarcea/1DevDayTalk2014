package server

import (
	"os"

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

}
