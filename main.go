package main

import (
	rails "github.com/ramin0/go-rails"
	_cfg "github.com/ramin0/go-rails-app/config"
)

func main() {
	rails.SetRouter(_cfg.Router)
	rails.Run()
}
