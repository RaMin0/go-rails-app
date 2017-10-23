package models

import (
	"fmt"

	rails "github.com/ramin0/go-rails"
)

type Post struct {
	rails.Model
}

func (c *Post) Title() string {
	return fmt.Sprintf("Title #%s", c.ID)
}

func (c *Post) Content() string {
	return "Lorem ipsum dolor sit amet, consectetur adipiscing elit. " +
		"Sed vel porttitor tortor. Nullam mi enim, dapibus at cursus eu, " +
		"condimentum nec ex. Fusce volutpat, magna id venenatis pharetra, " +
		"eros est euismod libero, quis porttitor metus libero ut elit. Donec " +
		"convallis leo dolor, id condimentum metus tempus id. Proin purus nisi, " +
		"lobortis ac lacus et, volutpat eleifend neque."
}
