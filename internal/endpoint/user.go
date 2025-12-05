package endpoint

import (
	"lastlegends-gateway-service/internal/service"
)

type UserEndpoints struct {
	Service service.IUserService
}

func MakeUserEndpoints(s service.IUserService) UserEndpoints {
	return UserEndpoints{
		Service: s,
	}
}
