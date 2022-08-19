package response

type RegisterResponse struct {
	Success   bool   `json:"success"`
	UserName  string `json:"user_name"`
	ID        uint   `json:"id"`
	Token     string `json:"token"`
	Picture   string `json:"picture"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Verified  bool   `json:"verified"`
}
