package main

import (
	"bufio"
	"os"
)

// scanningFile. Receives the file that contain the input data.
// It scans line by line of the file
func scanningFile(file *os.File) (string, error) {
	var allConfig string
	scanner := bufio.NewScanner(file)
	var resourceCapacity, foregroundTasks, backgroundTasks string
	i := 0

	// for that scans line by line to storage its values
	for scanner.Scan() {
		i++
		switch i {
		case 1:
			// First line has de resource capacity
			resourceCapacity = scanner.Text()
		case 2:
			// Second line has the foreground tasks
			foregroundTasks = scanner.Text()
		case 3:
			// Third line has the background tasks
			backgroundTasks = scanner.Text()

			// Send to the processData function this three values
			allConfig += processData(resourceCapacity, foregroundTasks, backgroundTasks) + "\n"
			i = 0
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return allConfig, nil
}

// createFile. Creates the final file (challenge.out) with the optimous configuration
func createFile(finalConfig string) error {
	outputFile, err := os.Create("challenge.out")
	if err != nil {
		return err
	}

	// Write the information in the file
	outputFile.WriteString(finalConfig)
	defer outputFile.Close()

	return nil
}
