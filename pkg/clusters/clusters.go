package clusters

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/container/v1"
	"k8skey/pkg/config"
	"log"
	"os"
	"os/exec"
	"syscall"
)

// GetClustersK8s get clusters name from project
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
		config.ExportConfig(p, c.Name, c.Location)
	}
}

// ConnectCluster get-credential from cluster, projects and region
func ConnectCluster(p, c, r string) {

	arg1 := "container"
	arg2 := "clusters"
	arg3 := "get-credentials"
	arg4 := c // Cluster
	arg5 := "--region"
	arg6 := r // region
	arg7 := "--project"
	arg8 := p // Project

	binary, lookErr := exec.LookPath("gcloud")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"gcloud", arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
	// Finish
	os.Exit(0)
}
