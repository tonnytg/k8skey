package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/google"
	"k8skey/pkg/config"
	"k8skey/pkg/http"
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

	config.ExportConfig()

	//GetClustersK8s()
}

// ListProjects get all projects to create a database
func ListProjects() {
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

	for _, v := range p.Projects {
		fmt.Println("Project:", v.ProjectID)
		GetClustersK8s(v.ProjectID)
	}
}

func GetClustersK8s(p string) {
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
	//parent := "projects/ultra-sound-324019/locations/-" // TODO: Update placeholder value.
	parent := fmt.Sprintf("projects/%s/locations/-", p)

	resp, err := containerService.Projects.Locations.Clusters.List(parent).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%#v\n", resp)
	for i, c := range resp.Clusters {
		fmt.Printf("Cluster[%d]: %s - %s\n", i, c.Name, c.Location)
	}
}
