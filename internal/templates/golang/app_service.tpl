package service

import (
	"go.uber.org/zap"
	"{{.pkg}}/internal/app/{{.name}}/repository"
)

type {{.className}}Service struct {
	logger *zap.Logger
	repo repository.I{{.className}}Repository
}

func New{{.className}}Service(logger *zap.Logger, repo repository.I{{.className}}Repository) *{{.className}}Service {
	return &{{.className}}Service{logger: logger, repo: repo}
}
