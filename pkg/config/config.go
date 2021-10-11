package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Response struct {
	Clusters []Config `json:"clusters"`
}

type Config struct {
	Project string            `json:"project"`
	Cluster string            `json:"cluster"`
	Region  string            `json:"region"`
	Tags    map[string]string `json:"tags"`
}

func ExportConfig() {
	c := []Config{
		{
			Project: "localhost1",
			Cluster: "autopilot-gcp-gke",
			Region:  "us-central1",
			Tags:    map[string]string{"a": "b"},
		},
		{
			Project: "localhost2",
			Cluster: "autopilot-gcp-gke",
			Region:  "us-central1",
			Tags:    map[string]string{"a": "b"},
		},
		{
			Project: "localhost3",
			Cluster: "autopilot-gcp-gke",
			Region:  "us-central1",
			Tags:    map[string]string{"a": "b"},
		},
	}

	jSend := Response{
		Clusters: c,
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
	fmt.Println(json.Valid(data))

}

func LoadConfig(p, c string) {
	fmt.Printf("Project: %s\t Cluster: %s\n", p, c)
}
