package main

import (
	"heartbeat"
	"log"
	"net/rpc"
)

type Rc struct {
	client *rpc.Client
}

func(rc *Rc) HeartBeat() string {
	result := new(string)
	err := rc.client.Call("Rs.HeartBeat", "", &result)
	heartbeat.ChkErrOnExit(err, "Call Rs.HeartBeat")
	return *result
}

func main() {
	rc := &Rc{}
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	heartbeat.ChkErrOnExit(err, "rpc.DialHTTP failed!")
	rc.client = client
	log.Printf("HeartBeat: %s", rc.HeartBeat())
}
