package config

import (
	"encoding/json"
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
	a[0] = bytes
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

	for i:= 0; i < len(c); i++ {
		f.Write(c[i])
	}
}
