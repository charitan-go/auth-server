package service

type JwtService interface {
}

type jwtServiceImpl struct {
}

func NewJwtService() JwtService {
	return &jwtServiceImpl{}
}
