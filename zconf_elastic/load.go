package zconf_elastic

import (
	"github.com/olivere/elastic/v7"
	"github.com/zionkit/zion/zconf"
)

func New(key string) (client *elastic.Client, err error) {
	var opts Options
	if err = zconf.Load(key, &opts); err != nil {
		return
	}
	client, err = elastic.NewClient(opts.Unwrap()...)
	return
}
