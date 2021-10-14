package config

import (
	"encoding/json"
	"fmt"
	"k8skey/entity/projects"
	"log"
	"os"
	"strconv"
	"strings"
)

func ExportConfig() {

	p := projects.GetProjects()

	fmt.Println("Exported:", p)

	jSend := projects.Projects{
		Projects: p,
	}

	// convert string to JSON format
	//bytes, _ := json.Marshal(jSend) // json format in each line
	bytes, _ := json.MarshalIndent(jSend, "", "    ") // json format in many lines

	fmt.Println(string(bytes))

	Save(bytes)
}

// Save create a data store with all clusters
func Save(b []byte) {

	dir, err := os.UserHomeDir()
	err = os.Mkdir(dir + "/.k8skey", 0755)
	if err != nil {
		log.Println(err)
	}

	file := dir + "/.k8skey/clusters.json"

	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		log.Println(err)
	}

	f.Write(b)
}

func ListProjectsByFile() {
	dir, _ := os.UserHomeDir()
	file := dir + "/.k8skey/clusters.json"

	data, _ := os.ReadFile(file)
	if !json.Valid(data) {
		fmt.Println("Error: json file don't have json format")
		os.Exit(1)
	}

	projects := projects.Projects{}
	json.Unmarshal(data, &projects)
	for i, _ := range projects.Projects {
		fmt.Printf("Project[%d]: %s\n", i, projects.Projects[i].Project)
	}
}

func GetProjectClusterByFile(p, c string) (string, string) {

	dir, _ := os.UserHomeDir()
	file := dir + "/.k8skey/clusters.json"
	data, _ := os.ReadFile(file)
	if !json.Valid(data) {
		fmt.Println("Error: json file don't have json format")
		os.Exit(1)
	}

	projects := projects.Projects{}
	json.Unmarshal(data, &projects)

	pr := strings.Trim(p, "\n")
	cl := strings.Trim(c, "\n")
	idpr, _ := strconv.Atoi(pr)
	idcl, _ := strconv.Atoi(cl)

	fmt.Printf("Project[%d]: %s\n", idpr, projects.Projects[idpr].Project)
	project := projects.Projects[idpr].Project

	fmt.Printf("Cluster[%d]: %s\n", idcl, projects.Projects[idpr].Clusters[idcl].Cluster)
	cluster := projects.Projects[idpr].Clusters[idcl].Cluster
	return project, cluster
}

func ListClustersByFile(project string) {

	dir, _ := os.UserHomeDir()
	file := dir + "/.k8skey/clusters.json"
	data, _ := os.ReadFile(file)
	if !json.Valid(data) {
		fmt.Println("Error: json file don't have json format")
		os.Exit(1)
	}

	//id := strconv.Atoi(project)
	// remove \n
	p := strings.Trim(project, "\n")
	// parse string to int
	id, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println(err)
	}

	projects := projects.Projects{}
	json.Unmarshal(data, &projects)

	for j, _ := range projects.Projects[id].Clusters {
		fmt.Printf("Cluster[%d]: %s\n", j, projects.Projects[id].Clusters[j].Cluster)
	}
}
