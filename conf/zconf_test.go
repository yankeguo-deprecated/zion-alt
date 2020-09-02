package conf

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

type testStruct struct {
	Hello struct {
		World struct {
			Truth bool `json:"t"`
		} `json:"w"`
	} `json:"h"`
}

func TestLoad(t *testing.T) {
	_ = os.Setenv(EnvDirectory, "testdata")
	var err error
	var ts testStruct
	err = Load("test1", &ts)
	assert.NoError(t, err)
	assert.True(t, ts.Hello.World.Truth)
	ts = testStruct{}
	err = Load("test2", &ts)
	assert.NoError(t, err)
	assert.True(t, ts.Hello.World.Truth)
	ts = testStruct{}
	err = Load("test3", &ts)
	assert.Error(t, err)
	assert.IsType(t, &notFoundError{}, err)
	assert.True(t, IsNotFound(err))

	dir, _ := os.Getwd()
	dir2 := filepath.Join(dir, ".")
	assert.Equal(t, dir, dir2)
}
