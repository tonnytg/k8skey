package clusters

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/container/v1"
	"log"
	"os"
	"os/exec"
	"syscall"
)

type Cluster struct {
	Cluster string `json:"cluster"`
	Region  string `json:"region"`
}

// GetClustersK8s get clusters name from project
func GetClustersK8s(project string) []Cluster {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, container.CloudPlatformScope)
	if err != nil {
		log.Println(err)
	}

	containerService, err := container.New(c)
	if err != nil {
		log.Println(err)
	}

	// The parent (project and location) where the clusters will be listed.
	// Specified in the format 'projects/*/locations/*'.
	// Location "-" matches all zones and all regions.
	parent := fmt.Sprintf("projects/%s/locations/-", project)

	resp, err := containerService.Projects.Locations.Clusters.List(parent).Context(ctx).Do()
	if err != nil {
		log.Println(err)
	}

	var clusters []Cluster

	for i, c := range resp.Clusters {
		fmt.Printf("Cluster[%d]: %s - %s\n", i, c.Name, c.Location)
		clusters = append(clusters, Cluster{Cluster: c.Name, Region: c.Location})
	}
	return clusters
}

// ConnectCluster get-credential from cluster, projects and region
func ConnectCluster(project, cluster, region string) {

	args := []string {
		"gcloud",			// command gcloud
		"container",		// arguments of gcloud to connect cluster
		"clusters",
		"get-credentials",
		cluster,
		"--region",
		region,
		"--project",
		project,
	}

	// find gcloud on PATH
	binary, lookErr := exec.LookPath("gcloud")
	if lookErr != nil {
		panic(lookErr)
	}

	env := os.Environ()

	// execErr run gcloud with arguments of args
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
	// Finish return 0 to no problem
	os.Exit(0)
}
