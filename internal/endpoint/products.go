package endpoint

import (
	"lastlegends-gateway-service/internal/service"
)

type ProductEndpoints struct {
	Service service.IProductService
}

func MakeProductEndpoints(s service.IProductService) ProductEndpoints {
	return ProductEndpoints{
		Service: s,
	}
}
