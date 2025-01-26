package main

import (
	"log"
	"net/rpc"

	"github.com/syniol/go-rpc/pkg"
)

func main() {
	rpcClient, err := rpc.Dial("tcp", "syniol_rpc_server:8080")
	if err != nil {
		log.Println("rpcClient", err)
	}

	var resp pkg.Response
	err = rpcClient.Call(
		"Handler.Execute",
		&pkg.Request{Name: "Syniol Limited"},
		&resp,
	)
	if err != nil {
		log.Println("Call", err)
	}

	log.Println(resp.Message)
}
