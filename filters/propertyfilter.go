package filters

import d "github.com/appPlant/alpinepass/data"

//PropertyFilter allows filtering a Config for the content of its properties.
type PropertyFilter struct {
	Slices []string
}

func (p PropertyFilter) filter(data d.Config) d.Config {
	//var remove = false
	return data
}

func filterProperty() {

}
