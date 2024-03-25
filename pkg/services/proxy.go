package main

import (
	"io"
	"net"

	"github.com/Velousss/sampleproject/pkg/handler"
)

func main() {
	listener, err := net.Listen("tcp","localhost:1234")
	handler.HandleError(err)
	defer listener.Close()

	for{
		conn,err:=listener.Accept()
		handler.HandleError(err)
		go handleServer(conn)
	}
}

func forwardProxy(from io.Reader,to io.Writer) error{
	fromW,fromIsW:=from.(io.Writer)
	toR,toIsR:=to.(io.Reader)

	if fromIsW&&toIsR{
		go func(){
			_,err:=io.Copy(fromW,toR)
			handler.HandleError(err)
			return
		}()
	}
	_,err:=io.Copy(to,from)
	handler.HandleError(err)
	return err
}

func handleServer(to net.Conn){
	defer to.Close()

	from,err:=net.Dial("tcp","localhost:9999")
	handler.HandleError(err)
	defer to.Close()

	err=forwardProxy(from,to)
	handler.HandleError(err)
}