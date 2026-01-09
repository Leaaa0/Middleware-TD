package nats

import (
	"log"
	"strings"

	"github.com/nats-io/nats.go"
)

func InitJetStream() nats.JetStreamContext {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	jsc, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	_, err = jsc.AddStream(&nats.StreamConfig{
		Name:     "EVENTS",
		Subjects: []string{"EVENTS.>"},
	})
	if err != nil && !strings.Contains(err.Error(), "already in use") {
		log.Fatal(err)
	}
	return jsc
}
