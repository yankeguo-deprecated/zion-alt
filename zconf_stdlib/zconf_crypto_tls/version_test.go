package zconf_crypto_tls

import (
	"crypto/tls"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVersion(t *testing.T) {
	type testStruct struct {
		V Version `json:"v"`
	}
	var ts testStruct
	err := json.Unmarshal([]byte(`{"v":"tls1.2"}`), &ts)
	require.NoError(t, err)
	require.Equal(t, uint16(tls.VersionTLS12), ts.V.Unwrap())
}
