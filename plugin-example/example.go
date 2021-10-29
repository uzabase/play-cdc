package main

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"net/http"
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

func (h *handler) NotifySpecExecutionStarting(c context.Context, m *gm.SpecExecutionStartingRequest) (*gm.Empty, error) {
    fmt.Println("Received SpecExecutionStartingRequest")
    return &gm.Empty{}, nil
}

func (h *handler) NotifyStepExecutionStarting(c context.Context, m *gm.StepExecutionStartingRequest) (*gm.Empty, error) {
    fmt.Println("Received StepExecutionStarting")
    fmt.Printf("Step actual text: %v\n", m.CurrentExecutionInfo.CurrentStep.Step.ActualStepText)
    fmt.Printf("Step parsed text: %v\n", m.CurrentExecutionInfo.CurrentStep.Step.ParsedStepText)
    for _, p := range m.CurrentExecutionInfo.CurrentStep.Step.Parameters {
        fmt.Printf("%v, %v, %v", p.ParameterType, p.Value, p.Name)
    }

    return &gm.Empty{}, nil
}

func (h *handler) NotifyScenarioExecutionStarting(c context.Context, m *gm.ScenarioExecutionStartingRequest) (*gm.Empty, error) {
    fmt.Println("Received ScenarioExecutionStartingRequest")
    fmt.Printf("Current scenario name: %v\n", m.CurrentExecutionInfo.CurrentScenario.Name)

    escapedStep := url.QueryEscape(m.CurrentExecutionInfo.CurrentScenario.Name)
    uri := fmt.Sprintf("http://localhost:18080/notifySpec?name=%v", escapedStep)
    fmt.Printf("url: %v", uri)

    resp, err := http.Get(uri)
    if err != nil {
        panic(err)
    }

    if resp.StatusCode == 200 {
        fmt.Println("Request succeeded")
    } else {
        panic(resp.StatusCode)
    }

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
