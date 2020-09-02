package zconf_elastic

import (
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/zionkit/zion/zconf_time"
)

const (
	BackoffZero        = "zero"
	BackoffStop        = "stop"
	BackoffConstant    = "constant"
	BackoffExponential = "exponential"
)

type OptionsSrc struct {
	URLs      []string `json:"urls"`
	BasicAuth struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"basic_auth"`
	Gzip    bool `json:"gzip"`
	Sniffer struct {
		Enabled        bool                `json:"enabled"`
		Interval       zconf_time.Duration `json:"interval"`
		Timeout        zconf_time.Duration `json:"timeout"`
		TimeoutStartup zconf_time.Duration `json:"timeout_startup"`
	} `json:"sniffer"`
	HealthCheck struct {
		Enabled        bool                `json:"enabled"`
		Interval       zconf_time.Duration `json:"interval"`
		Timeout        zconf_time.Duration `json:"timeout"`
		TimeoutStartup zconf_time.Duration `json:"timeout_startup"`
	} `json:"health_check"`
	Retrier struct {
		Backoff    string              `json:"backoff"`
		Timeout    zconf_time.Duration `json:"timeout"`
		TimeoutMin zconf_time.Duration `json:"timeout_min"`
		TimeoutMax zconf_time.Duration `json:"timeout_max"`
	} `json:"retrier"`
}

func (opts OptionsSrc) Build() (ret []elastic.ClientOptionFunc) {
	ret = []elastic.ClientOptionFunc{
		elastic.SetURL(opts.URLs...),
		elastic.SetGzip(opts.Gzip),
		elastic.SetSniff(opts.Sniffer.Enabled),
		elastic.SetHealthcheck(opts.HealthCheck.Enabled),
	}
	if opts.BasicAuth.Username != "" && opts.BasicAuth.Password != "" {
		ret = append(ret, elastic.SetBasicAuth(opts.BasicAuth.Username, opts.BasicAuth.Password))
	}
	if opts.Sniffer.Interval != 0 {
		ret = append(ret, elastic.SetSnifferInterval(opts.Sniffer.Interval.Unwrap()))
	}
	if opts.Sniffer.Timeout != 0 {
		ret = append(ret, elastic.SetSnifferTimeout(opts.Sniffer.Timeout.Unwrap()))
	}
	if opts.Sniffer.TimeoutStartup != 0 {
		ret = append(ret, elastic.SetSnifferTimeoutStartup(opts.Sniffer.TimeoutStartup.Unwrap()))
	}
	if opts.HealthCheck.Interval != 0 {
		ret = append(ret, elastic.SetHealthcheckInterval(opts.HealthCheck.Interval.Unwrap()))
	}
	if opts.HealthCheck.Timeout != 0 {
		ret = append(ret, elastic.SetHealthcheckTimeout(opts.HealthCheck.Timeout.Unwrap()))
	}
	if opts.HealthCheck.TimeoutStartup != 0 {
		ret = append(ret, elastic.SetHealthcheckTimeoutStartup(opts.HealthCheck.TimeoutStartup.Unwrap()))
	}

	switch opts.Retrier.Backoff {
	case BackoffZero:
		ret = append(ret, elastic.SetRetrier(elastic.NewBackoffRetrier(elastic.ZeroBackoff{})))
	case BackoffStop:
		ret = append(ret, elastic.SetRetrier(elastic.NewBackoffRetrier(elastic.StopBackoff{})))
	case BackoffConstant:
		if opts.Retrier.Timeout != 0 {
			ret = append(ret, elastic.SetRetrier(elastic.NewBackoffRetrier(elastic.NewConstantBackoff(opts.Retrier.Timeout.Unwrap()))))
		}
	case BackoffExponential:
		if opts.Retrier.TimeoutMin != 0 && opts.Retrier.TimeoutMax != 0 {
			ret = append(ret, elastic.SetRetrier(elastic.NewBackoffRetrier(elastic.NewExponentialBackoff(opts.Retrier.TimeoutMin.Unwrap(), opts.Retrier.TimeoutMax.Unwrap()))))
		}
	}

	return
}

type Options struct {
	optionFuncs []elastic.ClientOptionFunc
}

func (c Options) Unwrap() []elastic.ClientOptionFunc {
	return c.optionFuncs
}

func (c *Options) UnmarshalJSON(bytes []byte) (err error) {
	var opts OptionsSrc
	if err = json.Unmarshal(bytes, &opts); err != nil {
		return
	}
	c.optionFuncs = opts.Build()
	return
}
