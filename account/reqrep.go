package account

// each endpoint in go kit acts like an rpc
// all these structs need to be PUBLIC!

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"Password"`
	}

	CreateUserResponse struct {
		ok string `json:"ok"`
	}

	GetUserRequest struct {
		Email string `json:"email"`
	}

	GetUserResponse struct {
		Id string `json:"id"`
	}
)
