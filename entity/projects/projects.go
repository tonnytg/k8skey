package projects

import (
	"encoding/json"
	"fmt"
	"k8skey/entity/clusters"
	"k8skey/pkg/http"
	"time"
)

// ProjectsRough used to convert return of API to struct format
type ProjectsRough struct {
	Projects []struct {
		ProjectNumber string    `json:"projectNumber"`
		ProjectID     string    `json:"projectId"`
		Name          string    `json:"name"`
		CreateTime    time.Time `json:"createTime"`
	} `json:"projects"`
}

type Projects struct {
	Projects []Project
}

type Project struct {
	Project  string
	Clusters []clusters.Cluster
}

// GetProjects get all projects to create a database
func GetProjects() []Project {
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

	// like a diamond needs to be cut
	var rough ProjectsRough

	// decode return of API to struct format
	if err := json.Unmarshal(b, &rough); err != nil {
		panic(err)
	}

	var p []Project

	for i, v := range rough.Projects {
		fmt.Printf("Project[%d]: %s\n", i, v.ProjectID)
		c := clusters.GetClustersK8s(v.ProjectID)
		p = append(p, Project{Project: v.ProjectID, Clusters: c})
	}

	fmt.Println(p)
	return p
}
