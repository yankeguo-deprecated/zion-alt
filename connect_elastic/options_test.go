package connect_elastic

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"testing"
)

const (
	elasticTestData = `
urls:
  - http://127.0.0.1:9200
  - http://127.0.0.1:8200
basic_auth:
  username: admin
  password: qwerty
gzip: true
sniffer:
  enabled: true
  interval: 5s
  timeout: 5s
  timeout_startup: 10s
health_check:
  enabled: true
  interval: 5s
  timeout: 5s
  timeout_startup: 10s
retrier:
  backoff: exponential
  timeout_min: 200ms
  timeout_max: 20s
`
)

func TestLoad(t *testing.T) {
	var opts Options
	var m map[string]interface{}
	err := yaml.Unmarshal([]byte(elasticTestData), &m)
	buf, _ := json.Marshal(m)
	err = json.Unmarshal(buf, &opts)
	require.NoError(t, err)
	require.Equal(t, 12, len(opts.Unwrap()))
}
