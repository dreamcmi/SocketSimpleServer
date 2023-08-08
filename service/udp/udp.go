package udp

import (
	"SocketSimpleServer/internal/config"
	"SocketSimpleServer/internal/log"
	"encoding/hex"
	"net"
	"strconv"
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
		if config.Config.Udp.SimpleAck {
			simpleAckBuf := "len:" + strconv.Itoa(len(dataString))
			_, err = listen.WriteToUDP([]byte(simpleAckBuf), addr)
			if err != nil {
				log.Log.Error().Msgf("UDP(%s) simple ack Write Error: %v", addr, err)
			}
		}
		if config.Config.Udp.WholeAck {
			_, err = listen.WriteToUDP(data, addr)
			if err != nil {
				log.Log.Error().Msgf("UDP(%s) whole ack Write Error: %v", addr, err)
			}
		}
		log.Log.Info().Msgf("UDP(%s) Receive(%d):%s", addr, len(dataString), dataString)
		log.Log.Info().Msgf("UDP(%s) Receive Hex:%s", addr, dataHex)
	}
}
