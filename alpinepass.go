package main

import (
	"strings"

	"github.com/ghodss/yaml"
)

// Separator separates the different parts of an ID
const Separator string = "."

// Filler is used to replace whitespace in strings
const Filler string = "-"

func parseData(data string) YamlData {
	yamlData := YamlData{}
	err := yaml.Unmarshal([]byte(data), &yamlData)
	checkError(err)
	return yamlData
}

func createConfig(definition Definition) Config {
	config := Config{}
	config.Title = definition.Title
	config.ID = createID(definition)
	config.Password = definition.Password
	config.User = definition.User
	//TODO host
	return config
}

func createID(definition Definition) string {
	id := ""
	id = id + cleanString(definition.Location) + Separator
	id = id + cleanString(definition.Env) + Separator
	id = id + cleanString(definition.Type) + Separator
	id = id + cleanString(definition.Title)
	return id
}

func cleanString(s string) string {
	return strings.Replace(s, " ", Filler, -1)
}

func filterConfig(config Config) Config {
	return config
}

func main() {
	runApp()
}
