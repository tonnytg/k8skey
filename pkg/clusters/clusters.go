package clusters

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/container/v1"
	"log"
)

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
