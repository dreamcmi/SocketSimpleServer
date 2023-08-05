package config

import (
	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog/log"
	"net"
)

var Config config

type config struct {
	Tcp tcp
	Udp udp
}

type tcp struct {
	Ip      string
	Port    int
	MaxSize int
}

type udp struct {
	Ip      string
	Port    int
	MaxSize int
}

func Parse() {
	var (
		ip net.IP
	)
	if _, err := toml.DecodeFile("config.toml", &Config); err != nil {
		log.Panic().Msgf("Decode config.toml Error: %v", err)
	}

	ip = net.ParseIP(Config.Tcp.Ip)
	if ip == nil {
		log.Panic().Msg("TCP ip config parse Error")
	}

	ip = net.ParseIP(Config.Udp.Ip)
	if ip == nil {
		log.Panic().Msg("UDP ip config parse Error")
	}

	if Config.Tcp.MaxSize <= 0 {
		log.Panic().Msg("TCP maxSize <= 0")
	}

	if Config.Udp.MaxSize <= 0 {
		log.Panic().Msg("UDP maxSize <= 0")
	}
}
