package zconf_tls_x509

import (
	"crypto/x509"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
)

type CertPoolSrc struct {
	System           bool     `json:"system"`
	Certificates     []string `json:"certificates"`
	CertificateFiles []string `json:"certificate_files"`
}

func (c CertPoolSrc) CertPool() (p *x509.CertPool, err error) {
	if c.System {
		if p, err = x509.SystemCertPool(); err != nil {
			return
		}
	} else {
		p = x509.NewCertPool()
	}
	for i, cert := range c.Certificates {
		if !p.AppendCertsFromPEM([]byte(cert)) {
			err = errors.New("failed to load x509.CertPool, .certs#" + strconv.Itoa(i))
			return
		}
	}
	for i, certFile := range c.CertificateFiles {
		var buf []byte
		if buf, err = ioutil.ReadFile(certFile); err != nil {
			return
		}
		if !p.AppendCertsFromPEM(buf) {
			err = errors.New("failed to load x509.CertPool, .cert_files#" + strconv.Itoa(i))
			return
		}
	}
	return
}

type CertPool struct {
	*x509.CertPool
}

func (c *CertPool) UnmarshalJSON(bytes []byte) (err error) {
	var src CertPoolSrc
	if err = json.Unmarshal(bytes, &src); err != nil {
		return
	}
	c.CertPool, err = src.CertPool()
	return
}

func (c CertPool) Unwrap() *x509.CertPool {
	return c.CertPool
}
