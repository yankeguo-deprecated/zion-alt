package zconf_crypto_tls

import (
	"crypto/tls"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCurveID(t *testing.T) {
	type testStruct struct {
		Curves CurveIDs `json:"curves"`
	}
	var ts testStruct
	err := json.Unmarshal([]byte(`{"curves":["pxxx","x25519"]}`), &ts)
	require.Error(t, err)
	ts = testStruct{}
	err = json.Unmarshal([]byte(`{"curves":["p521","x25519"]}`), &ts)
	require.NoError(t, err)
	require.Equal(t, tls.X25519, ts.Curves[1].Unwrap())
}
