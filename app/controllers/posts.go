package controllers

import (
	rails "github.com/ramin0/go-rails"
	"github.com/ramin0/go-rails-app/app/models"
)

var PostsController = rails.NewController().
	// index
	AddAction(rails.ActionIndex, func(ctx *rails.Context) {
		ctx.Var("posts", []*models.Post{
			&models.Post{rails.Model{ID: "1"}},
			&models.Post{rails.Model{ID: "2"}},
			&models.Post{rails.Model{ID: "3"}},
		})
	}).

	// show
	AddAction(rails.ActionShow, func(ctx *rails.Context) {
		ctx.Var("post", &models.Post{rails.Model{ID: ctx.Param("id")}})
	})
