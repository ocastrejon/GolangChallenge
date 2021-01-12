package main

import (
	"strconv"
	"strings"
)

// processData. This function receives the three values of the scenarios device, and
// call another functions to finally get the file format config
func processData(resourceCapacity, foregroundTasks, backgroundTasks string) string {
	arrayForegroundTasks := getResourceConsumption(foregroundTasks)
	arrayBackgroundTasks := getResourceConsumption(backgroundTasks)
	config := makeConfiguration(resourceCapacity, arrayForegroundTasks, arrayBackgroundTasks)
	fileFormatConfig := makeFinalFormat(config)
	return fileFormatConfig
}

// getResourceConsumption. Receives the foreground or background configuration tasks,
// and process the string to filter just the resource consumption value of the task.
func getResourceConsumption(tasks string) []int {

	// Remove all the "(" and ")"
	noParentesis := strings.Replace(strings.Replace(tasks, "(", "", -1), ")", "", -1)

	// Create an array of values separated for ", "
	eachTask := strings.Split(noParentesis, ", ")

	var tasksArray []int

	// Iterates over eachTask array to remove all the spaces and clean the string
	for i := 1; i < len(eachTask); i += 2 {
		noSpaceTask := strings.Replace(eachTask[i], " ", "", -1)

		// Finally append the integer value of the resource consumption task, in the tasksArray
		tasksArray = append(tasksArray, stringToInt(noSpaceTask))
	}

	return tasksArray
}

// makeConfiguration. Gets the optimous configuration, but it works to know if the device scenarios has an
// optimous configuration, else, iterates until get any optimous configuration
func makeConfiguration(resourceCapacity string, arrayForegroundTasks, arrayBackgroundTasks []int) []string {
	newResourceCapacity := stringToInt(resourceCapacity)

	// Get the optimous configuration with all the resource used.
	configurationArray := everyConfiguration(newResourceCapacity, arrayForegroundTasks, arrayBackgroundTasks, 0)

	i := 0
	// It repeats the everyConfiguration function until get any configuration, but with resource to spare.
	for len(configurationArray) == 0 {
		i++
		configurationArray = everyConfiguration(newResourceCapacity, arrayForegroundTasks, arrayBackgroundTasks, i)
	}

	return configurationArray
}

// everyConfiguration. This function iterates over foreground and background tasks to compare each
// resource consumption and finally decides if it is the optimous configuration.
// The flag parameter works to know if the configuration has the optimous configuration, else, needs
// decrement until find some optimous configuration.
func everyConfiguration(resourceCapacity int, arrayForegroundTasks, arrayBackgroundTasks []int, flag int) []string {
	var configurationArray []string

	// First for loop iterates over the foreground tasks
	for i := 0; i < len(arrayForegroundTasks); i++ {

		// Second for loop iterates over the background tasks
		for m := 0; m < len(arrayBackgroundTasks); m++ {
			sum := arrayForegroundTasks[i] + arrayBackgroundTasks[m]

			// If the sum of the foreground task resource consumption and the background
			//  task resource consumption matches with the resource capacity, storage the optimous configuration.
			if sum+flag == resourceCapacity {
				config := "(" + strconv.Itoa(i+1) + ", " + strconv.Itoa(m+1) + ")"
				configurationArray = append(configurationArray, config)
			}
		}
	}

	return configurationArray
}
