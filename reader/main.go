package reader

import (
	"os"

	"gopkg.in/yaml.v3"
)
type Config struct {
	Name	string   `yaml:"name"`
	WorkDir string `yaml:"work-dir"`
	Commands []string `yaml:"sequence"`
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
