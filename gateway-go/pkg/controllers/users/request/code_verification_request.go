package request

type CodeVerificationRequest struct {
	Code  string `json:"code" validator:"not_blank"`
	Email string `json:"email" validator:"email"`
}
