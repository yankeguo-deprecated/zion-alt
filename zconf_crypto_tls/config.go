package zconf_crypto_tls

import (
	"crypto/tls"
	"encoding/json"
	"github.com/zionkit/zion/zconf_crypto_x509"
)

type ConfigSrc struct {
	Certificates             Certificates               `json:"certificates"`
	RootCA                   zconf_crypto_x509.CertPool `json:"root_ca"`
	NextProtos               []string                   `json:"next_protos"`
	ServerName               string                     `json:"server_name"`
	ClientAuth               ClientAuthType             `json:"client_auth"`
	ClientCA                 zconf_crypto_x509.CertPool `json:"client_ca"`
	InsecureSkipVerify       bool                       `json:"insecure_skip_verify"`
	CipherSuites             CipherSuites               `json:"cipher_suites"`
	PreferServerCipherSuites bool                       `json:"prefer_server_cipher_suites"`
	SessionTicketsDisabled   bool                       `json:"session_tickets_disabled"`
	MinVersion               Version                    `json:"min_version"`
	MaxVersion               Version                    `json:"max_version"`
	CurvePreferences         CurveIDs                   `json:"curve_preferences"`
}

func (cs ConfigSrc) Build() *tls.Config {
	return &tls.Config{
		Certificates:             cs.Certificates.Unwrap(),
		RootCAs:                  cs.RootCA.Unwrap(),
		NextProtos:               cs.NextProtos,
		ServerName:               cs.ServerName,
		ClientAuth:               cs.ClientAuth.Unwrap(),
		ClientCAs:                cs.ClientCA.Unwrap(),
		InsecureSkipVerify:       cs.InsecureSkipVerify,
		CipherSuites:             cs.CipherSuites.UnwrapIDs(),
		PreferServerCipherSuites: cs.PreferServerCipherSuites,
		SessionTicketsDisabled:   cs.SessionTicketsDisabled,
		MinVersion:               cs.MinVersion.Unwrap(),
		MaxVersion:               cs.MaxVersion.Unwrap(),
		CurvePreferences:         cs.CurvePreferences.Unwrap(),
	}
}

type Config struct {
	*tls.Config
}

func (c *Config) UnmarshalJSON(bytes []byte) (err error) {
	var src ConfigSrc
	if err = json.Unmarshal(bytes, &src); err != nil {
		return
	}
	c.Config = src.Build()
	return
}

func (c Config) Unwrap() *tls.Config {
	return c.Config
}
