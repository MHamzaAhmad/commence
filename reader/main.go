package reader

import (
	"os"

	"gopkg.in/yaml.v3"
)
type Config struct {
	Name	string   `yaml:"name"`
	Commands []Command `yaml:"sequence"`
}

type Command struct {
	Command string `yaml:"cmd"`
	SubCommand []string `yaml:"then,omitempty"`
}


func Read(path string) Config {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	res := &Config{}
	err = yaml.Unmarshal(data, res)

	if err != nil {
		panic(err)
	}

	return *res
}
