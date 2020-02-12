package REST_api

import (
	"github.com/RezaZahedi/Go-Gin/fibonacci"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FibonacciAPI struct {
	FibonacciService fibonacci.FibonacciService
}

func ProvideFibonacciAPI(f fibonacci.FibonacciService) *FibonacciAPI {
	return &FibonacciAPI{FibonacciService: f}
}

func (*FibonacciAPI) ShowGetFibonacciNumberPage(c *gin.Context)  {
	render(c,
		gin.H{"title": "Get a New Fibonacci Number"},
		"get-fibonacci-number.html")
}

func (f *FibonacciAPI) GetFibonacciAnswer (c *gin.Context) {
	number, err := strconv.Atoi(c.PostForm("number"))
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	output, err := (*f.FibonacciService.FibonacciCalculator)(number)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	render(c,
		gin.H{"title": "Submission Successful",
			"payload": output},
		"fibonacci-answer.html")
}