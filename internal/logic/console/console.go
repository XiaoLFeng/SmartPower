package console

import "SmartPower/internal/service"

type sConsole struct {
}

func init() {
	service.RegisterConsole(New())
}

func New() *sConsole {
	return &sConsole{}
}
