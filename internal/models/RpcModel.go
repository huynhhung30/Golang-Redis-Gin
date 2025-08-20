package models

import (
	"Golang-Redis-Gin/internal/utils/functions"
	"net"
	"net/rpc"
	"os"
)

type Listener int

func StartRpcServer() {
	rpcPort := os.Getenv("RPC_PORT")
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+rpcPort)
	if err != nil {
		functions.ShowLog("StartRpcServerError1", err)
	}
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		functions.ShowLog("StartRpcServerError2", err)
	}
	listener := new(Listener)
	err = rpc.Register(listener)
	if err != nil {
		functions.ShowLog("StartRpcServerError3", err)
	}
	go rpc.Accept(inbound)
}

