package core

import "context"

type Core interface {
	Run(ctx context.Context) error
	Stop(ctx context.Context) error
}
