package graph

import "github.com/Jonsy13/GoLang-GraphQL/graph/model"

//go:generate go run github.com/99designs/gqlgen

//Resolver is used to resolve.
type Resolver struct {
	videos []*model.Video
}
