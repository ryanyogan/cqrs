package db

import (
	"context"

	"github.com/ryanyogan/cqrs/schema"
)

// Repository defines the three actions one may perform
type Repository interface {
	Close()
	InsertShout(ctx context.Context, shout schema.Shout) error
	ListShouts(ctx context.Context, skip uint64, take uint64) ([]schema.Shout, error)
}

var impl Repository

// SetRepository takes a repostiroy and binds it to the impl ref of Repository
func SetRepository(repository Repository) {
	impl = repository
}

// Close will terminate the connection
func Close() {
	impl.Close()
}

// InsertShout takes a context and a shout schema
func InsertShout(ctx context.Context, shout schema.Shout) error {
	return impl.InsertShout(ctx, shout)
}

// ListShouts returns all of the shouts
func ListShouts(ctx context.Context, skip uint64, take uint64) ([]schema.Shout, error) {
	return impl.ListShouts(ctx, skip, take)
}
