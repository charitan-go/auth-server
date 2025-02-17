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

type RegisterCharityRequestDto struct {
	OrganizationName string `json:"organizationName"`
	TaxCode          string `json:"taxCode"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Address          string `json:"address"`
}

type RegisterResponseDto struct {
	Message string `json:"message"`
}

type LoginUserRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponseDto struct {
	Token string `json:"token"`
}

type GetMeResponseDto struct {
	ProfileReadableId string `json:"profileReadableIdId"`
	Email             string `json:"email"`
	Role              string `json:"role"`

	DonorDetails   *GetMeDonorDetailsResponseDto   `json:"donorDetails,omitempty"`
	CharityDetails *GetMeCharityDetailsResponseDto `json:"charityDetails,omitempty"`
}

type GetMeDonorDetailsResponseDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
}

type GetMeCharityDetailsResponseDto struct {
	OrganizationName string `json:"organizationName"`
	TaxCode          string `json:"taxCode"`
}

type MessageResponseDto struct {
	Message string `json:"message"`
}

type ErrorResponseDto struct {
	Message    string `json:"message"`
	StatusCode uint   `json:"statusCode"`
}
