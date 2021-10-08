package main

import (
	"context"
	"encoding/json"
	"fmt"
	"k8skey/services/http"
	"log"
	"time"

	"google.golang.org/api/container/v1"
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

func GetClustersK8s() {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, container.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	containerService, err := container.New(c)
	if err != nil {
		log.Fatal(err)
	}

	// The parent (project and location) where the clusters will be listed.
	// Specified in the format 'projects/*/locations/*'.
	// Location "-" matches all zones and all regions.
	parent := "projects/ultra-sound-324019/locations/-" // TODO: Update placeholder value.

	resp, err := containerService.Projects.Locations.Clusters.List(parent).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Change code below to process the `resp` object:
	fmt.Printf("%#v\n", resp)
}
