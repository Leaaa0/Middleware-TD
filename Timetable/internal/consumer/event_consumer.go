package consumer

import (
	"context"
	"encoding/json"
	"middleware/timetable/internal/models"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
)

func EventConsumer(nc *nats.Conn) (*jetstream.Consumer, error) {
	// On crée le contexte JetStream à partir de la connexion reçue
	js, err := jetstream.New(nc)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// On récupère le stream USERS ou EVENTS (selon ton scheduler)
	stream, err := js.Stream(ctx, "EVENTS")
	if err != nil {
		return nil, err
	}

	// On crée le consumer durable
	consumer, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable: "timetable_processor",
		Name:    "timetable_processor",
	})

	return &consumer, err
}

func Consume(consumer jetstream.Consumer) error {
	cc, err := consumer.Consume(func(msg jetstream.Msg) {
		// 1. Décodage du message reçu du Scheduler
		var event models.Event
		json.Unmarshal(msg.Data(), &event)

		logrus.Infof("Message reçu : %s", event.Summary)

		// 2. TODO : Appeler ta logique de stockage ici
		// saveToDatabase(event)

		msg.Ack() // On confirme la lecture
	})
	if err != nil {
		return err
	}

	// Garde le consumer ouvert
	select {}
	<-cc.Closed()
	cc.Stop()

	return err
}
