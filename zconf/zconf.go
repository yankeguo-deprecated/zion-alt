package zconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Loader convert various configuration formats to map[string]interface{}
type Loader interface {
	Name() string
	Extensions() []string
	Load(filename string, buf []byte, out *map[string]interface{}) (err error)
}

const (
	EnvDirectory = "ZCONF_DIR"
)

var (
	DefaultDirectories = []string{".", "conf", "config"}
)

var (
	loaders = []Loader{
		&jsonLoader{},
		&yamlLoader{},
	}
)

// notFoundError error emitted when nothing is found in conf directory
type notFoundError struct {
	key  string
	dirs []string
}

func (n *notFoundError) Error() string {
	return fmt.Sprintf("zconf: conf with key \"%s\" not found in dirs \"%s\"", n.key, strings.Join(n.dirs, ","))
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	if _, ok := err.(*notFoundError); ok {
		return true
	}
	return false
}

func Load(key string, out interface{}) (err error) {
	dirs := DefaultDirectories
	if dir := strings.TrimSpace(os.Getenv(EnvDirectory)); dir != "" {
		dirs = append([]string{dir}, dirs...)
	}
	var buf []byte
	for _, dir := range dirs {
		for _, loader := range loaders {
			for _, ext := range loader.Extensions() {
				filename := filepath.Join(dir, key+ext)
				if buf, err = ioutil.ReadFile(filename); err != nil {
					if os.IsNotExist(err) {
						err = nil
						continue
					} else {
						return
					}
				}
				// unmarshal to general map
				m := map[string]interface{}{}
				if err = loader.Load(filename, buf, &m); err != nil {
					return
				}
				// re-marshal to json
				if buf, err = json.Marshal(m); err != nil {
					return
				}
				// unmarshal from json
				if err = json.Unmarshal(buf, out); err != nil {
					return
				}
				return
			}
		}
	}
	err = &notFoundError{dirs: dirs, key: key}
	return
}
