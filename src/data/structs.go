package data

const DB = "db"
const SERVER = "server"
const WEB = "web"

// Definition stores information about a system, used for importing data.
type Definition struct {
	ID          string
	Name        string
	Type        string
	User        string
	Password    string
	URL         string
	Server      string
	DB          string
	Host        string
	Port        string
	SID         string
	Env         string
	Location    []string
	Responsible []string
	Notes       []string
}

// YamlData stores information about all systems, used for importing data.
type YamlData struct {
	Defs []Definition
}

// Config stores data about a system, used for exporting data.
type Config struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	User        string   `json:"user"`
	Password    string   `json:"password,omitempty"`
	URL         string   `json:"url"`
	Server      string   `json:"server,omitempty"`
	DB          string   `json:"db,omitempty"`
	Host        string   `json:"host,omitempty"`
	Port        string   `json:"port,omitempty"`
	SID         string   `json:"sid,omitempty"`
	Env         string   `json:"env"`
	Location    []string `json:"location,omitempty"`
	Responsible []string `json:"responsible,omitempty"`
	Notes       []string `json:"notes,omitempty"`

	IsValid bool `json:"-"`
}
