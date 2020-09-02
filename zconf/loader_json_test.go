package conf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	jsonTestDoc = `{
"hello": {
  "world": {
    "truth": false
  }
}
}
`
)

func TestJsonLoader_Load(t *testing.T) {
	l := jsonLoader{}
	m := map[string]interface{}{}
	err := l.Load("test.json", []byte(jsonTestDoc), &m)
	assert.NoError(t, err)
	truth := m["hello"].(map[string]interface{})["world"].(map[string]interface{})["truth"].(bool)
	assert.False(t, truth)
}
