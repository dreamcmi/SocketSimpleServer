package tcp

import (
	"SocketSimpleServer/internal/config"
	"SocketSimpleServer/internal/log"
	"encoding/hex"
	"net"
	"strconv"
)

func Run() {
	var address string
	log.Log.Info().Msgf("TCP Run: IP(%s) PORT(%d) MAXSIZE(%d)",
		config.Config.Tcp.Ip,
		config.Config.Tcp.Port,
		config.Config.Tcp.MaxSize)
	address = config.Config.Tcp.Ip + ":" + strconv.Itoa(config.Config.Tcp.Port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Log.Panic().Msgf("TCP Listen Error: %v", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Log.Error().Msgf("TCP Accept Error: %v", err)
			return
		}
		log.Log.Info().Msgf("TCP client connect(%s)", conn.RemoteAddr().String())
		go tcpProcess(conn) // 客户端接入进协程
	}
}

func tcpProcess(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Log.Error().Msgf("%v", err)
		}
	}(conn)

	for {
		buf := make([]byte, config.Config.Tcp.MaxSize)
		n, err := conn.Read(buf)
		if err != nil {
			log.Log.Error().Msgf("TCP(%s) Error: %v", conn.RemoteAddr().String(), err)
			break
		}
		dataString := string(buf[:n])
		dataHex := hex.EncodeToString(buf[:n])
		if config.Config.Tcp.SimpleAck {
			simpleAckBuf := "len:" + strconv.Itoa(len(dataString))
			_, err = conn.Write([]byte(simpleAckBuf))
			if err != nil {
				log.Log.Error().Msgf("TCP(%s) simple ack Write Error: %v", conn.RemoteAddr().String(), err)
			}
		}
		if config.Config.Tcp.WholeAck {
			_, err = conn.Write(buf)
			if err != nil {
				log.Log.Error().Msgf("TCP(%s) whole ack Write Error: %v", conn.RemoteAddr().String(), err)
			}
		}
		log.Log.Info().Msgf("TCP(%s) Receive(%d):%s", conn.RemoteAddr().String(), len(dataString), dataString)
		log.Log.Info().Msgf("TCP(%s) Receive Hex:%s", conn.RemoteAddr().String(), dataHex)
	}
	log.Log.Info().Msgf("TCP client disconnect(%s)", conn.RemoteAddr().String())
}
