package cmd

import (
	"flag"
	"fmt"
	"k8skey/pkg/clusters"
	"k8skey/pkg/interactive"
	"k8skey/pkg/projects"
	"log"
	"os"
)

func Flags() {
	db := flag.String("database", "nil", "--database update")
	con := flag.Bool("select", false, "--select true")
	prompt := flag.String("prompt", "off", "--prompt on")
	project := flag.String("project", "", "--project PROJECT_ID")
	cluster := flag.String("cluster", "", "--cluster CLUSTER")
	region := flag.String("region", "", "--region REGION")

	flag.Parse()

	if *prompt == "on" {
		SetShell()
		os.Exit(0)
	}

	if *db == "update" {
		fmt.Println("Updating database of projects and clusters")
		projects.List()
		os.Exit(0)
	}

	if *con == true {
		interactive.Menu()
	}
	if *project != "" {
		if *cluster != "" {
			if *region != "" {
				clusters.ConnectCluster(*project, *cluster, *region)
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

func SetShell() {

	dir, err := os.UserHomeDir()
	err = os.Mkdir(dir+"/.k8skey", 0755)

	file := dir + "/.k8skey/kube_ps1.zsh"

	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		log.Println(err)
	}

	f.WriteString("PROMPT='$(kube_ps1)'$PROMPT\n")

	fmt.Println("source  ~/.k8skey/kube_ps1.zsh")
	os.Exit(0)
}
