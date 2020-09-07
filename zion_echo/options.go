package zion_echo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/random"
	"github.com/zionkit/zion/zconf_stdlib/zconf_time"
	"time"
)

type Options struct {
	Debug     bool `json:"debug"`
	Port      int  `json:"port"`
	Gzip      int  `json:"gzip"`
	BasicAuth struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"basic_auth"`
	Static struct {
		Dirs  map[string]string `json:"dirs"`
		Files map[string]string `json:"files"`
	} `json:"static"`
	CORS struct {
		AllowOrigins     []string            `json:"allow_origins"`
		AllowMethods     []string            `json:"allow_methods"`
		AllowHeaders     []string            `json:"allow_headers"`
		AllowCredentials bool                `json:"allow_credentials"`
		ExposeHeaders    []string            `json:"expose_headers"`
		MaxAge           zconf_time.Duration `json:"max_age"`
	} `json:"cors"`
	BodyLimit string            `json:"body_limit"`
	Rewrites  map[string]string `json:"rewrites"`
}

func (opts Options) Create() (e *echo.Echo) {
	e = echo.New()
	e.Debug = opts.Debug
	e.HideBanner = true
	e.HidePort = true
	e.Server.Addr = fmt.Sprintf(":%d", opts.Port)
	e.Use(middleware.Recover())
	if opts.BasicAuth.Username != "" && opts.BasicAuth.Password != "" {
		e.Use(middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
			return opts.BasicAuth.Username == username && opts.BasicAuth.Password == password, nil
		}))
	}
	if opts.Gzip != 0 {
		e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: opts.Gzip}))
	}
	if len(opts.CORS.AllowOrigins) != 0 {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     opts.CORS.AllowOrigins,
			AllowMethods:     opts.CORS.AllowMethods,
			AllowHeaders:     opts.CORS.AllowHeaders,
			AllowCredentials: opts.CORS.AllowCredentials,
			ExposeHeaders:    opts.CORS.ExposeHeaders,
			MaxAge:           int(opts.CORS.MaxAge.Unwrap() / time.Second),
		}))
	}
	for k, v := range opts.Static.Dirs {
		e.Static(k, v)
	}
	for k, v := range opts.Static.Files {
		e.File(k, v)
	}
	if opts.BodyLimit != "" {
		e.Use(middleware.BodyLimit(opts.BodyLimit))
	}
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{Generator: func() string {
		return random.String(16, random.Alphanumeric)
	}}))
	if len(opts.Rewrites) != 0 {
		e.Pre(middleware.Rewrite(opts.Rewrites))
	}
	return
}
