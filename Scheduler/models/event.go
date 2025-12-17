package models

type Event struct {
	UID         string `json:"uid"`
	Summary     string `json:"summary"`
	Location    string `json:"location"`
	Description string `json:"description"`
	DTStart     string `json:"dtstart"`
	DTEnd       string `json:"dtend"`
}
