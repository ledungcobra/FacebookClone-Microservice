package response

type RegisterResponse struct {
	Success  bool   `json:"success"`
	UserName string `json:"user_name"`
	UserID   uint   `json:"user_id"`
	Token    string `json:"token"`
}
