package main

import (
	"fmt"
	"net"

	"github.com/Velousss/sampleproject/pkg/handler"
	"github.com/Velousss/sampleproject/pkg/types"
)

func main() {
	listener, err := net.Listen("tcp","localhost:9999")
	handler.HandleError(err)
	defer listener.Close()
	for{
		conn,err:=listener.Accept()
		handler.HandleError(err)
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn){
	defer conn.Close()
	for{
		payload,err:=types.Decode(conn)
		handler.HandleError(err)
		fmt.Println("Recieved: "+string(payload.Bytes()))

		data:=types.Binary("Server has recieved: "+string(payload.Bytes()))
		_,err=data.WriteTo(conn)
		handler.HandleError(err)
	}
}