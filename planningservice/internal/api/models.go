package api

import (
	"github.com/mkaganm/providergrpc"
	"planningservice/internal/data"
	"planningservice/internal/planning"
)

type FailedResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type PlanningRequest struct {
}

type PlanningResponse struct {
	Status string `json:"status"`
	Data   Datum  `json:"data"`
}

type Datum struct {
	Planning   planning.Plan       `json:"planning"`
	Developers []data.Developer    `json:"developers"`
	Tasks      []providergrpc.Task `json:"tasks"`
}
