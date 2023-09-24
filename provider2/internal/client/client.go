package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"provider2/internal/config"
)

// RequestTasks requests tasks from provider1 returns a slice of tasks
func RequestTasks() ([]TaskMap, error) {

	// Get tasks from provider1
	resp, err := http.Get(config.EnvConfigs.Provider2Url)
	if err != nil {
		log.Default().Println("Error getting tasks from provider1: ", err)
		return nil, err
	}

	// Close the response body when everything is done
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Default().Println("Error closing response body: ", err)
			return
		}
	}(resp.Body)

	// Check the status code
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse the response body
	tasks, err := parseResponse(resp)
	if err != nil {
		log.Default().Println("Error parsing response: ", err)
		return nil, err
	}

	return tasks, nil
}

// parseResponse parses the response body and returns a slice of tasks
func parseResponse(resp *http.Response) ([]TaskMap, error) {

	var tasks []TaskMap

	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
