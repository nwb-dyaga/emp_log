package rmq

import "os"

type Config struct {
	username string
	password string
	host     string
	port     string
}

func init() {
	var config Config
	config.username = os.Getenv("RMQ_USERNAME")
	config.password = os.Getenv("RMQ_PASSWORD")
	config.host = os.Getenv("RMQ_HOST")
	config.port = os.Getenv("RMQ_PORT")
}
