package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

type Config struct {
}

type Module struct {
}

func New(c Config) Module {

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("topic", "channel", config)
	if err != nil {
		log.Fatal(err)
	}

	return Module{}
}
