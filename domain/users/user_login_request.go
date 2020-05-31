package users

// LoginRequest is a struct to store user email and password for oauth api
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
