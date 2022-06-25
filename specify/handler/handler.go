package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	gm "specify/gauge_messages"
)

func Start() {
	address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer(grpc.MaxRecvMsgSize(1024 * 1024 * 1024))

	handler := newHandler(server)

	gm.RegisterReporterServer(server, handler)

	fmt.Printf("Listening on port:%d\n", listener.Addr().(*net.TCPAddr).Port)
	server.Serve(listener)
}

type handler struct {
	*gm.UnimplementedReporterServer
	server *grpc.Server
}

func newHandler(s *grpc.Server) *handler {
	return &handler{UnimplementedReporterServer: new(gm.UnimplementedReporterServer)}
}

func (h *handler) NotifyExecutionStarting(c context.Context, m *gm.ExecutionStartingRequest) (*gm.Empty, error) {
	fmt.Println("Received ExecutionStartingRequest")
	return &gm.Empty{}, nil
}

func (h *handler) NotifyExecutionEnding(c context.Context, m *gm.ExecutionEndingRequest) (*gm.Empty, error) {
	fmt.Println("Received ExecutionEndingRequest")
	return &gm.Empty{}, nil
}

func (h *handler) NotifySpecExecutionStarting(c context.Context, m *gm.SpecExecutionStartingRequest) (*gm.Empty, error) {
	fmt.Println("Received SpecExecutionStartingRequest")
	return &gm.Empty{}, nil
}

func (h *handler) NotifySpecExecutionEnding(c context.Context, m *gm.SpecExecutionEndingRequest) (*gm.Empty, error) {
	fmt.Println("Received SpecExecutionEndingRequest")
	return &gm.Empty{}, nil
}

func (h *handler) NotifyScenarioExecutionStarting(c context.Context, m *gm.ScenarioExecutionStartingRequest) (*gm.Empty, error) {
	fmt.Println("Received ScenarioExecutionStartingRequest")
	return &gm.Empty{}, nil
}

func (h *handler) NotifyScenarioExecutionEnding(c context.Context, m *gm.ScenarioExecutionEndingRequest) (*gm.Empty, error) {
	fmt.Println("Received ScenarioExecutionEndingRequest")
	return &gm.Empty{}, nil
}

func (h *handler) NotifyStepExecutionStarting(c context.Context, m *gm.StepExecutionStartingRequest) (*gm.Empty, error) {
	fmt.Println("Received StepExecutionStartingRequest")
	return &gm.Empty{}, nil
}

func (h *handler) NotifyStepExecutionEnding(c context.Context, m *gm.StepExecutionEndingRequest) (*gm.Empty, error) {
	fmt.Println("Received StepExecutionEndingRequest")
	return &gm.Empty{}, nil
}

func (h *handler) NotifySuiteResult(c context.Context, m *gm.SuiteExecutionResult) (*gm.Empty, error) {
	fmt.Println("Received SuiteExecutionResult")
	return &gm.Empty{}, nil
}
