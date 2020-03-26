package endpoint

import (
	"context"

	io "github.com/faith/Accounts2/accounts/pkg/io"
	service "github.com/faith/Accounts2/accounts/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateUserRequest collects the request parameters for the CreateUser method.
type CreateUserRequest struct {
	User io.User `json:"user"`
}

// CreateUserResponse collects the response parameters for the CreateUser method.
type CreateUserResponse struct {
	S1 string `json:"Id"`
	E0 error  `json:"Error"`
}

// MakeCreateUserEndpoint returns an endpoint that invokes CreateUser on the service.
func MakeCreateUserEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		s1, e0 := s.CreateUser(ctx, req.User)
		return CreateUserResponse{
			S1: s1,
			E0: e0,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateUserResponse) Failed() error {
	return r.E0
}

// GetUserRequest collects the request parameters for the GetUser method.
type GetUserRequest struct {
	ID string `json:"id"`
}

// GetUserResponse collects the response parameters for the GetUser method.
type GetUserResponse struct {
	S0 string `json:"Email"`
	E1 error  `json:"Error"`
}

// MakeGetUserEndpoint returns an endpoint that invokes GetUser on the service.
func MakeGetUserEndpoint(s service.AccountsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		s0, e1 := s.GetUser(ctx, req.ID)
		return GetUserResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateUser implements Service. Primarily useful in a client.
func (e Endpoints) CreateUser(ctx context.Context, user io.User) (s1 string, e0 error) {
	request := CreateUserRequest{User: user}
	response, err := e.CreateUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateUserResponse).S1, response.(CreateUserResponse).E0
}

// GetUser implements Service. Primarily useful in a client.
func (e Endpoints) GetUser(ctx context.Context, id string) (s0 string, e1 error) {
	request := GetUserRequest{ID: id}
	response, err := e.GetUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserResponse).S0, response.(GetUserResponse).E1
}
