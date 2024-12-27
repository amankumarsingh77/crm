package auth

import "github.com/labstack/echo/v4"

type Handler interface {
	// Auth
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	ResetPassword() echo.HandlerFunc

	// Authenticated Users
	Logout() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetMe() echo.HandlerFunc
	GetAllApplicationStatus() echo.HandlerFunc
	GetApplicationStatus() echo.HandlerFunc
	GetAllDocuments() echo.HandlerFunc
	UploadDocument() echo.HandlerFunc
	GetNotification() echo.HandlerFunc

	//Admin Only
	GetAllUsers() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeactivateUser() echo.HandlerFunc
	GetAllApplications() echo.HandlerFunc
	GetApplication() echo.HandlerFunc
	UpdateApplication() echo.HandlerFunc
	UpdateBulkApplication() echo.HandlerFunc
	GetDocuments() echo.HandlerFunc
	VerifyDocument() echo.HandlerFunc
	RemoveDocument() echo.HandlerFunc
	GetAnalytics() echo.HandlerFunc
	GetApplicationReport() echo.HandlerFunc
	GetUserActivityReport() echo.HandlerFunc
}
