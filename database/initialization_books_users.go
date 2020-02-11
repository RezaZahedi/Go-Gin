package database

import "github.com/RezaZahedi/Go-Gin/model"

var books = []model.Book{
	{
		ID: model.ID{
			BackField: 1,
		},
		Name:        "The Perks of Being a Wallflower",
		Description: "Des1",
	},
	{
		ID: model.ID{
			BackField: 2,
		},
		Name:        "The Man Without Qualities",
		Description: "Des2",
	},
	{
		ID: model.ID{
			BackField: 3,
		},
		Name:        "Where the Wild Things Are",
		Description: "Des3",
	},
}

var users = []model.User{
	{
		model.ID{
			BackField: 1,
		},
		"reza",
		"rezareza",
	},
	{
		model.ID{
			BackField: 2,
		},
		"user",
		"password",
	},
}
