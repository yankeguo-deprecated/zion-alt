package zconf_crypto_tls

import (
	"crypto/tls"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"testing"
)

const (
	configTestData = `c:
  certificates:
    - certificate: |
        -----BEGIN CERTIFICATE-----
        MIIBhTCCASugAwIBAgIQIRi6zePL6mKjOipn+dNuaTAKBggqhkjOPQQDAjASMRAw
        DgYDVQQKEwdBY21lIENvMB4XDTE3MTAyMDE5NDMwNloXDTE4MTAyMDE5NDMwNlow
        EjEQMA4GA1UEChMHQWNtZSBDbzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABD0d
        7VNhbWvZLWPuj/RtHFjvtJBEwOkhbN/BnnE8rnZR8+sbwnc/KhCk3FhnpHZnQz7B
        5aETbbIgmuvewdjvSBSjYzBhMA4GA1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggr
        BgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1UdEQQiMCCCDmxvY2FsaG9zdDo1
        NDUzgg4xMjcuMC4wLjE6NTQ1MzAKBggqhkjOPQQDAgNIADBFAiEA2zpJEPQyz6/l
        Wf86aX6PepsntZv2GYlA5UpabfT2EZICICpJ5h/iI+i341gBmLiAFQOyTDT+/wQc
        6MF9+Yw1Yy0t
        -----END CERTIFICATE-----
      private_key: |
        -----BEGIN EC PRIVATE KEY-----
        MHcCAQEEIIrYSSNQFaA2Hwf1duRSxKtLYX5CB04fSeQ6tF1aY/PuoAoGCCqGSM49
        AwEHoUQDQgAEPR3tU2Fta9ktY+6P9G0cWO+0kETA6SFs38GecTyudlHz6xvCdz8q
        EKTcWGekdmdDPsHloRNtsiCa697B2O9IFA==
        -----END EC PRIVATE KEY-----
  cipher_suites:
    - TLS_AES_128_GCM_SHA256
    - TLS_AES_128_GCM_SHA256
  min_version: tls1.1
  max_version: tls1.3
  curve_preferences:
    - p521
    - x25519
`
)

func TestConfig(t *testing.T) {
	type testStruct struct {
		C Config `json:"c"`
	}
	var m map[string]interface{}
	var ts testStruct
	err := yaml.Unmarshal([]byte(configTestData), &m)
	buf, _ := json.Marshal(m)
	err = json.Unmarshal(buf, &ts)
	require.NoError(t, err)
	tc := ts.C.Unwrap()
	require.Equal(t, tls.TLS_AES_128_GCM_SHA256, tc.CipherSuites[1])
	require.Equal(t, uint16(tls.VersionTLS11), tc.MinVersion)
	require.Equal(t, uint16(tls.VersionTLS13), tc.MaxVersion)
	require.Equal(t, tls.X25519, tc.CurvePreferences[1])
}
