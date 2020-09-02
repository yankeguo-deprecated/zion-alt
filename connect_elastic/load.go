package connect_elastic

import (
	"github.com/olivere/elastic/v7"
	"github.com/zionkit/zion/conf"
)

func New(key string) (client *elastic.Client, err error) {
	var opts Options
	if err = conf.Load(key, &opts); err != nil {
		return
	}
	client, err = elastic.NewClient(opts.Unwrap()...)
	return
}
