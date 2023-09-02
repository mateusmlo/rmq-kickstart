package main

import (
	"os"

	"github.com/mateusmlo/rabbitmq-hello-world/config"
	"github.com/mateusmlo/rabbitmq-hello-world/receive"
	"github.com/mateusmlo/rabbitmq-hello-world/send"
)

func main() {
	config.GetEnvs()

	arg1 := os.Args[1]
	ch, q, conn := config.ConnectAMQP()
	defer ch.Close()
	defer conn.Close()

	if arg1 == "send" {
		send.Send(ch, q)
	}

	if arg1 == "receive" {
		receive.Receive(ch, q)
	}
}
