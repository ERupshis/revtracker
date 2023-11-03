package data

import (
	"github.com/erupshis/revtracker.git/db/models"
)

//go:generate easyjson -all data.go

type Data struct {
	Homework  models.Homework   `json:"Homework"`
	Questions []models.Question `json:"Questions"`
}

type FrontMessage struct {
	Data Data `json:"Data"`
}
