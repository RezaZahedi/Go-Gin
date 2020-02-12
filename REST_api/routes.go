package REST_api

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, p *ProductAPI, u *UserAPI, f *FibonacciAPI) error {
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

	// Group book related routes together
	articleRoutes := router.Group("/book")
	{
		// Handle GET requests at /book/view/some_book_id
		articleRoutes.GET("/view/:book_id", p.GetBook)

		// Handle the GET requests at /book/create
		// Show the book creation page
		// Ensure that the user is logged in by using the middleware
		articleRoutes.GET("/create", ensureLoggedIn(), p.ShowBookCreatingPage)

		// Handle POST requests at /book/create
		// Ensure that the user is logged in by using the middleware
		articleRoutes.POST("/create", ensureLoggedIn(), p.CreateBook)
	}

	// Group Fibonacci related routes together
	fiboRoutes := router.Group("/fibo")
	{
		// Handle the GET requests at /fibo
		// Show the get fibonacci number page
		fiboRoutes.GET("/", f.ShowGetFibonacciNumberPage)

		// Handle POST requests at /fibo
		fiboRoutes.POST("/", f.GetFibonacciAnswer)
	}

	return nil
}
