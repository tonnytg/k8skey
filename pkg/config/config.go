package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

func ListConfig() {

	data, _ := os.ReadFile("clusters.json")
	if ! json.Valid(data) {
		fmt.Println("Error: json file don't have json format")
		os.Exit(1)
	}

	projects := Projects{}
	json.Unmarshal(data, &projects)
	for i, _ := range projects.Projects {
		fmt.Printf("Project[%d]: %s\n", i,projects.Projects[i].Project)
		for j, _ := range projects.Projects[i].Clusters {
			fmt.Printf("\tCluster[%d]: %s\n", j,projects.Projects[i].Clusters[j].Cluster)
		}
	}


}

func LoadConfig(p, c string) {
	fmt.Printf("Project: %s\t Cluster: %s\n", p, c)
}
