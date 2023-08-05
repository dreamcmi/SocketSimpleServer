package udp

import (
	"SocketSimpleServer/internal/config"
	"SocketSimpleServer/internal/log"
	"encoding/hex"
	"net"
)

func Run() {
	log.Log.Info().Msgf("UDP Run: IP(%s) PORT(%d) MAXSIZE(%d)",
		config.Config.Udp.Ip,
		config.Config.Udp.Port,
		config.Config.Udp.MaxSize)

	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP(config.Config.Udp.Ip),
		Port: config.Config.Udp.Port,
	})
	if err != nil {
		log.Log.Panic().Msgf("UDP Listen Error: %v", err)
		return
	}
	for {
		data := make([]byte, config.Config.Udp.MaxSize)
		count, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			log.Log.Error().Msgf("UDP Read Failed: %v", err)
			continue
		}
		dataString := string(data[:count])
		dataHex := hex.EncodeToString(data[:count])
		log.Log.Info().Msgf("UDP(%s) Receive(%d):%s", addr, len(dataString), dataString)
		log.Log.Info().Msgf("UDP(%s) Receive Hex:%s", addr, dataHex)
	}
}
