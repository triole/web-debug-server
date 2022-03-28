package main

import (
	"fmt"
	"net/http"

	"web-debug-server/src/logging"

	"github.com/sirupsen/logrus"
)

var (
	lg logging.Logging
)

type tConf struct {
	Bind          string
	Port          int
	ResponseDelay []int
}

func main() {
	parseArgs()

	conf := tConf{
		Port:          CLI.Port,
		ResponseDelay: CLI.ResponseDelay,
	}

	lg = logging.Init(CLI.LogFile, CLI.JSONLog)

	bind := ""
	if bind == "" {
		bind = fmt.Sprintf(":%d", conf.Port)
	}

	httpServer := &http.Server{
		Addr:    bind,
		Handler: &handler{},
	}

	if lg.LogToFile == true {
		fmt.Printf(
			"Run server\tport=%d, verbose=%v, logfile=%s\n",
			conf.Port, CLI.Verbose, CLI.LogFile,
		)
	}

	lg.LogInfo("Run server", logrus.Fields{
		"port":    conf.Port,
		"verbose": CLI.Verbose,
		"logfile": CLI.LogFile,
	})

	lg.LogError(httpServer.ListenAndServe(), nil)
}
