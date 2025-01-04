package middleware

import (
	"github.com/amankumarsingh77/cmr/config"
	"github.com/amankumarsingh77/cmr/internal/auth"
	"github.com/amankumarsingh77/cmr/pkg/logger"
)

type MiddlewareManager struct {
	authUC  auth.UseCase
	cfg     *config.Config
	origins []string
	logger  logger.Logger
}

// Middleware manager constructor
func NewMiddlewareManager(authUC auth.UseCase, cfg *config.Config, origins []string, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{authUC: authUC, cfg: cfg, origins: origins, logger: logger}
}
