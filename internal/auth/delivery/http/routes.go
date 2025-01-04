package http

import (
	"github.com/amankumarsingh77/cmr/config"
	"github.com/amankumarsingh77/cmr/internal/auth"
	"github.com/amankumarsingh77/cmr/internal/middleware"
	"github.com/labstack/echo/v4"
)

func MapAuthRoutes(authGroup *echo.Group, h auth.Handler, mw *middleware.MiddlewareManager, authUC auth.UseCase, cfg *config.Config) {
	authGroup.POST("/register", h.Register())
	authGroup.POST("/login", h.Login())
	authGroup.POST("/logout", h.Logout())
	authGroup.GET("/:user_id", h.GetUserByID(), mw.OwnerOrAdminMiddleware())
	authGroup.Use(mw.AuthJWTMiddleware(authUC, cfg))
	authGroup.GET("/me", h.GetMe())
	authGroup.PUT("/:user_id", h.Update(), mw.OwnerOrAdminMiddleware())
	authGroup.DELETE("/:user_id", h.DeactivateUser(), mw.RoleBasedAuthMiddleware([]string{"admin"}))
}
