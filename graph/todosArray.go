package graph

import (
	"pritodo/graph/model"
)

var todosArray = []*model.Todo{
	{
		ID:   "1",
		Text: "PriTestTodo1true",
		Done: true,
	},
	{
		ID:   "2",
		Text: "PriTestTodo2true",
		Done: true,
	},
	{
		ID:   "3",
		Text: "PriTestTodo1false",
		Done: false,
	},
	{
		ID:   "4",
		Text: "PriTestTodo2false",
		Done: false,
	},
	{
		ID:   "5",
		Text: "testRefresh",
		Done: false,
	},
}
