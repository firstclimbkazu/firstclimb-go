package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"firstclimb-go/entity"
	"firstclimb-go/graph/generated"
	"firstclimb-go/model"
	"fmt"
	"strconv"
)

func (r *articleResolver) TextItems(ctx context.Context, obj *model.Article) ([]*model.TextItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *articleResolver) MovieItems(ctx context.Context, obj *model.Article) ([]*model.MovieItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *movieItemResolver) Order(ctx context.Context, obj *model.MovieItem) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserData) (*model.User, error) {
	record := entity.User{
		Name:     input.Name,
		Password: input.Password,
		Email:    input.Email,
	}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	res := model.NewUserFromEntity(&record)

	return res, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UserData) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.PostData) (*model.Post, error) {
	record := entity.Post{}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	res := model.NewPostFromEntity(&record)

	return res, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, input model.PostData) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.ArticleData) (*model.Article, error) {
	postId, err := strconv.Atoi(input.PostID)
	if err != nil {
		return nil, err
	}
	record := entity.Article{
		Title:  input.Title,
		PostID: fmt.Sprintf("%d", postId),
	}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	res := model.NewArticleFromEntity(&record)

	return res, nil
}

func (r *mutationResolver) UpdateArticle(ctx context.Context, input model.ArticleData) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTextItem(ctx context.Context, input model.TextItemData) (*model.TextItem, error) {
	articleId, err := strconv.Atoi(input.ArticleID)
	if err != nil {
		return nil, err
	}
	order, err := strconv.Atoi(input.Order)
	if err != nil {
		return nil, err
	}
	record := entity.TextItem{
		Value:     input.Value,
		Order:     uint(order),
		ArticleID: uint(articleId),
	}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	res := model.NewTextItemFromEntity(&record)

	return res, nil
}

func (r *mutationResolver) UpdateTextItem(ctx context.Context, input model.TextItemData) (*model.TextItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMovieItem(ctx context.Context, input model.MovieItemData) (*model.MovieItem, error) {
	articleId, err := strconv.Atoi(input.ArticleID)
	if err != nil {
		return nil, err
	}
	order, err := strconv.Atoi(input.Order)
	if err != nil {
		return nil, err
	}
	record := entity.MovieItem{
		Value:     input.Value,
		Order:     uint(order),
		ArticleID: uint(articleId),
	}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	res := model.NewMovieItemFromEntity(&record)

	return res, nil
}

func (r *mutationResolver) UpdateMovieItem(ctx context.Context, input model.MovieItemData) (*model.MovieItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) Articles(ctx context.Context, obj *model.Post) ([]*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	var records []entity.Post
	if err := r.DB.Find(&records).Error; err != nil {
		return nil, err
	}
	posts := []*model.Post{}
	for _, record := range records {
		posts = append(posts, model.NewPostFromEntity(&record))
	}

	return posts, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *textItemResolver) Order(ctx context.Context, obj *model.TextItem) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Article returns generated.ArticleResolver implementation.
func (r *Resolver) Article() generated.ArticleResolver { return &articleResolver{r} }

// MovieItem returns generated.MovieItemResolver implementation.
func (r *Resolver) MovieItem() generated.MovieItemResolver { return &movieItemResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// TextItem returns generated.TextItemResolver implementation.
func (r *Resolver) TextItem() generated.TextItemResolver { return &textItemResolver{r} }

type articleResolver struct{ *Resolver }
type movieItemResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type textItemResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *articleResolver) Post(ctx context.Context, obj *model.Article) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *articleResolver) TextItem(ctx context.Context, obj *model.Article) ([]*model.TextItem, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *articleResolver) MovieItem(ctx context.Context, obj *model.Article) ([]*model.MovieItem, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *postResolver) Article(ctx context.Context, obj *model.Post) ([]*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *movieItemResolver) Article(ctx context.Context, obj *model.MovieItem) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *textItemResolver) Article(ctx context.Context, obj *model.TextItem) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}
