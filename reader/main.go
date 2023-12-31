package reader

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Sequence struct {
	Commands []string `yaml:"sequence"`
}

func Read(path string) Sequence {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	res := &Sequence{}
	err = yaml.Unmarshal(data, res)

	if err != nil {
		panic(err)
	}

	return *res
}
