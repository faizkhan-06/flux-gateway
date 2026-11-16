package config

type Gateway struct {
	Server Server   `yaml:"server"`
	Routes []Routes `yaml:"routes"`
}

type Server struct {
	Port int `yaml:"port"`
}

type Routes struct {
	Path   string `yaml:"path"`
	Target string `yaml:"target"`
}
