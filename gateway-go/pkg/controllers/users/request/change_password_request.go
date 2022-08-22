package request

type ChangePasswordRequest struct {
	Password     string `json:"password" validator:"not_blank;length:8-50"`
	Email        string `json:"email" validator:"email"`
	Token        string `json:"token" validator:"not_blank"`
	ConfPassword string `json:"conf_password" validator:"not_blank;length:8-50"`
}
