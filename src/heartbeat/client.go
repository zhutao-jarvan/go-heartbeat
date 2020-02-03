package heartbeat

import (
	"encoding/binary"
	"fmt"
	log "github.com/cihub/seelog"
	"net"
)

type Client struct {
	UdpConn *net.UDPConn
	Id      uint64
	Magic   uint64
}

func (c *Client) HeartbeatResp() (err error, resp *HeartbeatMsg) {
	resp = &HeartbeatMsg{}
	buf := make([]byte, CMsgLen)

	var len int
	len, err = c.UdpConn.Read(buf)

	if len != CMsgLen {
		log.Errorf("Read len[%d] != CMsgLen[%d]", len, CMsgLen)
		return err, nil
	}

	resp.Id = binary.BigEndian.Uint64(buf[CMsgIdOffset:])
	resp.Magic = binary.BigEndian.Uint64(buf[CMsgMagicOffset:])
	resp.Type = buf[CMsgTypeOffset]
	resp.Backlog = binary.BigEndian.Uint32(buf[CMsgBacklogOffset:])

	return nil, resp
}

func (c *Client) HeartbeatSync() (err error) {
	req := make([]byte, CMsgLen)

	binary.BigEndian.PutUint64(req[CMsgIdOffset:], c.Id)
	binary.BigEndian.PutUint64(req[CMsgMagicOffset:], c.Magic)
	req[CMsgTypeOffset] = CMsgTypeSync

	_, err = c.UdpConn.Write(req)
	return err
}

func (c *Client) NewUdpClient(domain string, port uint16) {
	udpAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", domain, port))
	ChkErrOnExit(err, "ResolveUDPAddr fail!")

	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	ChkErrOnExit(err, "ListenUDP fail!")

	log.Debug("Udp dial ok!")
	c.UdpConn = udpConn
}

// 开启客户端
// @servCfgFile: 服务配置文件
// @logCfgFile:  日志配置文件
func NewClient(domain string, port uint16, logCfgFile string) (c *Client){
	defer LogFlush()
	c = &Client{}
	log.Debug("Client init start...")
	InitLogAsFile(logCfgFile)

	c.NewUdpClient(domain, port)

	return c
}
