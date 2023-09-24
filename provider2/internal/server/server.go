package server

import (
	"context"
	"provider2/internal/client"
	"provider2/internal/providergrpc"
)

type Server struct {
	providergrpc.UnimplementedProviderServer
}

// GetTasks is a function that returns tasks
func (s *Server) GetTasks(ctx context.Context, req *providergrpc.Request) (*providergrpc.Response, error) {

	// Get tasks from provider1
	tasks, err := client.RequestTasks()
	if err != nil {
		return nil, err
	}

	// Convert tasks to providergrpc.Task
	returnTasks, err := convertTasks(tasks)

	// Return tasks
	return &providergrpc.Response{
		Tasks: returnTasks,
	}, nil
}

// convertTasks converts tasks to providergrpc.Task
func convertTasks(tasks []client.TaskMap) ([]*providergrpc.Task, error) {

	var convertedTasks []*providergrpc.Task

	for _, task := range tasks {

		for taskName, detail := range task {

			convertedTasks = append(convertedTasks, &providergrpc.Task{
				Name:     taskName,
				Level:    int32(detail.Level),
				Duration: int32(detail.EstimatedDuration),
			})

		}
	}

	return convertedTasks, nil
}
