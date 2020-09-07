package zconf_crypto_tls

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"strconv"
)

var (
	clientAuthTypes = map[string]tls.ClientAuthType{
		"no":                 tls.NoClientCert,
		"request":            tls.RequestClientCert,
		"require-any":        tls.RequireAnyClientCert,
		"verify-if-given":    tls.VerifyClientCertIfGiven,
		"require-and-verify": tls.RequireAndVerifyClientCert,
	}
)

type ClientAuthType tls.ClientAuthType

func (c ClientAuthType) Unwrap() tls.ClientAuthType {
	return tls.ClientAuthType(c)
}

func (c ClientAuthType) MarshalJSON() ([]byte, error) {
	for k, v := range clientAuthTypes {
		if v == c.Unwrap() {
			return json.Marshal(k)
		}
	}
	return nil, errors.New("unknown tls.ClientAuthType: " + strconv.Itoa(int(c)))
}

func (c *ClientAuthType) UnmarshalJSON(bytes []byte) (err error) {
	var s string
	if err = json.Unmarshal(bytes, &s); err != nil {
		return
	}
	for k, v := range clientAuthTypes {
		if k == s {
			*c = ClientAuthType(v)
			return
		}
	}
	err = errors.New("unknown tls.ClientAuthType: " + s)
	return
}
