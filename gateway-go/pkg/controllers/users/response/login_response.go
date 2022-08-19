package response

type LoginResponse struct {
	Token     string `json:"token"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	ID        uint   `json:"id"`
	Picture   string `json:"picture"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
