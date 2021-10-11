package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
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

func ExportConfig() {
	c := []Project{
		{
			Project: "localhost1",
			Clusters: []Cluster{{
				Cluster: "autopilot-gcp-gke1",
				Region:  "us-central1",
				Tags:    map[string]string{"a": "b"},
			},
				{
					Cluster: "autopilot-gcp-gke2",
					Region:  "us-central1",
					Tags:    map[string]string{"c": "d"},
				},
			},
		},
		{
			Project: "localhost2",
			Clusters: []Cluster{{
				Cluster: "autopilot-gcp-gke3",
				Region:  "us-central1",
				Tags:    map[string]string{"a": "b"},
			},
			},
		},
	}

	jSend := Projects{
		Projects: c,
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

func ListProjects()  {
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

func GetProjectCluster(p, c string) (string, string){

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

func ListClusters(project string) {
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


func ListConfig() {

	data, _ := os.ReadFile("clusters.json")
	if !json.Valid(data) {
		fmt.Println("Error: json file don't have json format")
		os.Exit(1)
	}

	projects := Projects{}
	json.Unmarshal(data, &projects)
	for i, _ := range projects.Projects {
		fmt.Printf("Project[%d]: %s\n", i, projects.Projects[i].Project)
		for j, _ := range projects.Projects[i].Clusters {
			fmt.Printf("\tCluster[%d]: %s\n", j, projects.Projects[i].Clusters[j].Cluster)
		}
	}

	project := projects.Projects[2].Project
	cluster := projects.Projects[2].Clusters[0].Cluster
	region := projects.Projects[2].Clusters[0].Region
	ConnectCluster(project, cluster, region)
}

func LoadConfig(p, c string) {
	fmt.Printf("Project: %s\t Cluster: %s\n", p, c)
}

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
