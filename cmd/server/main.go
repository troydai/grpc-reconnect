package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const _socketPath = "workspace/server.sock"

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	socketPath := path.Join(cwd, _socketPath)
	if _, err := os.Stat(socketPath); !os.IsNotExist(err) {
		fmt.Println("Removing socket ", socketPath)
		os.Remove(socketPath)
	}

	fmt.Println(socketPath)

	lis, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	defer func() {
		server.Stop()
		os.Remove(socketPath)
	}()

	go func() {
		server.Serve(lis)
	}()

	waitTerm := make(chan os.Signal, 1)
	signal.Notify(waitTerm, os.Interrupt)

	fmt.Println("Ctrl+C to exit")
	<-waitTerm
}
