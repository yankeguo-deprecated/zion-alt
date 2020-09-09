package zion

import "context"

type Service interface {
	Run(ctx context.Context) (err error)
}
