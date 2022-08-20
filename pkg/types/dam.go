package types

import (
	"encoding/json"

	"github.com/jptalukdar/dam-visualizer/pkg/utils"
	"sigs.k8s.io/yaml"
)

type DAM struct {
	Name        string        `json:"name"`
	Version     string        `json:"version"`
	Description string        `json:"description"`
	Author      string        `json:"author,omitempty"`
	Application []Application `json:"application"`
	Users       []string      `json:"users"`
}

type Application struct {
	Name    string         `json:"name"`
	Type    string         `json:"type"`
	Request string         `json:"request"`
	Depends []Dependencies `json:"depends-on"`
}

type Dependencies struct {
	Name       string `json:"name"`
	Access     string `json:"access"`
	TimeToLive string `json:"ttl,omitempty"`
	Type       string `json:"type"`
}

func (d *DAM) GenerateJson() string {
	json, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (d *DAM) GenerateYaml() string {
	j := d.GenerateJson()
	yaml, err := yaml.JSONToYAML([]byte(j))
	if err != nil {
		panic(err)
	}
	return string(yaml)
}

func (d *DAM) ReadYaml(path string) {
	y := utils.ReadFromFile(path)
	err := yaml.Unmarshal([]byte(y), d)
	if err != nil {
		panic(err)
	}
	d.Validate()
}

func (d *DAM) Validate() {
	if d.Name == "" {
		panic("DAM name is empty")
	}
}

func NewDAM() *DAM {
	return &DAM{}
}
