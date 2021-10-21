package main

import (
	"fmt"
	"net"
	"google.golang.org/grpc"
)

type Event interface {}

func handleEvents(e chan Event) {
	// TODO
}

type handler struct {
	server *grpc.Server
	e chan Event
}

func NewHandler(s *grpc.Server, e chan Event) *handler {
	return &handler{server: s, e: e}
}

type ReporterServer interface {
	// TODO
}

func RegisterReporterServer(s *grpc.Server, src ReporterServer) {
	// TODO
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
	RegisterReporterServer(server, h)

	fmt.Printf("Listening on port:%d\n", l.Addr().(*net.TCPAddr).Port)
	server.Serve(l)
}
