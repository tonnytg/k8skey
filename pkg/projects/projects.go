package projects

import (
	"encoding/json"
	"fmt"
	"k8skey/pkg/clusters"
	"k8skey/pkg/http"
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

// List get all projects to create a database
func List() {
	// Get a key and export to authentication
	// gcloud auth application-default print-access-token
	// export GCP_TOKEN="ya29.a0ARrdaM_6DfC..."

	// You can use sub-shell but sometimes gcloud return with update info
	// export GCP_TOKEN=`gcloud auth application-default print-access-token`
	url := "https://cloudresourcemanager.googleapis.com/v1/projects"
	b, err := http.GetRequest(url)
	if err != nil {
		fmt.Println("error:", err)
	}

	var p AllProjects

	if err := json.Unmarshal(b, &p); err != nil {
		panic(err)
	}

	for i, v := range p.Projects {
		fmt.Printf("Project[%d]: %s\n", i, v.ProjectID)
		clusters.GetClustersK8s(v.ProjectID)
	}
}
