package models

import (
	"github.com/gofrs/uuid"
)

type Calendar struct {
	Id    *uuid.UUID `json:"id"`
	UcaId int        `json:"ucaId"`
	Name  string     `json:"name"`
}

type Alerts struct {
	Id            *uuid.UUID `json:"id"`
	ressource     *uuid.UUID `json:"ressource"`
	allRessources bool       `json:"allRessources"`
	mail          string     `json:"mail"`
}
