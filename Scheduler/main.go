package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// Retrieve data from EDT
	resp, err := http.Get("https://edt.uca.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=13295,13345&projectId=3&calType=ical&nbWeeks=4&displayConfigId=128")
	if err != nil {
		panic(err) // TODO done: manage error
	}
	defer resp.Body.Close()

	// Read all and store in rawData
	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err) // TODO done: manage error
	}

	// Create a line-reader from data
	scanner := bufio.NewScanner(bytes.NewReader(rawData))

	// Create vars
	var eventArray []map[string]string
	currentEvent := map[string]string{}

	currentKey := ""
	currentValue := ""

	inEvent := false

	// Inspect each line
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore calendar lines until we reach an event
		if !inEvent && line != "BEGIN:VEVENT" {
			continue
		}

		// If new event, go to next line
		if line == "BEGIN:VEVENT" {
			inEvent = true
			currentEvent = map[string]string{}
			continue
		}

		if line == "END:VEVENT" {
			inEvent = false
			eventArray = append(eventArray, currentEvent)
			continue
		}

		// Lines starting with a space = continuation of previous key
		if strings.HasPrefix(line, " ") {
			currentValue += strings.TrimSpace(line)
			currentEvent[currentKey] = currentValue
			continue
		}

		// Split scan
		splitted := strings.SplitN(line, ":", 2)
		if len(splitted) != 2 {
			continue // Invalid line, skip
		}

		currentKey = splitted[0]
		currentValue = splitted[1]

		// Store current event attribute
		currentEvent[currentKey] = currentValue
	}

	// You could define a struct, but here we stay simple.

	jsonData, err := json.MarshalIndent(eventArray, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
