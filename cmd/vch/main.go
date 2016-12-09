package main

import (
	"fmt"
	"github.com/begizi/vch-client/client"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = "localhost:9001"
	}

	client, err := client.NewVCHClient(serverAddr)
	if err != nil {
		panic(err)
	}

	// Error chan
	errc := make(chan error)

	// Interrupt handler
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errc <- client.Tunnel()
	}()

	fmt.Println("Exit: ", <-errc)
}
