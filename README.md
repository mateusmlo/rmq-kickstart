Honestly I just wanted to flex my golang muscles which had been unused for a while and this was a good enough mini project, which consists of going through the [RabbitMQ tutorials](https://www.rabbitmq.com/getstarted.html) and modernizing them (they won't work past go v1.14). Will do something more "production grade" next time.
To play around, first get RabbitMQ online by running `docker-compose up -d` with a .env file and then `go run cmd/tutorials/main.go` which takes an argument:
- `send` sends a hardcoded message to the broker
- `receive` should be called in another terminal, then receives the messages
- `task` secretly takes an additional arg that is the message to be parsed. Each "." is a second that it takes to complete. Uses default message if not provided.
- `worker` will process the messages received by the "task" queue and fake some time to return
- `emit_logs` will broadcast messages to any `receive_logs` terminal that is up

## TODO:

- [x] Tutorial 1
- [x] Tutorial 2
- [x] Tutorial 3
- [x] Tutorial 4
- [ ] Tutorial 5
- [ ] Tutorial 6
- [ ] Tutorial 7
