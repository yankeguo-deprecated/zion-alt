package zconf_crypto_tls

import (
	"crypto/tls"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCipherSuite(t *testing.T) {
	type testStruct struct {
		C CipherSuites `json:"c"`
	}
	var ts testStruct
	err := json.Unmarshal([]byte(`{"c":["TLS_RSA_WITH_AES_256_CBC_SHA"]}`), &ts)
	require.NoError(t, err)
	require.Equal(t, tls.TLS_RSA_WITH_AES_256_CBC_SHA, ts.C.UnwrapIDs()[0])
}
