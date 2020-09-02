package zconf_crypto_tls

import (
	"crypto/tls"
	"encoding/json"
	"errors"
)

type CipherSuite struct {
	*tls.CipherSuite
}

func (c CipherSuite) Unwrap() *tls.CipherSuite {
	return c.CipherSuite
}

func (c CipherSuite) UnwrapID() uint16 {
	if c.CipherSuite == nil {
		return 0
	} else {
		return c.CipherSuite.ID
	}
}

func (c *CipherSuite) UnmarshalJSON(bytes []byte) (err error) {
	var s string
	if err = json.Unmarshal(bytes, &s); err != nil {
		return
	}
	for _, cs := range tls.CipherSuites() {
		if cs.Name == s {
			c.CipherSuite = cs
			return
		}
	}
	for _, cs := range tls.InsecureCipherSuites() {
		if cs.Name == s {
			c.CipherSuite = cs
			return
		}
	}
	err = errors.New("unknown tls.CipherSuite: " + s)
	return
}

func (c CipherSuite) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.CipherSuite.ID)
}

type CipherSuites []CipherSuite

func (cs CipherSuites) Unwrap() []*tls.CipherSuite {
	if cs == nil {
		return nil
	}
	ret := make([]*tls.CipherSuite, 0, len(cs))
	for _, c := range cs {
		ret = append(ret, c.Unwrap())
	}
	return ret
}

func (cs CipherSuites) UnwrapIDs() []uint16 {
	if cs == nil {
		return nil
	}
	ret := make([]uint16, 0, len(cs))
	for _, c := range cs {
		ret = append(ret, c.UnwrapID())
	}
	return ret
}
