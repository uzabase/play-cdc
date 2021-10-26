package main

import (
	"context"
	"fmt"
	"net"
	"google.golang.org/grpc"

	gm "play-cdc/gauge_messages"
)

type Event interface {}

type handler struct {
	*gm.UnimplementedReporterServer
	server *grpc.Server
}

func (h *handler) NotifyExecutionStarting(c context.Context, m *gm.ExecutionStartingRequest) (*gm.Empty, error) {
	fmt.Println("Received ExecutionStartingRequest")
	return &gm.Empty{}, nil
}

func NewHandler(s *grpc.Server) *handler {
	return &handler{UnimplementedReporterServer: new(gm.UnimplementedReporterServer), server: s}
}

func main() {
	e :=make(chan Event)
	startAPI(e)
}

func startAPI(e chan Event) {
	fmt.Println("Starting API")

	address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		panic("failed to resolve tcp addr.")
	}

	l, err := net.ListenTCP("tcp", address)
	if err != nil {
		panic("failed to listen tcp.")
	}

	server := grpc.NewServer(grpc.MaxRecvMsgSize(1024 * 1024 * 1024))

	h := NewHandler(server)

	fmt.Println("Registering Server")
	gm.RegisterReporterServer(server, h)

	fmt.Printf("Listening on port:%d\n", l.Addr().(*net.TCPAddr).Port)
	server.Serve(l)
}
