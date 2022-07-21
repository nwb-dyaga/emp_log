package rmq

import (
	"log"
)

func (rcl *RabbitClient) Consume(n string, f func([]byte) bool) {
	for {
		_, err := rcl.channel(true, true)
		if err == nil {
			break
		}
	}
	log.Printf("--- connected to consume '%s' ---\r\n", n)
	q, err := rcl.recChan.QueueDeclare(
		n,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	m, err := rcl.recChan.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("--- failed to consume ---")
	}
	var forever chan struct{}
	go func() {
		for d := range m {
			f(d.Body)
		}
	}()
	<-forever
}
