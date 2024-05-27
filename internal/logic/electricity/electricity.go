package electricity

import "SmartPower/internal/service"

type sElectric struct {
}

func init() {
	service.RegisterElectric(New())
}

func New() *sElectric {
	return &sElectric{}
}
