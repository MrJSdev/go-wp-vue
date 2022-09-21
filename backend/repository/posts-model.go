package repository

import (
	"fmt"

	"github.com/MrJSdev/go-wp-vue/database"
	"github.com/MrJSdev/go-wp-vue/entity"
)

func FindPostByID(id string) (post *entity.Post, err error) {
	db, err := database.GetDB()

	if err != nil {
		return nil, err
	}

	result := db.QueryRow("SELECT id,post_title,post_content,post_date,post_author,post_status,comment_status,comment_count,post_name,post_modified,post_parent,guid ,post_type FROM wp_posts WHERE id = ?", id)

	post = &entity.Post{}
	err = result.Scan(&post.ID, &post.Title, &post.Content, &post.Date, &post.Author, &post.Status, &post.CommentStatus, &post.CommCount, &post.Name, &post.Modified, &post.Parent, &post.Guid, &post.Type)

	if err != nil {
		return nil, err
	}

	return post, err
}

func FindAllPosts() (posts []entity.Post, err error) {
	db, err := database.GetDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	result, err := db.Query("SELECT id,post_title,post_content,post_date,post_author,post_status,comment_status,comment_count,post_name,post_modified,post_parent,guid ,post_type FROM wp_posts  LIMIT ?", 10)

	if err != nil {
		return nil, err
	}

	var post entity.Post

	for result.Next() {
		err = result.Scan(&post.ID, &post.Title, &post.Content, &post.Date, &post.Author, &post.Status, &post.CommentStatus, &post.CommCount, &post.Name, &post.Modified, &post.Parent, &post.Guid, &post.Type)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, err
}
