package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Project string
	Cluster string
	Region  string
	Tags    map[string]string
}

func ExportConfig() {
	c := Config{
		Project: "localhost",
		Cluster: "autopilot-gcp-gke",
		Region:  "us-central1",
		Tags:    map[string]string{"a": "b"},
	}

	// convert to JSON format
	bytes, _ := json.Marshal(c)

	a := map[int][]byte{}
	for i:=0; i < 3; i++ {
		a[i] = bytes
	}

	Save(a)
}

// Save create a data store with all clusters
func Save(c map[int][]byte) {
	file := "clusters.db"

	f, err := os.Create(file)
	defer f.Close()

	if err != nil {
		log.Println(err)
	}

	var s []string
	for i:= 0; i < len(c); i++ {
		s = append(s, string(c[i]))
		s = append(s, "\n")
	}

	for i, _ := range s {
		f.WriteString(s[i])
	}
}

func ListConfig() {
	fmt.Println("a", "b")
}

func LoadConfig(p, c string) {
	fmt.Printf("Project: %s\t Cluster: %s\n", p, c)
}