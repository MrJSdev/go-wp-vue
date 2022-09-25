package service

import (
	"github.com/MrJSdev/go-wp-vue/entity"
)

// store comments in a map
//
func GroupCommentsByParent(comments []entity.Comment) []entity.Comment {
	var groupedComments = []entity.Comment{}
	parentCompIDs := make(map[int]int)

	for index, comment := range comments {
		if comment.CommentParent != 0 {
			groupedComments[parentCompIDs[comment.CommentParent]].Children = append(
				groupedComments[parentCompIDs[comment.CommentParent]].Children, comment,
			)
		} else {
			parentCompIDs[comment.CommentID] = index
			groupedComments = append(groupedComments, comment)
		}
	}

	return groupedComments
}
