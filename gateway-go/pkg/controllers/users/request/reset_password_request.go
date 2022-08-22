package request

type ResetPasswordRequest struct {
	Email string `json:"email" validator:"email;not_blank"`
}
