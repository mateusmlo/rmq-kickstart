package main

import (
	"os"

	"github.com/mateusmlo/rabbitmq-hello-world/config"
	"github.com/mateusmlo/rabbitmq-hello-world/internal/logs"
	"github.com/mateusmlo/rabbitmq-hello-world/internal/receive"
	"github.com/mateusmlo/rabbitmq-hello-world/internal/send"
	"github.com/mateusmlo/rabbitmq-hello-world/internal/task"
	"github.com/mateusmlo/rabbitmq-hello-world/internal/worker"
)

func main() {
	config.GetEnvs()

	arg1 := os.Args[1]
	ch, conn := config.ConnectAMQP()
	defer ch.Close()
	defer conn.Close()

	q, _ := config.CreateQueue(ch, "hello", true, false)
	config.CreateQueue(ch, "", false, true)
	config.CreateExchange(ch, "logs")

	if arg1 == "send" {
		send.Send(ch, &q)
	}

	if arg1 == "receive" {
		receive.Receive(ch, &q)
	}

	if arg1 == "task" {
		task.Task(ch, &q)
	}

	if arg1 == "worker" {
		worker.Worker(ch, &q)
	}

	if arg1 == "emit_logs" {
		logs.EmitLogs(ch)
	}

	if arg1 == "receive_logs" {
		logs.ReceiveLogs(ch)
	}
}
