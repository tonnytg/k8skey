package cmd

import (
	"flag"
	"fmt"
	"k8skey/pkg/config"
	"k8skey/pkg/interactive"
	"os"
)

func Flags() {
	db := flag.String("database", "nil", "--database update")
	con := flag.Bool("select", false, "--select true")
	project := flag.String("project", "", "--project PROJECT_ID" )
	cluster := flag.String("cluster", "", "--cluster CLUSTER")
	region := flag.String("region", "", "--region REGION")

	flag.Parse()
	if *db == "update" {
		fmt.Println("Updating database of projects and clusters")
	}

	if *con == true {
		interactive.Menu()
	}
	if *project != "" {
		if *cluster != "" {
			if *region != "" {
				config.ConnectCluster(*project, *cluster, *region)
			} else {
				os.Exit(1)
			}
		} else {
			os.Exit(1)
		}
	} else {
		os.Exit(1)
	}
}
