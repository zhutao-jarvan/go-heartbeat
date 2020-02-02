package heartbeat

import (
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"io/ioutil"
)

type ServerConfig struct {
	Domain   string `json:domain`
	Port     uint16 `json:"port"`
	RpcPort  uint16 `json:"rpc_port"`
	DbUser   string `json:"db_user"`
	DbPwd    string `json:"db_password"`
	DbDomain string `json:"db_domain"`
	DbPort   string `json:"db_port"`
	DbName   string `json:"db_name"`
}

func GetServerConfig(configFile string) (cfg *ServerConfig) {
	log.Infof("Parse config file %s ...", configFile)
	cfg = new(ServerConfig)

	data, err := ioutil.ReadFile(configFile)
	ChkErrOnExit(err, fmt.Sprintf("Parse config file %s fail", configFile))

	log.Debug(string(data))
	err = json.Unmarshal(data, cfg)
	ChkErrOnExit(err, fmt.Sprintf("Unmarshal config file %s fail", configFile))

	return cfg
}
