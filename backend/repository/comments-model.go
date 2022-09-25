package repository

import (
	"fmt"
	"time"

	"github.com/MrJSdev/go-wp-vue/database"
	"github.com/MrJSdev/go-wp-vue/entity"
)

func FindCommentsByPostID(postID int) (comments []entity.Comment, err error) {
	db, err := database.GetDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	result, err := db.Query("SELECT comment_ID, comment_post_ID, comment_author, comment_date, comment_content, comment_parent, user_id FROM wp_comments WHERE comment_post_ID=?", postID)

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

func AddNewCommentByPostID(postID int, comment entity.Comment) (err error) {
	db, err := database.GetDB()

	if err != nil {
		return err
	}

	comment.CommentDate = time.Now().Format("2006-01-02 15:04:05")

	_, err = db.Exec("INSERT INTO wp_comments (comment_post_ID, comment_author, comment_author_email, comment_date, comment_content, comment_parent, user_id) VALUES (?, ?, ?, ?, ?, ?)", postID, comment.CommentAuthor, comment.CommentAuthorEmail, comment.CommentDate, comment.CommentContent, comment.CommentParent, comment.UserID)

	if err != nil {
		return err
	}

	return err
}
