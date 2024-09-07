package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct{
	logger *slog.Logger
}

func main() {
	
	//Define a new command-line flag with the name 'addr', a default value of ':4000'
	addr := flag.String("addr", ":4000", "HTTP network address")
	//Parse the command-line flag
	flag.Parse()

	//logger a structured logger to write the standard out stream and uses the default settings.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//Initialize a new instance of our application struct, containing the dependencies
	app := &application{
		logger: logger,
	}

	//Server
	logger.Info("starting server","addr", *addr)
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
