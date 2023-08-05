package main

import (
	"SocketSimpleServer/internal/config"
	"SocketSimpleServer/internal/log"
	"SocketSimpleServer/service/tcp"
	"SocketSimpleServer/service/udp"
	"time"
)

func main() {
	log.Log.Info().Msg("Welcome to SocketSimpleServer.")
	log.Log.Info().Msg("Copyright 2023 dreamcmi. All rights reserved.")
	config.Parse()
	go tcp.Run()
	go udp.Run()
	for {
		time.Sleep(1 * time.Second)
	}
}
