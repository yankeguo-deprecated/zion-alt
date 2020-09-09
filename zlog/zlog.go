package zlog

import (
	"github.com/zionkit/zion/zion_elastic"
	"sync"
)

const (
	NONAME = "noname"
)

var (
	gOptions = Options{
		Hostname:  NONAME,
		Workload:  NONAME,
		Namespace: NONAME,
		Cluster:   NONAME,
	}
	gLock sync.Locker = &sync.Mutex{}
)

type Options struct {
	Hostname  string `json:"hostname"`
	Workload  string `json:"workload"`
	Namespace string `json:"namespace"`
	Cluster   string `json:"cluster"`
	Console   struct {
		Topics []string `json:"topics"`
	} `json:"console"`
	Elasticsearch struct {
		Topics  []string              `json:"topics"`
		Connect *zion_elastic.Options `json:"connect"`
	} `json:"elasticsearch"`
}
