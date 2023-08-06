package config

import (
	"SocketSimpleServer/internal/log"
	"github.com/BurntSushi/toml"
	"net"
)

var Config config

type config struct {
	Tcp socketConf
	Udp socketConf
}

type socketConf struct {
	Ip        string
	Port      int
	MaxSize   int
	SimpleAck bool
	WholeAck  bool
}

//type tcp struct {
//	Ip        string
//	Port      int
//	MaxSize   int
//	SimpleAck bool
//	WholeAck  bool
//}
//
//type udp struct {
//	Ip        string
//	Port      int
//	MaxSize   int
//	SimpleAck bool
//	WholeAck  bool
//}

func Parse() {
	var (
		ip net.IP
	)
	if _, err := toml.DecodeFile("config.toml", &Config); err != nil {
		log.Log.Panic().Msgf("Decode config.toml Error: %v", err)
	}

	ip = net.ParseIP(Config.Tcp.Ip)
	if ip == nil {
		log.Log.Panic().Msg("TCP ip config parse Error")
	}

	ip = net.ParseIP(Config.Udp.Ip)
	if ip == nil {
		log.Log.Panic().Msg("UDP ip config parse Error")
	}

	if Config.Tcp.MaxSize <= 0 {
		log.Log.Panic().Msg("TCP maxSize <= 0")
	}

	if Config.Udp.MaxSize <= 0 {
		log.Log.Panic().Msg("UDP maxSize <= 0")
	}
}
