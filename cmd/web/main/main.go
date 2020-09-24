package main

import (
	"strconv"

	"github.com/mcabezas/agileEngine/internal/logs"
)

func main() {
	conf := getConfigs()
	logger := logs.NewSugaredLogger()
	mux := Routes(logger)
	srv := newServer(strconv.Itoa(conf.Port), mux, logger)
	srv.Start()
}