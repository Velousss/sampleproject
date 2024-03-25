package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/Velousss/sampleproject/pkg/handler"
	"github.com/Velousss/sampleproject/pkg/types"
)

func main() {
	dialer, err := net.Dial("tcp","localhost:1234")
	handler.HandleError(err)
	defer dialer.Close()

	scanner:=bufio.NewScanner(os.Stdin)
	var msg string

	for{
		fmt.Print("Message: ")
		scanner.Scan()
		msg = scanner.Text()

		data := types.Binary(msg)
		_,err = data.WriteTo(dialer)
		handler.HandleError(err)

		payload,err := types.Decode(dialer)
		handler.HandleError(err)

		err=dialer.SetReadDeadline(time.Now().Add(5*time.Second))
		if err!=nil{
			if netErr,ok:=err.(net.Error);ok&&netErr.Timeout(){
				fmt.Println(err)
			}else{
				handler.HandleError(err)
			}
		}
		fmt.Println(string(payload.Bytes()))
	}
}