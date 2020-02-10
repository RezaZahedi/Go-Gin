package database

var books = []Book{
	{
		ID: ID{
			BackField: 1,
		},
		Name:        "The Perks of Being a Wallflower",
		Description: "Des1",
	},
	{
		ID: ID{
			BackField: 2,
		},
		Name:        "The Man Without Qualities",
		Description: "Des2",
	},
	{
		ID: ID{
			BackField: 3,
		},
		Name:        "Where the Wild Things Are",
		Description: "Des3",
	},
}

var users = []User{
	{
		ID{
			BackField: 1,
		},
		"reza",
		"rezareza",
	},
}
