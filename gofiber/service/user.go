package service

type UserRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	CustomerID int    `json:"customer_id"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type UserService interface {
	Register(UserRequest) (*UserResponse, error)
	Login(string, string) (*UserResponse, error)
}
