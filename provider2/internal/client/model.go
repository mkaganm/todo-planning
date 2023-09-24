package client

// TaskMap is a map that represents a task from provider response
type TaskMap map[string]TaskDetail

type TaskDetail struct {
	Level             int `json:"level"`
	EstimatedDuration int `json:"estimated_duration"`
}
