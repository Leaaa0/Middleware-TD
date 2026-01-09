package repository

import (
	"Scheduler/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func GetAgendaIds() string {

	resp, err := http.Get("http://localhost:8091/calendars")
	if err != nil {
		log.Printf("Erreur : Impossible de joindre l'API Config : %v", err)
		return ""
	}
	defer resp.Body.Close()

	var calendar []models.Calendar
	if err := json.NewDecoder(resp.Body).Decode(&calendar); err != nil {
		log.Printf("Erreur décodage JSON Config : %v", err)
		return ""
	}

	// extrait UcaIDs puis joint avec des virgules
	var ids []string
	for _, a := range calendar {
		// On vérifie que l'ID n'est pas vide (0)
		if a.UcaID != 0 {
			// Converti int vers string pour l'URL
			ids = append(ids, fmt.Sprintf("%d", a.UcaID))
		}
	}

	return strings.Join(ids, ",")
}
