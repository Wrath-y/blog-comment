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

func (*CommentFactory) CreateComment(comment po.Comment) entity.Comment {
	return entity.Comment{
		Pid:        comment.Pid,
		ArticleId:  comment.ArticleId,
		Name:       comment.Name,
		Email:      comment.Email,
		Url:        comment.Url,
		Content:    comment.Content,
		CreateTime: comment.CreateTime,
	}
}
