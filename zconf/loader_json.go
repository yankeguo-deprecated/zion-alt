package zconf

import "encoding/json"

type jsonLoader struct {
}

func (j *jsonLoader) Name() string {
	return "JSON"
}

func (j *jsonLoader) Extensions() []string {
	return []string{".json"}
}

func (j *jsonLoader) Load(filename string, buf []byte, out *map[string]interface{}) error {
	return json.Unmarshal(buf, out)
}
