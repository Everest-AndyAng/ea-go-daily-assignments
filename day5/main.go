package main

import (
	"encoding/json"
	"os"
)

type Job struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func main() {

}

func jobsFromFile() []Job {
	var jobs []Job
	data, _ := os.ReadFile("jobs.json")
	json.Unmarshal(data, &jobs)
	return jobs
}
