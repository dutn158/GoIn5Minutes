package main

import (
	"net/rpc"
	"net"
	"fmt"
)

type Server struct {}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil;
}

func (this *Server) Reverse(s string, reply *string) error {
	b := []byte(s)
	b2 := make([]byte, len(b))
	for i := 0; i < len(b); i++ {
		b2[len(b2) - i - 1] = b[i]
	}
	*reply = string(b2)
	return nil
}

func server () {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
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

func client () {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64
	err = c.Call("Server.Negate", int64(1508), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(1508) = ", result)
	}

	var reply string
	err = c.Call("Server.Reverse", string("Tran Ngoc Du"), &reply)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Reverse(\"Tran Ngoc Du\") = ", reply)
	}
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
