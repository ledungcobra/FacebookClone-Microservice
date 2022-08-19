package request

type LoginRequest struct {
	Email    string `json:"email" validator:"email"`
	Password string `json:"password" validator:"not_blank"`
}
