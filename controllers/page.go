package controllers

import (
	"base/views"
)

func NewPage() *Page {
	return &Page{
		HomeView: views.NewView("main", "views/pages/home.gohtml"),
		ContactView: views.NewView("main", "views/pages/contact.gohtml"),
	}
}

type Page struct {
	HomeView *views.View
	ContactView *views.View
}
