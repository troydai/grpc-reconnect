package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/troydai/cron"
	"github.com/troydai/grpc-reconnect/internal/socket"
	echopb "github.com/troydai/grpc-reconnect/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const _socketPath = "workspace/server.sock"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	conn, err := grpc.Dial("unix://"+socket.MustGetPath(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := echopb.NewEchoClient(conn)

	wait, err := cron.Start(
		ctx,
		cron.PlainJob(func() {
			resp, err := client.Status(ctx, &echopb.StatusRequest{})
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				return
			}

			fmt.Println("STATUS: ", resp.Status)

			echo, err := client.Echo(ctx, &echopb.EchoRequest{Message: uuid.New().String()})
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				return
			}

			fmt.Println("  ECHO: ", echo.Message)
		}),
		cron.WithInterval(500*time.Millisecond),
	)
	if err != nil {
		log.Fatal(err)
	}

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt)
	fmt.Println("Ctrl+C to exit")

	select {
	case <-term:
		fmt.Println("Exiting ...")
		cancel()
	case <-wait:
		os.Exit(0)
	}

	select {
	case <-wait:
		fmt.Println("Terminated")
	case <-time.After(time.Second):
		fmt.Println("Terminated after timeout")
	}
}
