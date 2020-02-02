package heartbeat

import (
	"fmt"
	log "github.com/cihub/seelog"
	"net"
	"strings"
)

type Server struct {
	udpConn *net.UDPConn
}

func (s *Server) heartbeatServer() {
	for {
		buf := make([]byte, 1024)
		len, udpAddr, err := s.udpConn.ReadFromUDP(buf)
		ChkErrOnExit(err, fmt.Sprintf("Server ReadFromUDP failed!"))

		msg := strings.Replace(string(buf), "\n", "", 1)
		log.Debug("Read len[%d], msg: %s\n", len, msg)

		_, err = s.udpConn.WriteToUDP([]byte("ok\n"), udpAddr)
		ChkErrOnExit(err, fmt.Sprintf("Server WriteToUDP failed! IP: %s", udpAddr.IP.String()))
	}
}

func (s *Server) NewUdpServer(cfg *ServerConfig) {
	udpAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", cfg.Domain, cfg.Port))
	ChkErrOnExit(err, fmt.Sprintf("ResolveUDPAddr failed!"))

	udpConn, err := net.ListenUDP("udp", udpAddr)
	ChkErrOnExit(err, fmt.Sprintf("ListenUDP failed!"))
	s.udpConn = udpConn
}

// 开启服务
// @servCfgFile: 服务配置文件
// @logCfgFile:  日志配置文件
func NewAndRunServer(servCfgFile string, logCfgFile string) {
	defer LogFlush()
	s := &Server{}
	log.Info("Server init start...")
	InitLogAsFile(logCfgFile)

	cfg := GetServerConfig(servCfgFile)

	s.NewUdpServer(cfg)
	defer s.udpConn.Close()
	go s.heartbeatServer()

	log.Info("Server initialized")
	RpcRun(cfg)
	log.Warn("Server Exit...")
}
