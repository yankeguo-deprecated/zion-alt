package conf

import "gopkg.in/yaml.v3"

type yamlLoader struct {
}

func (y *yamlLoader) Name() string {
	return "YAML"
}

func (y *yamlLoader) Extensions() []string {
	return []string{".yaml", ".yml"}
}

func (y *yamlLoader) Load(filename string, buf []byte, out *map[string]interface{}) error {
	return yaml.Unmarshal(buf, out)
}
