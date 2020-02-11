package REST_api

import (
	"github.com/RezaZahedi/Go-Gin/user"
	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	UserService user.UserService
}

func ProvideUserAPI(p user.UserService) UserAPI {
	return UserAPI{UserService: p}
}

func (*UserAPI) ShowLoginPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

