package validation

import (
	"github.com/appPlant/alpinepass/src/data"
	"github.com/appPlant/alpinepass/src/util"
)

//Validate validates the input.
func Validate(configs []data.Config) {
	for i := 0; i < len(configs); i++ {
		verifyType(configs[i])
	}
}

func verifyType(config data.Config) {
	if !util.StringInArray(config.Type, []string{"db", "server", "web"}) {
		util.ThrowConfigError("The Config has an invalid type!", config)
	}
}
