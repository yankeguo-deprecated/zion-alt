package zconf_time

import (
	"encoding/json"
	"time"
)

type Duration time.Duration

func (d Duration) Unwrap() time.Duration {
	return time.Duration(d)
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

func (d *Duration) UnmarshalJSON(bytes []byte) (err error) {
	var s string
	if err = json.Unmarshal(bytes, &s); err != nil {
		return
	}
	var x time.Duration
	if x, err = time.ParseDuration(s); err != nil {
		return
	}
	*d = Duration(x)
	return
}

func (d Duration) MarshalText() (text []byte, err error) {
	text = []byte(time.Duration(d).String())
	return
}

func (d *Duration) UnmarshalText(text []byte) (err error) {
	var x time.Duration
	if x, err = time.ParseDuration(string(text)); err != nil {
		return
	}
	*d = Duration(x)
	return
}
