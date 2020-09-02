package zconf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	yamlTestDoc = `
hello:
  world:
    truth: false
`
)

func TestYamlLoader_Load(t *testing.T) {
	l := yamlLoader{}
	m := map[string]interface{}{}
	err := l.Load("test.json", []byte(yamlTestDoc), &m)
	assert.NoError(t, err)
	truth := m["hello"].(map[string]interface{})["world"].(map[string]interface{})["truth"].(bool)
	assert.False(t, truth)
}
