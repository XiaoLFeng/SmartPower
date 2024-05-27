package token

import "SmartPower/internal/service"

type sToken struct {
}

func init() {
	service.RegisterToken(New())
}

func New() *sToken {
	return &sToken{}
}
