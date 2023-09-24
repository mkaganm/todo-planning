package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mkaganm/providergrpc"
	"log"
	"planningservice/internal/client"
	"planningservice/internal/config"
	"planningservice/internal/data"
	"planningservice/internal/planning"
	"sync"
)

// Planning service
func Planning(c *fiber.Ctx) error {

	// Provider ports
	var providerPorts []string
	providerPorts = append(providerPorts, config.EnvConfigs.Provider1Port)
	providerPorts = append(providerPorts, config.EnvConfigs.Provider2Port)

	// Provider tasks
	var tasks []providergrpc.Task

	// create a mutex
	var mutex sync.Mutex
	// create a WaitGroup
	var wg sync.WaitGroup

	// set number of goroutines to wait for
	wg.Add(len(providerPorts))

	// concurrent calls to the provider
	for _, port := range providerPorts {

		go func(port string) {

			// defer the WaitGroup
			defer wg.Done()

			// call the GetTasks function
			result, err := client.GetTasks(port)
			if err != nil {
				log.Default().Println("Error while calling the GetTasks function: ", err)
				return
			}

			// lock the mutex
			mutex.Lock()
			// Dereference each task in result and append to tasks
			for _, taskPtr := range result {
				tasks = append(tasks, *taskPtr)
			}
			// unlock the mutex
			mutex.Unlock()

		}(port)

	}

	// Wait for all goroutines to finish
	wg.Wait()

	developers, err := data.GetDevelopers()
	if err != nil {
		log.Default().Println("Error while calling the GetDevelopers function: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(FailedResponse{
			Status:  "failed",
			Message: "internal server error",
			Error:   err.Error(),
		})
	}

	taskList := tasks
	plan := planning.Planning(developers, tasks)

	return c.Status(fiber.StatusOK).JSON(PlanningResponse{
		Status: "success",
		Data: Datum{
			Planning:   plan,
			Developers: developers,
			Tasks:      taskList,
		},
	})

}
