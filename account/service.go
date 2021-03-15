package account

import "context"

// interface of the service
// create interface with methods that are to be exposed to tranport method
// interface then used to implement business logic
type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}

// to implement the transport, need to turn the methods into endpoints
// need to create structs that represent request and response for each of the methods
