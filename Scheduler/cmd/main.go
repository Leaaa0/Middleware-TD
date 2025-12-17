package main

import (
	"Scheduler/models"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/zhashkevych/scheduler"
)

var jsc nats.JetStreamContext

func initNATS() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	jsc, err = nc.JetStream()
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
}

func getAgendaIds() string {

	//resp, err := http.Get("http://localhost:8080/agendas")
	//if err != nil {
	//	log.Printf("Erreur : Impossible de joindre l'API Config : %v", err)
	//	return ""
	//}
	//defer resp.Body.Close()

	//var agendas []Agenda
	//if err := json.NewDecoder(resp.Body).Decode(&agendas); err != nil {
	//	log.Printf("Erreur décodage JSON Config : %v", err)
	//	return ""
	//}

	// extrait UcaIDs puis joint avec des virgules
	//var ids []string
	//for _, a := range agendas {
	//	if a.UcaID != "" {
	//		ids = append(ids, a.UcaID)
	//	}
	//}

	//return strings.Join(ids, ",")
	return "13295,7224,7225,62962,62090"
}

func runTask(ctx context.Context) {
	ids := getAgendaIds()
	url := fmt.Sprintf("https://edt.uca.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=%s&projectId=3&calType=ical&nbWeeks=1&displayConfigId=128", ids)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Erreur HTTP: %v", err)
		return
	}
	defer resp.Body.Close()

	rawData, _ := io.ReadAll(resp.Body)
	scanner := bufio.NewScanner(bytes.NewReader(rawData))

	inEvent := false
	currentEventMap := map[string]string{}

	for scanner.Scan() {
		line := scanner.Text()
		if !inEvent && line != "BEGIN:VEVENT" {
			continue
		}
		if line == "BEGIN:VEVENT" {
			inEvent = true
			currentEventMap = map[string]string{}
			continue
		}

		if line == "END:VEVENT" {
			inEvent = false
			event := models.Event{
				UID:         currentEventMap["UID"],
				Summary:     currentEventMap["SUMMARY"],
				Location:    currentEventMap["LOCATION"],
				Description: currentEventMap["DESCRIPTION"],
			}
			data, _ := json.Marshal(event)
			_, err := jsc.Publish("EVENTS.new", data)
			if err != nil {
				log.Printf("Erreur NATS: %v", err)
			}
			continue
		}

		splitted := strings.SplitN(line, ":", 2)
		if len(splitted) == 2 {
			currentEventMap[splitted[0]] = splitted[1]
		}
	}
	log.Println("Tâche de synchronisation terminée.")

	return
}

func main() {
	initNATS()

	ctx := context.Background()
	sc := scheduler.NewScheduler()

	// à changer car c'est toute les 10secs
	sc.Add(ctx, runTask, time.Second*10)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	sc.Stop()
}
