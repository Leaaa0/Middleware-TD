package nats

import (
	natsgo "github.com/nats-io/nats.go"

	"log"
	"strings"
)

func InitJetStream() natsgo.JetStreamContext {
	nc, err := natsgo.Connect(natsgo.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	jsc, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	_, err = jsc.AddStream(&natsgo.StreamConfig{
		Name:     "EVENTS",
		Subjects: []string{"EVENTS.>"},
	})
	if err != nil && !strings.Contains(err.Error(), "already in use") {
		log.Fatal(err)
	}
	return jsc
}
