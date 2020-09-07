package zconf_crypto_tls

import (
	"crypto/tls"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClientAuthType(t *testing.T) {
	type testStruct struct {
		A ClientAuthType `json:"a"`
	}
	var ts testStruct
	err := json.Unmarshal([]byte(`{"a":"verify-if-given"}`), &ts)
	require.NoError(t, err)
	require.Equal(t, tls.VerifyClientCertIfGiven, ts.A.Unwrap())
}
