package main

var list []Article

type Article struct {
	UID     string `json:"uid"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Teaser  string `json:"teaser"`
	Content string `json:"content"`
}
