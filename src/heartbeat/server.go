package heartbeat

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"io/ioutil"
	"net"
)

type ClientNode struct {
	Id       uint64
	Magic    uint64 // IP & PORT | 32BIT index
	LastHbTs uint32 // 最后心跳时间，单位秒
	Backlog  uint32 // 待处理事务数量
}

type Server struct {
	udpConn     *net.UDPConn
	cns         *ClientNode
	OnlineCount uint64
	TotalCount  uint64
}

type ServerConfig struct {
	Domain    string `json:domain`
	Port      uint16 `json:"port"`
	RpcPort   uint16 `json:"rpc_port"`
	CacheSize uint32 `json: cache_size`
	DbUser    string `json:"db_user"`
	DbPwd     string `json:"db_password"`
	DbDomain  string `json:"db_domain"`
	DbPort    string `json:"db_port"`
	DbName    string `json:"db_name"`
}

func parseServerConfig(configFile string) (cfg *ServerConfig) {
	log.Infof("Parse config file %s ...", configFile)
	cfg = new(ServerConfig)

	data, err := ioutil.ReadFile(configFile)
	ChkErrOnExit(err, fmt.Sprintf("Parse config file %s fail", configFile))

	log.Debug(string(data))
	err = json.Unmarshal(data, cfg)
	ChkErrOnExit(err, fmt.Sprintf("Unmarshal config file %s fail", configFile))

	return cfg
}

func (s *Server) heartbeatServer() {
	buf := make([]byte, CMsgLen)
	for {
		len, udpAddr, err := s.udpConn.ReadFromUDP(buf)
		ChkErrOnExit(err, fmt.Sprintf("Server ReadFromUDP failed!"))

		if len != CMsgLen {
			log.Errorf("Read len[%d] != CMsgLen[%d]", len, CMsgLen)
			continue
		}

		buf[CMsgTypeOffset] = CMsgTypeRsp
		binary.BigEndian.PutUint32(buf[CMsgBacklogOffset:], 1)

		_, err = s.udpConn.WriteToUDP(buf, udpAddr)
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

	cfg := parseServerConfig(servCfgFile)

	s.NewUdpServer(cfg)
	defer s.udpConn.Close()
	go s.heartbeatServer()

	log.Info("Server initialized")
	RpcRun(cfg)
	log.Warn("Server Exit...")
}
