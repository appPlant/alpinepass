package filters

import (
	"strings"

	"regexp"

	d "github.com/appPlant/alpinepass/src/data"
	"github.com/appPlant/alpinepass/src/util"
	"github.com/fatih/structs"
)

//Separator is tha character used for separating filter keys and values.
const Separator string = "="

//PropertyFilter allows filtering a Config for the content of its properties.
type PropertyFilter struct {
	Slices []string
}

func (p PropertyFilter) filter(data d.Config) d.Config {
	verifyFlags(p.Slices) //TODO verify only once and not for every Config
	for i := 0; i < len(p.Slices); i++ {
		data = filterProperty(p.Slices[i], data)
	}
	return data
}

func filterProperty(property string, data d.Config) d.Config {
	var split = strings.Split(property, Separator)
	var key = split[0]
	var value = split[1]
	var regex = regexp.MustCompile(value)

	//TODO rather use reflection?
	switch key {
	case "id":
		if !regex.MatchString(data.ID) {
			data.IsValid = false
		}
	case "location":
		if !regex.MatchString(data.Location) {
			data.IsValid = false
		}
	case "environment":
		if !regex.MatchString(data.Env) {
			data.IsValid = false
		}
	case "user":
		if !regex.MatchString(data.User) {
			data.IsValid = false
		}
	case "host":
		if !regex.MatchString(data.Host) {
			data.IsValid = false
		}
	}

	return data
}

//verifyFlags checks that the input flags are valid.
func verifyFlags(flags []string) {
	for i := 0; i < len(flags); i++ {
		flag := flags[i]

		//Do some simple checks.
		if !strings.Contains(flag, Separator) {
			util.ThrowError("The filter does not contain '" + Separator + "'! Flag: " + flag)
		}
		if strings.Count(flag, Separator) > 1 {
			util.ThrowError("The filter contains too many '" + Separator + "'! Flag: " + flag)
		}

		//check that the key part of a flag matches the fields available in struct Config.
		data := d.Config{}
		fieldNames := structs.Names(data)
		key := strings.ToLower(strings.Split(flag, Separator)[0])
		isContained := false
		for j := 0; j < len(fieldNames); j++ {
			if key == strings.ToLower(fieldNames[j]) {
				isContained = true
			}
		}
		if !isContained {
			util.ThrowError("The filter type is not valid! Flag: " + flag)
		}
	}
}
