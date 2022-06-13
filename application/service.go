package application

import (
	"github.com/kcalixto/penny-api/domain/contract"
	"github.com/kcalixto/penny-api/domain/service"
)

type AppService struct {
	app App
}

func (svc AppService) Instance() contract.InstanceService {
	return service.NewInstanceService()
}
