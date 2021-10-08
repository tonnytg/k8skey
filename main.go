package main

import (
	"encoding/json"
	"fmt"
	"k8skey/services/http"
	"time"
)

type AllProjects struct {
	Projects []struct {
		ProjectNumber string    `json:"projectNumber"`
		ProjectID     string    `json:"projectId"`
		Name          string    `json:"name"`
		CreateTime    time.Time `json:"createTime"`
	} `json:"projects"`
}

func main() {
	fmt.Println("List Projects")
	ListProjects()

	//GetClustersK8s()
}

// ListProjects get all projects to create a database
func ListProjects() {
	// gcloud auth application-default print-access-token
	// export GCP_TOKEN="ya29.a0ARrdaM_6DfC..."
	url := "https://cloudresourcemanager.googleapis.com/v1/projects"
	b, err := http.GetRequest(url)
	if err != nil {
		fmt.Println("error:", err)
	}

	var p AllProjects

	if err := json.Unmarshal(b, &p); err != nil {
		panic(err)
	}

	for _, v := range p.Projects {
		fmt.Println(v.ProjectID)
	}
}
