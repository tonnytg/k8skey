package interactive

import (
	"bufio"
	"fmt"
	"k8skey/pkg/clusters"
	"k8skey/pkg/config"
	"os"
)

func Menu() {

	for {

		config.ListProjectsByFile()
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Choose Project: ")
		project, _ := reader.ReadString('\n')
		fmt.Printf("Project: %s", project)
		fmt.Println("---")

		if project == "\n" {
			fmt.Println("Project must be value")
			break
		}
		if project == "exit\n" {
			fmt.Println("Bye bye...")
			os.Exit(0)
		}

		config.ListClustersByFile(project)
		fmt.Print("Choose Cluster: ")
		cluster, _ := reader.ReadString('\n')
		fmt.Printf("Cluster: %s", cluster)
		fmt.Println("---")

		p, c := config.GetProjectClusterByFile(project, cluster)

		if cluster == "\n" {
			fmt.Println("Cluster must be value")
			break
		}
		r := "us-central1"
		clusters.ConnectCluster(p, c, r)
		os.Exit(0)
	}
}
