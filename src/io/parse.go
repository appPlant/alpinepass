package io

import (
	yaml "gopkg.in/yaml.v2"

	d "github.com/appPlant/alpinepass/src/data"
	"github.com/appPlant/alpinepass/src/util"
)

//Separator is the character used for separating the parts of a generated ID.
const Separator string = "_"

func parseYaml(data string) d.YamlData {
	yamlData := d.YamlData{}
	err := yaml.Unmarshal([]byte(data), &yamlData)
	util.CheckError(err, "Parsing YAML failed!")
	return yamlData
}

func createID(definition d.Definition) string {
	//TODO find a way to generate an ID
	id := ""
	id = id + util.CleanString(definition.Env) + Separator
	id = id + util.CleanString(definition.Type) + Separator
	return id
}

func createConfig(def d.Definition) d.Config {
	config := d.Config{}

	if def.ID != "" {
		config.ID = def.ID
	} else {
		config.ID = createID(def)
	}

	config.Name = def.Name
	config.Type = def.Type
	config.User = def.User
	config.Password = def.Password
	config.URL = def.URL
	config.Server = def.Server
	config.DB = def.DB
	config.Host = def.Host
	config.Port = def.Port
	config.SID = def.SID
	config.Env = def.Env
	config.Location = def.Location
	config.Responsible = def.Responsible
	config.Notes = def.Notes

	config.IsValid = true

	return config
}

//ReadConfigs reads the input file and creates Config objects to work with.
func ReadConfigs(path string) []d.Config {
	data := readFile(path)
	yamlData := parseYaml(data)
	configs := []d.Config{}

	for _, definition := range yamlData.Defs {
		config := createConfig(definition)
		configs = append(configs, config)
	}

	return configs
}
