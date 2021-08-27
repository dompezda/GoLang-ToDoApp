package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"pritodo/graph/generated"
	"pritodo/graph/model"
	"strconv"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input string) ([]*model.Todo, error) {
	if input == "" {
		return nil, errors.New("zle dane")
	}

	var NewArray []*model.Todo
	var ArrayLength = len(todosArray) - 1
	NewIdAsString := todosArray[ArrayLength].ID
	i64, err := strconv.ParseInt(NewIdAsString, 10, 32)
	NewId := i64 + 1
	NewTodoId := strconv.Itoa(int(NewId))
	var newTodo = model.Todo{
		Text: input,
		ID:   NewTodoId,
		Done: false,
	}
	NewArray = append(todosArray, &newTodo)
	todosArray = NewArray

	return todosArray, err
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input string) ([]*model.Todo, error) {
	var newId = len(todosArray)
	var NewArray []*model.Todo

	for i := 0; i < newId; i++ {

		if todosArray[i].ID != input {
			NewArray = append(NewArray, todosArray[i])
		}
	}
	todosArray = NewArray
	return todosArray, nil
}

func (r *mutationResolver) MarkAsDone(ctx context.Context, input string) ([]*model.Todo, error) {
	var newId = len(todosArray)

	for i := 0; i < newId; i++ {

		if todosArray[i].ID == input {
			if todosArray[i].Done == true {
				todosArray[i].Done = false
			} else {
				todosArray[i].Done = true
			}

		}
	}

	return todosArray, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return todosArray, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func RemoveIndex(s []*model.Todo, index int) []*model.Todo {
	return append(s[:index], s[index+1:]...)
}
