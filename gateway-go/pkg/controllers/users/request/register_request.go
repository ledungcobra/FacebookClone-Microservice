package request

type RegisterRequest struct {
	FirstName  string `json:"first_name" validate:"not_blank"`
	LastName   string `json:"last_name" validator:"not_blank"`
	Email      string `json:"email" validator:"email"`
	Password   string `json:"password" validator:"not_blank"`
	BirthYear  int    `json:"birth_year" validator:"range:1900-9999"`
	BirthMonth int    `json:"birth_month" validator:"range:1-12"`
	BirthDay   int    `json:"birth_day" validator:"range:1-31"`
	Gender     string `json:"gender" validator:"in:male,female"`
}
