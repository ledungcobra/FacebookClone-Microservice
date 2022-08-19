package request

type ActivateAccountRequest struct {
	Email string `json:"email" validator:"email"`
	Token string `json:"token" validator:"not_blank"`
}
