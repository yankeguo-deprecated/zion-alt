package zconf_crypto_tls

import (
	"crypto/tls"
	"encoding/json"
)

type CertificateSrc struct {
	Certificate     string `json:"certificate"`
	CertificateFile string `json:"certificate_file"`

	PrivateKey     string `json:"private_key"`
	PrivateKeyFile string `json:"private_key_file"`
}

func (c CertificateSrc) Build() (tls.Certificate, error) {
	if c.CertificateFile != "" && c.PrivateKeyFile != "" {
		return tls.LoadX509KeyPair(c.CertificateFile, c.PrivateKeyFile)
	}
	return tls.X509KeyPair([]byte(c.Certificate), []byte(c.PrivateKey))
}

type Certificate tls.Certificate

func (c Certificate) Unwrap() tls.Certificate {
	return tls.Certificate(c)
}

func (c *Certificate) UnmarshalJSON(bytes []byte) (err error) {
	var src CertificateSrc
	if err = json.Unmarshal(bytes, &src); err != nil {
		return
	}
	var tc tls.Certificate
	if tc, err = src.Build(); err != nil {
		return
	}
	*c = Certificate(tc)
	return
}

type Certificates []Certificate

func (cs Certificates) Unwrap() []tls.Certificate {
	if cs == nil {
		return nil
	}
	ret := make([]tls.Certificate, 0, len(cs))
	for _, c := range cs {
		ret = append(ret, c.Unwrap())
	}
	return ret
}
