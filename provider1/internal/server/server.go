package server

import (
	"context"
	"provider1/internal/client"
	"provider1/internal/providergrpc"
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
func convertTasks(tasks []client.Task) ([]*providergrpc.Task, error) {

	var convertedTasks []*providergrpc.Task

	for _, task := range tasks {
		convertedTasks = append(convertedTasks, &providergrpc.Task{
			Name:     task.Name,
			Level:    int32(task.Level),
			Duration: int32(task.Duration),
		})
	}

	return convertedTasks, nil
}
