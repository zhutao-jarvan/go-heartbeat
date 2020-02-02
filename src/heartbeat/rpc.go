package heartbeat

import (
	"fmt"
	log "github.com/cihub/seelog"
	"net"
	"net/http"
	"net/rpc"
)

type Rs struct {
}

// 查看程序是否还在正常运行
func (r *Rs) HeartBeat(null string, reply *string) error {
	*reply = "I'm Ok"
	return nil
}

func RpcRun(cfg *ServerConfig)  {
	r := new(Rs)
	ChkErrOnExit(rpc.Register(r), "Register Rpc fail,")
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", cfg.RpcPort))
	ChkErrOnExit(err, fmt.Sprintf("Listen Rpc port %d fail,", cfg.RpcPort))

	log.Info("Rpc initialized")
	http.Serve(l, nil)
}
