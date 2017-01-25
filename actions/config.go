package actions

import "github.com/gobuffalo/buffalo"

type ConfigResource struct {
	buffalo.Resource
}

// ConfigShow default implementation.
func ConfigShow(c buffalo.Context) error {
	return c.Render(200, r.HTML("config/show.html"))
}

// ConfigEdit default implementation.
func ConfigEdit(c buffalo.Context) error {
	return c.Render(200, r.HTML("config/edit.html"))
}

// ConfigCreate default implementation.
func ConfigCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("config/create.html"))
}
