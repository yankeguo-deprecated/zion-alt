package zconf_time

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type testStructD struct {
	D Duration `json:"d"`
}

func TestDuration_Duration(t *testing.T) {
	d := Duration(time.Second * 3)
	td := testStructD{
		D: d,
	}
	buf, _ := json.Marshal(td)
	assert.Equal(t, `{"d":"3s"}`, string(buf))
	var td1 testStructD
	_ = json.Unmarshal(buf, &td1)
	assert.Equal(t, time.Second*3, td1.D.Unwrap())
	err := json.Unmarshal([]byte(`{}`), &td1)
	assert.NoError(t, err)
}
