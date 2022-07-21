package rmq

import (
	"github.com/streadway/amqp"
)

type RabbitClient struct {
	sendConn *amqp.Connection
	recConn  *amqp.Connection
	sendChan *amqp.Channel
	recChan  *amqp.Channel
}
