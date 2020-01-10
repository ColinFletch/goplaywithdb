package model

//Post to model posts to our REST API
type Post struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
