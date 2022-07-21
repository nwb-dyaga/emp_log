package main

import (
	"go_test/controllers"
	"go_test/rmq"
)

func main() {
	var rc rmq.RabbitClient
	rc.Consume("test_queue", controllers.CreateHttpLog)
}
