package rmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func (rcl *RabbitClient) connect(isRec, reconnect bool) (*amqp.Connection, error) {
	if reconnect {
		if isRec {
			rcl.recConn = nil
		} else {
			rcl.sendConn = nil
		}
	}
	if isRec && rcl.recConn != nil {
		return rcl.recConn, nil
	} else if !isRec && rcl.sendConn != nil {
		return rcl.sendConn, nil
	}
	var c string
	var config = Config{}
	if config.username == "" {
		c = fmt.Sprintf("amqp://%s:%s/", config.host, config.port)
	} else {
		c = fmt.Sprintf("amqp://%s:%s@%s/", config.username, config.password, config.host)
	}
	conn, err := amqp.Dial(c)
	if err != nil {
		log.Printf("\r\n--- could not create a conection ---\r\n")
		time.Sleep(1 * time.Second)
		return nil, err
	}
	if isRec {
		rcl.recConn = conn
		return rcl.recConn, nil
	} else {
		rcl.sendConn = conn
		return rcl.sendConn, nil
	}
}
