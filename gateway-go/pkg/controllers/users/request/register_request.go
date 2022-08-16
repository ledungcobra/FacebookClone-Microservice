package request

type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email" validator:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	BirthYear int `json:"birth_year"`
	BirthMonth int `json:"birth_month"`
	BirthDay int `json:"birth_day"`
	Gender string `json:"gender"`
}