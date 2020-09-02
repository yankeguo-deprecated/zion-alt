package zconf_crypto_tls

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"strconv"
)

var (
	curveIDs = map[string]tls.CurveID{
		"p256":   tls.CurveP256,
		"p384":   tls.CurveP384,
		"p521":   tls.CurveP521,
		"x25519": tls.X25519,
	}
)

type CurveID tls.CurveID

func (c *CurveID) UnmarshalJSON(bytes []byte) (err error) {
	var s string
	if err = json.Unmarshal(bytes, &s); err != nil {
		return
	}
	for k, v := range curveIDs {
		if k == s {
			*c = CurveID(v)
			return
		}
	}
	err = errors.New("unknown tls.CurveID: " + s)
	return
}

func (c CurveID) MarshalJSON() ([]byte, error) {
	for k, v := range curveIDs {
		if v == tls.CurveID(c) {
			return json.Marshal(k)
		}
	}
	return nil, errors.New("unknown tls.CurveID: " + strconv.Itoa(int(c)))
}

func (c CurveID) Unwrap() tls.CurveID {
	return tls.CurveID(c)
}

type CurveIDs []CurveID

func (cs CurveIDs) Unwrap() []tls.CurveID {
	if cs == nil {
		return nil
	}
	ret := make([]tls.CurveID, 0, len(cs))
	for _, c := range cs {
		ret = append(ret, c.Unwrap())
	}
	return ret
}
