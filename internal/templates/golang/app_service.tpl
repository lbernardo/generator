package service

import (
	"go.uber.org/zap"
)

type {{.className}}Service struct {
	logger *zap.Logger
}

func New{{.className}}Service(logger *zap.Logger) *{{.className}}Service {
	return &{{.className}}Service{logger: logger}
}
