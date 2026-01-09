package main

import (
	_ "Scheduler/internal/models"
	internalNats "Scheduler/internal/nats"
	"Scheduler/internal/parser"
	"Scheduler/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/zhashkevych/scheduler"
)

var jsc nats.JetStreamContext

func runTask(ctx context.Context) {
	ids := repository.GetAgendaIds()
	url := fmt.Sprintf("https://edt.uca.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=%s&projectId=3&calType=ical&nbWeeks=1&displayConfigId=128", ids)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Erreur lors de la récupération de l'emploi du temps: %v", err)
		return
	}
	defer resp.Body.Close()

	events := parser.ParseICal(resp.Body)

	for _, event := range events {
		data, _ := json.Marshal(event)
		jsc.Publish("EVENTS.new", data)
	}
	log.Printf("%d événements synchronisés.", len(events))
}

func main() {
	jsc = internalNats.InitJetStream()

	ctx := context.Background()
	sc := scheduler.NewScheduler()

	// lance toute les 10 minutes
	sc.Add(ctx, runTask, time.Minute*10)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	sc.Stop()
}
