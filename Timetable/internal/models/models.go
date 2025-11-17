package models

import (
	"github.com/gofrs/uuid"
)

type Event struct {
	Id    *uuid.UUID `json:"id"`
	UcaId int        `json:"ucaId"`
	Name  string     `json:"name"`
}
