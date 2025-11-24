package models

import (
	"github.com/gofrs/uuid"
)

type Calendar struct {
	Id    *uuid.UUID `json:"id"`
	UcaId int        `json:"ucaId"`
	Name  string     `json:"name"`
}

type Alert struct {
	Id           *uuid.UUID `json:"id"`
	Resource     *uuid.UUID `json:"resource"`
	AllResources bool       `json:"allResources"`
	Mail         string     `json:"mail"`
}
