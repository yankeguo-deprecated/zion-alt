package svc_math

import (
	"context"
	"testing"
)

func TestService_Add(t *testing.T) {
	s := &Service{}
	s.Add(context.Background(), AddRequest{}, &AddResponse{})
}
