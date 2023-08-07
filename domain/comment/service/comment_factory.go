package service

import (
	"comment/domain/comment/entity"
	"comment/domain/comment/repository/po"
)

type CommentFactory struct {
}

func NewCommentFactory() CommentFactory {
	return CommentFactory{}
}

func (*CommentFactory) CreateCommentPO(comment entity.Comment) po.Comment {
	return po.Comment{
		Pid:        comment.Pid,
		ArticleId:  comment.ArticleId,
		Name:       comment.Name,
		Email:      comment.Email,
		Url:        comment.Url,
		Content:    comment.Content,
		CreateTime: comment.CreateTime,
		UpdateTime: comment.UpdateTime,
	}
}

func (*CommentFactory) CreateComments(comments []po.Comment) []entity.Comment {
	list := make([]entity.Comment, 0, len(comments))
	for _, v := range comments {
		list = append(list, entity.Comment{
			Id:         v.Id,
			Pid:        v.Pid,
			ArticleId:  v.ArticleId,
			Name:       v.Name,
			Email:      v.Email,
			Url:        v.Url,
			Content:    v.Content,
			CreateTime: v.CreateTime,
		})
	}

	return list
}
