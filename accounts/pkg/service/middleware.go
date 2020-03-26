package service

import (
	"context"

	io "github.com/faith/Accounts2/accounts/pkg/io"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AccountsService) AccountsService

type loggingMiddleware struct {
	logger log.Logger
	next   AccountsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AccountsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AccountsService) AccountsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateUser(ctx context.Context, user io.User) (s1 string, e0 error) {
	defer func() {
		l.logger.Log("method", "CreateUser", "user", user, "s1", s1, "e0", e0)
	}()
	return l.next.CreateUser(ctx, user)
}
func (l loggingMiddleware) GetUser(ctx context.Context, id string) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "GetUser", "id", id, "s0", s0, "e1", e1)
	}()
	return l.next.GetUser(ctx, id)
}
