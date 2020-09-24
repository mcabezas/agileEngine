package main

import "flag"

type Config struct {
	Port         int
	TimeLocation string
}

func getConfigs() Config {
	port := flag.Int("port", 8080, "Set the http server port")
	location := flag.String("location", "America/Argentina/Buenos_Aires", "Set the server location which will be used to set the timezone")
	flag.Parse()
	return Config{
		Port:         *port,
		TimeLocation: *location,
	}
}
