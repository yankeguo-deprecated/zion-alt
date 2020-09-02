package zconf_redis_v7

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"testing"
)

const (
	redisTestData = `
network: tcp
address: 127.0.0.1:6378
database: 2
username: hello
password: world
`
	redisClusterTestData = `
addresses:
  - 127.0.0.1:6378
username: hello
password: world
`
)

func TestLoad(t *testing.T) {
	var opts Options
	var m map[string]interface{}
	err := yaml.Unmarshal([]byte(redisTestData), &m)
	buf, _ := json.Marshal(m)
	err = json.Unmarshal(buf, &opts)
	assert.NoError(t, err)
	require.Equal(t, opts.Address, "127.0.0.1:6378")
}

func TestLoadCluster(t *testing.T) {
	var opts ClusterOptions
	var m map[string]interface{}
	err := yaml.Unmarshal([]byte(redisClusterTestData), &m)
	buf, _ := json.Marshal(m)
	err = json.Unmarshal(buf, &opts)
	assert.NoError(t, err)
	require.Equal(t, opts.Password, "world")
}
