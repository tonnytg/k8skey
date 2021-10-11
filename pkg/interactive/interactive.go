package interactive

import (
	"bufio"
	"fmt"
	"os"
)

func Menu() {

	for {
		//ShowProjects() TODO: Create a function to show all projects
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Choose Project: ")
		project, _ := reader.ReadString('\n')
		fmt.Printf("Project: %s\n", project)

		if project == "\n" {
			fmt.Println("Project must be value")
			break
		}
		if project == "exit\n" {
			fmt.Println("Bye bye...")
			os.Exit(0)
		}

		//ShowClusters() TODO: Create a function to show all clusters
		fmt.Print("Choose Cluster: ")
		cluster, _ := reader.ReadString('\n')
		fmt.Printf("Cluster: %s\n", cluster)

		if cluster == "\n" {
			fmt.Println("Cluster must be value")
			break
		}
	}
}
