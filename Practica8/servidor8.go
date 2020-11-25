package main

import(
	"fmt"
	"net"
	"net/rpc"
)

type Server struct{}

func (this *Server) AddScore(paso []string, reply *[]string)error{
	paso = append(paso, "aaa")
	paso = append(paso, "popó")
	*reply = paso
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}