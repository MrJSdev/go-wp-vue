package entity

type Post struct {
	ID            int    `json:"id"`
	Title         string `json:"post_title"`
	Content       string `json:"post_content"`
	Date          string `json:"post_date"`
	Author        string `json:"post_author"`
	Status        string `json:"post_status"`
	CommentStatus string `json:"comment_status"`
	CommCount     int    `json:"comment_count"`
	Name          string `json:"post_name"`
	Modified      string `json:"post_modified"`
	Parent        int    `json:"post_parent"`
	Guid          string `json:"guid"`
	Type          string `json:"post_type"`
}
