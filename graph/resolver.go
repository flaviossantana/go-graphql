package graph

import "github.com/flaviossantana/go-graphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Categories []*model.Category
	Chapters   []*model.Chapter
	Courses    []*model.Course
}
