package service

import (
	"comment/domain/comment/entity"
	"comment/domain/comment/repository/facade"
	"comment/domain/comment/repository/persistence"
	"comment/infrastructure/common/context"
	baseEvent "comment/infrastructure/common/event"
)

type CommentDomainService struct {
	*context.Context
	CommentFactory
	facade.CommentRepositoryI
	baseEvent.PublisherI
}

func NewCommentDomainService(ctx *context.Context) CommentDomainService {
	return CommentDomainService{
		Context:            ctx,
		CommentFactory:     NewCommentFactory(),
		CommentRepositoryI: persistence.NewCommentRepository(),
		PublisherI:         baseEvent.NewBasePublisher(),
	}
}

func (a *CommentDomainService) Insert(comment entity.Comment) error {
	comment.Create()
	return a.CommentRepositoryI.Insert(a.CommentFactory.CreateCommentPO(comment))
}
