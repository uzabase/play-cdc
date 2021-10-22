package main

import (
	"fmt"
	"net"
	"google.golang.org/grpc"

	"github.com/uzabase/play-cdc/gauge-messages"
)

type Event interface {}

type handler struct {
	server *grpc.Server
	e chan Event
}

func NewHandler(s *grpc.Server, e chan Event) *handler {
	return &handler{server: s, e: e}
}

func main() {
	e :=make(chan Event)
	go startAPI(e)
}

func startAPI(e chan Event) {
	address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		panic("failed to resolve tcp addr.")
	}

	l, err := net.ListenTCP("tcp", address)
	if err != nil {
		panic("failed to listen tcp.")
	}

	server := grpc.NewServer(grpc.MaxRecvMsgSize(1024 * 1024 * 1024))

	h := NewHandler(server, e)
	gauge_messages.RegisterReporterServer(server, h)

	fmt.Printf("Listening on port:%d\n", l.Addr().(*net.TCPAddr).Port)
	server.Serve(l)
}
