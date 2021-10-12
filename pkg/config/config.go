package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Projects struct {
	Projects []Project `json:"projects"`
}

type Cluster struct {
	Cluster string            `json:"cluster"`
	Region  string            `json:"region"`
	Tags    map[string]string `json:"tags"`
}

type Project struct {
	Project  string    `json:"project"`
	Clusters []Cluster `json:"clusters"`
}

func ExportConfig(p, c, r string) {

	var projects []Project
	var clusters []Cluster

	clusters = append(clusters, Cluster{Cluster: c, Region: r, Tags: map[string]string{"a": "b"}})
	projects = append(projects, Project{Project: p, Clusters: clusters})

	d := projects

	fmt.Println("Exported:", d)

	jSend := Projects{
		Projects: d,
	}

	// convert to JSON format
	//bytes, _ := json.Marshal(jSend)
	bytes, _ := json.MarshalIndent(jSend, "", "    ")

	fmt.Println(string(bytes))

	Save(bytes)
}

// Save create a data store with all clusters
func Save(b []byte) {
	file := "clusters.json"

	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		log.Println(err)
	}

	f.Write(b)
}

func ListProjectsByFile() {
	data, _ := os.ReadFile("clusters.json")
	if !json.Valid(data) {
		fmt.Println("Error: json file don't have json format")
		os.Exit(1)
	}

	projects := Projects{}
	json.Unmarshal(data, &projects)
	for i, _ := range projects.Projects {
		fmt.Printf("Project[%d]: %s\n", i, projects.Projects[i].Project)
	}
}

func GetProjectClusterByFile(p, c string) (string, string) {

	data, _ := os.ReadFile("clusters.json")
	if !json.Valid(data) {
		fmt.Println("Error: json file don't have json format")
		os.Exit(1)
	}

	projects := Projects{}
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
	data, _ := os.ReadFile("clusters.json")
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

	projects := Projects{}
	json.Unmarshal(data, &projects)

	for j, _ := range projects.Projects[id].Clusters {
		fmt.Printf("Cluster[%d]: %s\n", j, projects.Projects[id].Clusters[j].Cluster)
	}
}
