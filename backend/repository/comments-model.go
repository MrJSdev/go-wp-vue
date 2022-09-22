package repository

import (
	"fmt"

	"github.com/MrJSdev/go-wp-vue/database"
	"github.com/MrJSdev/go-wp-vue/entity"
)

func FindCommentsByPostID(id string) (comments []entity.Comment, err error) {
	db, err := database.GetDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	result, err := db.Query("SELECT comment_ID, comment_post_ID, comment_author, comment_date, comment_content, comment_parent, user_id FROM wp_comments WHERE comment_post_ID=?", id)
	// result, err := db.Query("SELECT * FROM wp_comments WHERE comment_post_ID=?", id)

	if err != nil {
		return nil, err
	}

	var comment entity.Comment

	for result.Next() {
		err = result.Scan(&comment.CommentID, &comment.CommentPostID, &comment.CommentAuthor, &comment.CommentDate, &comment.CommentContent, &comment.CommentParent, &comment.UserID)

		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, err
}
