package service

type UserRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	CustomerID int8   `json:"customer_id"`
}

type UserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type UserService interface {
	Register(UserRequest) (*UserResponse, error)
	Login(string, string) (*UserResponse, error)
}
