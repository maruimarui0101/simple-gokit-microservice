package account

import "context"

// representing a user within the business logic
// need to be able to convert to/from JSON if serving via a HTTP API
type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// creating an interface to be able to interact with a database
type Respository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}
