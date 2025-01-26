package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/syniol/go-rpc/pkg"
)

func main() {
	rpcServer := rpc.NewServer()
	_ = rpcServer.Register(new(pkg.Handler))

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("error creating a TCP server on port 8080")
	}

	log.Println("🚀RPC Server is running on port: 8080")
	rpcServer.Accept(l)
}
