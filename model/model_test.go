package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBook(t *testing.T) {
	book := Book{
		ID:          ID{BackField: 2},
		Name:        "testName",
		Description: "testDesc",
	}
	//t.Logf("%v", book)
	assert.Equal(t, "{{2} testName testDesc}", fmt.Sprintf("%v", book))
}
func TestUser(t *testing.T) {
	user := User{
		ID:       ID{BackField: 2},
		Username: "Rez",
		Password: "RezRez",
	}
	assert.Equal(t, "{{2} Rez RezRez}", fmt.Sprintf("%v", user))
}
