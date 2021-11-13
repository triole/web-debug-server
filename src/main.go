package main

import (
	"fmt"
	"log"
	"net/http"
)

type tConf struct {
	Bind string
	Port int
}

func main() {
	parseArgs()

	conf := tConf{
		Port: CLI.Port,
	}

	bind := ""
	if bind == "" {
		bind = fmt.Sprintf(":%d", conf.Port)
	}

	httpServer := &http.Server{
		Addr:    bind,
		Handler: &handler{},
	}

	log.Printf("[INFO] Listening at %s", bind)
	log.Fatal(httpServer.ListenAndServe())
}
