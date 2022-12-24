package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSouldReturnJobsFromFile(t *testing.T) {
	var jobs []Job
	jobs = append(jobs, Job{342, "build-app"})
	jobs = append(jobs, Job{123, "build-server"})
	jobs = append(jobs, Job{78, "build-shared-library"})

	result := jobsFromFile()

	assert.Equal(t, jobs, result)
}

func TestShouldWriteStatusToFile(t *testing.T) {
	var statuses []Status
	statuses = append(statuses, Status{342, "SUCCESS"})
	statuses = append(statuses, Status{123, "FAILURE"})
	statuses = append(statuses, Status{78, "SUCCESS"})

	writeStatusToFile(statuses)

	var statusFromFile []Status
	data, _ := os.ReadFile("status.json")
	json.Unmarshal(data, &statusFromFile)

	assert.Equal(t, statuses, statusFromFile)
}
