package main

import (
	"github.com/stretchr/testify/assert"
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
