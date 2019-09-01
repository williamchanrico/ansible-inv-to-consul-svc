package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/williamchanrico/ansible-inv-to-consul-svc/aini"
	"gopkg.in/yaml.v3"
)

// Data struct contains consul_services ansible variable
// consul_services:
//   - name: servicename
//     address: "x.x.x.x"
//     tags:
//       - abc
type Data struct {
	ConsulServices []ConsulService `yaml:"consul_services"`
}

// ConsulService struct contains every consul service entry
type ConsulService struct {
	Address string   `yaml:"address"`
	Name    string   `yaml:"name"`
	Tags    []string `yaml:"tags,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		log.Printf("usage: %v inventory_file", os.Args[0])
		os.Exit(1)
	}

	inventoryFile, err := aini.NewFile(os.Args[1])
	if err != nil {
		log.Fatal("failed to read inventory file:", err)
	}

	data := Data{}
	for hostname, group := range inventoryFile.Groups {
		for _, host := range group {
			svc := ConsulService{
				Name:    normalizeHostgroup(hostname),
				Address: host.Name,
			}
			if host.ConsulTag != "" {
				svc.Tags = []string{host.ConsulTag}
			}

			data.ConsulServices = append(data.ConsulServices, svc)
		}
	}

	out, err := yaml.Marshal(data)
	if err != nil {
		log.Fatal("error marshal:", err)
	}
	fmt.Println(string(out))
}

func normalizeHostgroup(s string) string {
	s = strings.TrimPrefix(s, "old-")
	return s
}
