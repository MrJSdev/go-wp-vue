package entity

type Comment struct {
	CommentID          int       `json:"comment_id"`
	CommentPostID      int       `json:"comment_post_id"`
	CommentAuthor      string    `json:"comment_author"`
	CommentAuthorEmail string    `json:"comment_author_email"`
	CommentAuthorURL   string    `json:"comment_author_url"`
	CommentAuthorIP    string    `json:"comment_author_ip"`
	CommentDate        string    `json:"comment_date"`
	CommentDateGmt     string    `json:"comment_date_gmt"`
	CommentContent     string    `json:"comment_content"`
	CommentKarma       int       `json:"comment_karma"`
	CommentAgent       string    `json:"comment_agent"`
	CommentType        string    `json:"comment_type"`
	CommentParent      int       `json:"comment_parent"`
	UserID             int       `json:"user_id"`
	Children           []Comment `json:"children"`
}
