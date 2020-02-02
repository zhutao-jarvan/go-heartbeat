package main

import (
	"flag"
	"heartbeat"
)

var cfgFile *string = flag.String("c", "D:\\src\\go\\heartbeat\\src\\cmd\\server\\config.json", "Server config file, default is config.json")
var logCfgFile *string = flag.String("l", "D:\\src\\go\\heartbeat\\src\\cmd\\server\\log_config.xml", "Server log config file, default is log_config.xml")

func main() {
	heartbeat.NewAndRunServer(*cfgFile, *logCfgFile)
}
