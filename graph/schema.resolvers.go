package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/flaviossantana/go-graphql/graph/generated"
	"github.com/flaviossantana/go-graphql/graph/model"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category := model.Category{
		ID:          generateID(),
		Name:        input.Name,
		Description: &input.Description,
	}

	r.Categories = append(r.Categories, &category)

	return &category, nil

}

func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {

	course := model.Course{
		ID:          generateID(),
		Name:        input.Name,
		Description: &input.Description,
		Category:    findCategoryById(r, input.CategoryID),
	}

	r.Courses = append(r.Courses, &course)

	return &course, nil

}

func (r *mutationResolver) CreateChapter(ctx context.Context, input model.NewChapter) (*model.Chapter, error) {

	chapter := model.Chapter{
		ID:     generateID(),
		Name:   input.Name,
		Course: findCourseById(r, input.CourseID),
	}

	r.Chapters = append(r.Chapters, &chapter)

	return &chapter, nil

}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.Resolver.Categories, nil
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return r.Resolver.Courses, nil
}

func (r *queryResolver) Chapter(ctx context.Context) ([]*model.Chapter, error) {
	return r.Resolver.Chapters, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func generateID() string {
	return fmt.Sprintf("T%d", rand.Int())
}

func findCategoryById(r *mutationResolver, categoryID string) *model.Category {
	var category *model.Category
	for _, findCategory := range r.Categories {
		if findCategory.ID == categoryID {
			category = findCategory
		}
	}
	return category
}

func findCourseById(r *mutationResolver, categoryID string) *model.Course {
	var course *model.Course
	for _, findCourse := range r.Courses {
		if findCourse.ID == categoryID {
			course = findCourse
		}
	}
	return course
}
