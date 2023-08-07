package service

import (
	"comment/domain/comment/entity"
	"comment/domain/comment/service"
	"comment/infrastructure/common/context"
)

type CommentApplicationService struct {
	*context.Context
	commentDomainService service.CommentDomainService
}

func NewCommentApplicationService(ctx *context.Context) *CommentApplicationService {
	return &CommentApplicationService{
		ctx,
		service.NewCommentDomainService(ctx),
	}
}

func (c *CommentApplicationService) GetComments(articleId, id int64, limit int8) ([]entity.Comment, error) {
	return c.commentDomainService.GetComments(articleId, id, limit)
}

func (c *CommentApplicationService) GetCountByArticleId(articleId int64) (int64, error) {
	return c.commentDomainService.GetCountByArticleId(articleId)
}

func (c *CommentApplicationService) AddComment(comment entity.Comment) error {
	return c.commentDomainService.Insert(comment)
}
