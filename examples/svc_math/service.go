package svc_math

import (
	"context"
	"github.com/go-redis/redis/v7"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
	"reflect"
)

type Service struct {
	redis *redis.Client       `zion:""`
	es    *elastic.Client     `zion:""`
	rw    http.ResponseWriter `zion:""`
	req   *http.Request       `zion:""`
}

type AddRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type AddResponse struct {
	Value int `json:"value"`
}

func (a *Service) Add(ctx context.Context, req AddRequest, res *AddResponse) (err error) {
	typ := reflect.TypeOf(a)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		t, ok := f.Tag.Lookup("zion")
		log.Println(t, ok)
	}
	return
}
