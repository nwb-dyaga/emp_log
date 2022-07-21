package rmq

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

func (rcl *RabbitClient) channel(isRec, recreate bool) (*amqp.Channel, error) {
	if recreate {
		if isRec {
			rcl.recChan = nil
		} else {
			rcl.sendChan = nil
		}
	}
	if isRec && rcl.recConn == nil {
		rcl.recChan = nil
	}
	if !isRec && rcl.sendConn == nil {
		rcl.recChan = nil
	}
	if isRec && rcl.recChan != nil {
		return rcl.recChan, nil
	} else if !isRec && rcl.sendChan != nil {
		return rcl.sendChan, nil
	}
	for {
		_, err := rcl.connect(isRec, recreate)
		if err == nil {
			break
		}
	}
	var err error
	if isRec {
		rcl.recChan, err = rcl.recConn.Channel()
	} else {
		rcl.sendChan, err = rcl.sendConn.Channel()
	}
	if err != nil {
		log.Println("--- could not create channel ---")
		time.Sleep(1 * time.Second)
		return nil, err
	}
	if isRec {
		return rcl.recChan, err
	} else {
		return rcl.sendChan, err
	}
}
