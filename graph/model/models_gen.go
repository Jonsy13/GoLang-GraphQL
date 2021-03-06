// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Audio struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

type NewAudio struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type NewVideo struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	UserID string `json:"userId"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Video struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Author *User  `json:"author"`
}
