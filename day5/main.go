package main

import (
	"encoding/json"
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
