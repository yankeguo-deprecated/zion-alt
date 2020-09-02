package zconf_crypto_tls

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"strconv"
)

var (
	versions = map[string]uint16{
		"tls1.0": tls.VersionTLS10,
		"tls1.1": tls.VersionTLS11,
		"tls1.2": tls.VersionTLS12,
		"tls1.3": tls.VersionTLS13,
		"ssl3.0": tls.VersionSSL30,
	}
)

type Version uint16

func (ver Version) Unwrap() uint16 {
	return uint16(ver)
}

func (ver *Version) UnmarshalJSON(bytes []byte) (err error) {
	var s string
	if err = json.Unmarshal(bytes, &s); err != nil {
		return
	}
	for k, v := range versions {
		if k == s {
			*ver = Version(v)
			return
		}
	}
	err = errors.New("unknown version: " + s)
	return
}

func (ver Version) MarshalJSON() ([]byte, error) {
	for k, v := range versions {
		if v == uint16(ver) {
			return json.Marshal(k)
		}
	}
	return nil, errors.New("unknown version: 0x" + strconv.FormatUint(uint64(ver), 16))
}
