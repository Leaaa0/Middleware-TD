package parser

import (
	"Scheduler/internal/models"
	"bufio"
	"io"
	"strings"
)

func ParseICal(r io.Reader) []models.Event {
	var events []models.Event
	scanner := bufio.NewScanner(r)
	currentEventMap := map[string]string{}
	inEvent := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "BEGIN:VEVENT" {
			inEvent = true
			currentEventMap = map[string]string{}
			continue
		}
		if line == "END:VEVENT" && inEvent {
			e := models.Event{
				UID:         currentEventMap["UID"],
				Summary:     currentEventMap["SUMMARY"],
				Location:    currentEventMap["LOCATION"],
				Description: currentEventMap["DESCRIPTION"],
			}
			// Récupération des dates
			for key, val := range currentEventMap {
				if strings.HasPrefix(key, "DTSTART") {
					e.DTStart = val
				}
				if strings.HasPrefix(key, "DTEND") {
					e.DTEnd = val
				}
			}
			events = append(events, e)
			inEvent = false
			continue
		}
		if inEvent {
			splitted := strings.SplitN(line, ":", 2)
			if len(splitted) == 2 {
				currentEventMap[splitted[0]] = splitted[1]
			}
		}
	}
	return events
}
