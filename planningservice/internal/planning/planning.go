package planning

import (
	"planningservice/internal/data"
	"planningservice/internal/providergrpc"
	"sort"
)

// Planning is a function that plans a week
func Planning(devs []data.Developer, providersTasks []providergrpc.Task) Plan {

	var taskList []Task

	// convert providergrpc.Task to Task
	for _, task := range providersTasks {
		taskList = append(taskList, ConvertTask(&task))
	}

	// sort tasks by estimation
	taskList = sortTasks(taskList)

	var tasks *[]Task
	tasks = &taskList
	var plan Plan

	i := 0
	for len(*tasks) > 0 {

		i++

		// plan a week
		weekPlan := weekPlanner(i, devs, tasks)
		plan.WeekPlans = append(plan.WeekPlans, weekPlan)

	}

	return plan
}

// weekPlanner is a function that plans a week
func weekPlanner(weekNumber int, devs []data.Developer, tasks *[]Task) WeekPlan {

	// assign tasks to developers
	devTasks := taskAssigner(devs, tasks)

	return WeekPlan{
		WeekNumber:     weekNumber,
		DeveloperTasks: devTasks,
	}

}

// taskAssigner is a function that assigns tasks to developers
func taskAssigner(developers []data.Developer, tasks *[]Task) []DeveloperTasks {

	var devTasks []DeveloperTasks

	// sort developers by estimation
	developers = sortDevelopers(developers)

	// create a developerTasks
	for _, dev := range developers {
		devTasks = append(devTasks, DeveloperTasks{
			Developer: dev,
			Tasks:     []Task{},
		})
	}

	// assign tasks index list
	var assignedTasks []int

	j := 0
	for i := 0; i < len(*tasks); i++ {

		flag := false

		if j >= len(devTasks) {
			j = 0
		}

		for j < len(devTasks) {

			if devTasks[j].Developer.Estimation >= (*tasks)[i].Estimation {

				// assign task to developer
				devTasks[j].AssignTask((*tasks)[i])
				assignedTasks = append(assignedTasks, i)

				flag = true
				j++
				break

			} else {
				j++
			}
		}

		if flag {
			continue
		}
	}

	// remove assigned tasks from tasks
	removeTasksByIndexes(tasks, assignedTasks)

	return devTasks
}

// sortTasks is a function that sorts tasks by estimation
func sortTasks(tasks []Task) []Task {

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Estimation > tasks[j].Estimation
	})

	return tasks
}

// sortDevelopers is a function that sorts developers by estimation
func sortDevelopers(developers []data.Developer) []data.Developer {

	sort.Slice(developers, func(i, j int) bool {
		return developers[i].Estimation > developers[j].Estimation
	})
	return developers
}

// removeTask is a function that removes a task from tasks
func removeTasksByIndexes(tasks *[]Task, indexes []int) {

	// Create a set of indexes
	indexSet := make(map[int]bool)
	for _, index := range indexes {
		indexSet[index] = true
	}

	// New slice without removed tasks
	var newTasks []Task
	for i, task := range *tasks {
		if !indexSet[i] {
			newTasks = append(newTasks, task)
		}
	}

	// Update tasks
	*tasks = newTasks
}
