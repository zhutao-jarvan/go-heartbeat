package main

import (
	"flag"
	"fmt"
	"heartbeat"
)

var logCfgFile *string = flag.String("l", "D:\\src\\go\\heartbeat\\src\\cmd\\client\\log_config.xml", "Server log config file, default is log_config.xml")

func main() {
	c := heartbeat.NewClient("127.0.0.1", 18525, *logCfgFile)
	defer c.UdpConn.Close()

	c.Id = 9527
	c.Magic = 945201314

	c.HeartbeatSync()
	err, resp := c.HeartbeatResp()
	if err != nil {
		fmt.Printf("Heartbeat recv fail: %s\n", err.Error())
		return
	}

	fmt.Println("Heartbeat response: id[", resp.Id, "], magic[", resp.Magic, "], backlog[", resp.Backlog, "]")
}
