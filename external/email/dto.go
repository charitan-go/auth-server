package email

type SendRegisterAccountEmailRequestDto struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}
