package planning

import (
	"github.com/mkaganm/providergrpc"
	"planningservice/internal/data"
)

type Plan struct {
	WeekPlans []WeekPlan `json:"week_plans"`
}

type WeekPlan struct {
	WeekNumber     int              `json:"week_number"`
	DeveloperTasks []DeveloperTasks `json:"developer_tasks"`
}

// DeveloperTasks is a struct that represents a developer with his tasks
type DeveloperTasks struct {
	Developer data.Developer `json:"developer"`
	Tasks     []Task         `json:"tasks"`
}

// AssignTask is a function that assigns a task to a developer
func (dev *DeveloperTasks) AssignTask(task Task) {

	dev.Tasks = append(dev.Tasks, task)
	dev.Developer.Estimation -= task.Estimation

}

type Task struct {
	Name       string `json:"name"`
	Level      int    `json:"level"`
	Duration   int    `json:"duration"`
	Estimation int    `json:"estimation"`
}

// ConvertTask converts a providergrpc.Task to a Task
func ConvertTask(t *providergrpc.Task) Task {
	return Task{
		Name:       t.Name,
		Level:      int(t.Level),
		Duration:   int(t.Duration),
		Estimation: int(t.Level * t.Duration),
	}
}
