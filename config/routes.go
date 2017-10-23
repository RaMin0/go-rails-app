package config

import (
	rails "github.com/ramin0/go-rails"
	_c "github.com/ramin0/go-rails-app/app/controllers"
)

var Router = rails.NewRouter().
	Resources("posts", _c.PostsController)
