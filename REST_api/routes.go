package REST_api

import (
	"github.com/RezaZahedi/Go-Gin/fibonacci"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, p *ProductAPI, u *UserAPI) error  {
	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// Handle the index route
	router.GET("/", p.ShowIndexPage)

	// Group user related routes together
	userRoutes := router.Group("/u")
	{
		// Handle the GET requests at /u/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", ensureNotLoggedIn(), u.ShowLoginPage)

		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", ensureNotLoggedIn(), u.PerformLogin)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", ensureLoggedIn(), u.Logout)

		// Handle the GET requests at /u/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", ensureNotLoggedIn(), u.ShowRegistrationPage)

		// Handle POST requests at /u/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", ensureNotLoggedIn(), u.Register)
	}

	// Group article related routes together
	articleRoutes := router.Group("/book")
	{
		// Handle GET requests at /article/view/some_article_id
		articleRoutes.GET("/view/:book_id", p.GetBook)

		// Handle the GET requests at /article/create
		// Show the article creation page
		// Ensure that the user is logged in by using the middleware
		articleRoutes.GET("/create", ensureLoggedIn(), p.ShowBookCreatingPage)

		// Handle POST requests at /article/create
		// Ensure that the user is logged in by using the middleware
		articleRoutes.POST("/create", ensureLoggedIn(), p.CreateBook)
	}

	ff := func(a int) int { return a * 2 }
	FibonacciAPI := ProvideFibonacciAPI(fibonacci.ProvideFibonacciService(ff))
	// Group Fibonacci related routes together
	fiboRoutes := router.Group("/fibo")
	{
		//TODO: add routes
		fiboRoutes.GET("/", FibonacciAPI.ShowGetFibonacciNumberPage)

		fiboRoutes.POST("/", FibonacciAPI.GetFibonacciAnswer)
	}
	return nil
}
