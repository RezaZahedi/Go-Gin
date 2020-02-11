package REST_api

import (
	"github.com/RezaZahedi/Go-Gin/user"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

type UserAPI struct {
	UserService user.UserService
}

func ProvideUserAPI(p user.UserService) *UserAPI {
	return &UserAPI{UserService: p}
}

func (*UserAPI) ShowLoginPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func (u *UserAPI) PerformLogin(c *gin.Context) {
	// Obtain the POSTed username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Check if the username/password combination is valid
	isValid, err := u.UserService.IsUserValid(username, password)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !isValid {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
		return
	}
	// If the username/password is valid set the token in a cookie
	token := generateSessionToken()
	c.SetCookie("token", token, 3600, "", "", false, true)
	c.Set("is_logged_in", true)

	render(c,
		gin.H{"title": "Successful Login"},
		"login-successful.html")

}

func generateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

func (*UserAPI) Logout(c *gin.Context) {
	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (*UserAPI) ShowRegistrationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Register"}, "register.html")
}

func (u *UserAPI) Register(c *gin.Context) {
	// Obtain the POSTed username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

	_, err := u.UserService.RegisterNewUser(username, password)
	if err != nil {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest,
			"register.html",
			gin.H{"ErrorTitle":   "Registration Failed",
				"ErrorMessage": err.Error()})
		return
	}
	// If the user is created, set the token in a cookie and log the user in
	token := generateSessionToken()
	c.SetCookie("token", token, 3600, "", "", false, true)
	c.Set("is_logged_in", true)

	render(c,
		gin.H{"title": "Successful registration & Login"},
		"login-successful.html")
}