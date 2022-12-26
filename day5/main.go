package main

import (
	"encoding/json"
	"log"
	"os"
)

type Job struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type Status struct {
	Id     int32  `json:"id"`
	Status string `json:"status"`
}

func main() {
	jobs := jobsFromFile()
	statuses := executeJobs(jobs)
	writeStatusToFile(statuses)
}

func jobsFromFile() []Job {
	var jobs []Job
	data, _ := os.ReadFile("jobs.json")
	json.Unmarshal(data, &jobs)
	return jobs
}

func writeStatusToFile(statuses []Status) {
	data, _ := json.Marshal(statuses)
	os.WriteFile("status.json", data, 0644)
}

func executeJob(job Job, status chan<- Status) {
	log.Println("executing ", job.Name)
	status <- Status{job.Id, "SUCCESS"}
}

func executeJobs(jobs []Job) []Status {
	var statuses []Status
	statusChannel := make(chan Status)
	for _, job := range jobs {
		go executeJob(job, statusChannel)
		status, _ := <-statusChannel
		statuses = append(statuses, status)
	}
	close(statusChannel)
	return statuses
}
