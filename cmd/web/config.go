package main

import "flag"

type config struct {
	addr string
	staticDir string
}

func initConf() (config) {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir,"staticDir", "../../ui/static/", "Path to static folder")
	flag.Parse()
	return cfg
}