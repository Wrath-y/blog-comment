package service

import (
	"comment/domain/comment/entity"
	"comment/domain/comment/service"
	"comment/infrastructure/common/context"
)

type CommentApplicationService struct {
	*context.Context
	service.CommentDomainService
}

func NewCommentApplicationService(ctx *context.Context) *CommentApplicationService {
	return &CommentApplicationService{
		ctx,
		service.NewCommentDomainService(ctx),
	}
}

func (a *CommentApplicationService) Insert(comment entity.Comment) error {
	return a.CommentDomainService.Insert(comment)
}
