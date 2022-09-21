package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/troydai/grpc-reconnect/internal/echoserver"
	"github.com/troydai/grpc-reconnect/internal/socket"
	echopb "github.com/troydai/grpc-reconnect/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	socketPath := socket.MustGetPath()
	socket.MustEnsureSocket(socketPath)

	lis, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	echopb.RegisterEchoServer(server, echoserver.New())
	defer func() {
		server.Stop()
		os.Remove(socketPath)
	}()

	go func() {
		fmt.Println("Listening on ", socketPath, " ...")
		server.Serve(lis)
	}()

	waitTerm := make(chan os.Signal, 1)
	signal.Notify(waitTerm, os.Interrupt)

	fmt.Println("Ctrl+C to exit")
	<-waitTerm
}
