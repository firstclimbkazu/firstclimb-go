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
	"time"
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

	res := model.UserFromEntity(&record)

	return res, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UserData) (*model.User, error) {
	result := r.DB.Where("id = ?", input.UserID).Updates(&entity.User{Name: input.Name, Email: input.Email})
	if result.Error != nil {
		return nil, result.Error
	}
	record := entity.User{}
	r.DB.First(&record, input.UserID)
	res := model.UserFromEntity(&record)
	return res, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.PostData) (*model.Post, error) {
	record := entity.Post{}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	res := model.PostFromEntity(&record)

	return res, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, input model.PostData) (*model.Post, error) {
	t, _ := time.Parse(DATE_LAYOUT, *input.PostedAt)
	result := r.DB.Where("id = ?", input.PostID).Updates(&entity.Post{PostedAt: &t})
	if result.Error != nil {
		return nil, result.Error
	}
	record := entity.Post{}
	r.DB.First(&record, input.PostID)
	res := model.PostFromEntity(&record)
	return res, nil
}

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.ArticleData) (*model.Article, error) {
	postId, err := strconv.Atoi(input.PostID)
	if err != nil {
		return nil, err
	}
	record := entity.Article{
		Title:  input.Title,
		PostID: uint(postId),
	}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	res := model.ArticleFromEntity(&record)

	return res, nil
}

func (r *mutationResolver) UpdateArticle(ctx context.Context, input model.ArticleData) (*model.Article, error) {
	postId, err := strconv.Atoi(input.PostID)
	if err != nil {
		return nil, err
	}
	result := r.DB.Where("id = ?", input.ArticleID).Updates(&entity.Article{
		Title:  input.Title,
		Detail: input.Detail,
		PostID: uint(postId),
	})
	if result.Error != nil {
		return nil, result.Error
	}
	record := entity.Article{}
	r.DB.First(&record, input.ArticleID)
	res := model.ArticleFromEntity(&record)
	return res, nil
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

	res := model.TextItemFromEntity(&record)

	return res, nil
}

func (r *mutationResolver) UpdateTextItem(ctx context.Context, input model.TextItemData) (*model.TextItem, error) {
	articleId, err := strconv.Atoi(input.ArticleID)
	if err != nil {
		return nil, err
	}
	order, err := strconv.Atoi(input.Order)
	if err != nil {
		return nil, err
	}
	result := r.DB.Where("id = ?", input.TextItemID).Updates(&entity.TextItem{
		Order:     uint(order),
		Value:     input.Value,
		ArticleID: uint(articleId),
	})
	if result.Error != nil {
		return nil, result.Error
	}
	record := entity.TextItem{}
	r.DB.First(&record, input.ArticleID)
	res := model.TextItemFromEntity(&record)
	return res, nil
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

	res := model.MovieItemFromEntity(&record)

	return res, nil
}

func (r *mutationResolver) UpdateMovieItem(ctx context.Context, input model.MovieItemData) (*model.MovieItem, error) {
	articleId, err := strconv.Atoi(input.ArticleID)
	if err != nil {
		return nil, err
	}
	order, err := strconv.Atoi(input.Order)
	if err != nil {
		return nil, err
	}
	result := r.DB.Where("id = ?", input.MovieItemID).Updates(&entity.MovieItem{
		Order:     uint(order),
		Value:     input.Value,
		ArticleID: uint(articleId),
	})
	if result.Error != nil {
		return nil, result.Error
	}
	record := entity.MovieItem{}
	r.DB.First(&record, input.ArticleID)
	res := model.MovieItemFromEntity(&record)
	return res, nil
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
		posts = append(posts, model.PostFromEntity(&record))
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
const DATE_LAYOUT = "Mon Jan 2 15:04:05 MST 2006"

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
