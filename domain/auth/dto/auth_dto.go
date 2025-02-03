package dto

type RoleEnum string

const (
	RoleCharity RoleEnum = "CHARITY"
	RoleDonor   RoleEnum = "DONOR"
)

type RegisterDonorRequestDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Address   string `json:"address"`
}

type MessageResponseDto struct {
	Message string `json:"message"`
}

type ErrorResponseDto struct {
	Message    string `json:"message"`
	StatusCode uint   `json:"statusCode"`
}

type RegisterResponseDto struct {
	Message string `json:"message"`
}
