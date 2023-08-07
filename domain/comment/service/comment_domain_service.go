package service

import (
	"comment/domain/comment/entity"
	"comment/domain/comment/repository/facade"
	"comment/domain/comment/repository/persistence"
	"comment/domain/comment/repository/po"
	"comment/infrastructure/common/context"
	"comment/infrastructure/common/errcode"
	baseEvent "comment/infrastructure/common/event"
	"fmt"
)

type CommentDomainService struct {
	*context.Context
	commentFactory     CommentFactory
	commentRepositoryI facade.CommentRepositoryI
	publisherI         baseEvent.PublisherI
}

func NewCommentDomainService(ctx *context.Context) CommentDomainService {
	return CommentDomainService{
		Context:            ctx,
		commentFactory:     NewCommentFactory(),
		commentRepositoryI: persistence.NewCommentRepository(),
		publisherI:         baseEvent.NewBasePublisher(),
	}
}

func (c *CommentDomainService) GetComments(articleId, id int64, limit int8) ([]entity.Comment, error) {
	if limit == 0 {
		limit = 15
	}
	var (
		list []po.Comment
		err  error
	)
	if id == 0 {
		list, err = c.commentRepositoryI.FindByArticleId(articleId, limit)
		if err != nil {
			c.Logger.ErrorL("获取评论列表失败", fmt.Sprintf("articleId: %d, id: %d", articleId, id), err.Error())
			return nil, err
		}
	} else {
		list, err = c.commentRepositoryI.FindByArticleIdLastId(articleId, id, limit)
		if err != nil {
			c.Logger.ErrorL("获取评论列表失败", fmt.Sprintf("articleId: %d, id: %d", articleId, id), err.Error())
			return nil, err
		}
	}

	return c.commentFactory.CreateComments(list), nil
}

func (c *CommentDomainService) GetCountByArticleId(articleId int64) (int64, error) {
	count, err := c.commentRepositoryI.GetCountByArticleId(articleId)
	if err != nil {
		c.Logger.ErrorL("获取评论数量失败", articleId, err.Error())
		return 0, err
	}

	return count, nil
}

func (c *CommentDomainService) Insert(comment entity.Comment) error {
	if comment.ArticleId == 0 {
		c.Logger.ErrorL("articleId is 0", comment, errcode.CommentArticleIdErr.Error())
		return errcode.CommentArticleIdErr
	}
	comment.Create()
	return c.commentRepositoryI.Insert(c.commentFactory.CreateCommentPO(comment))
}
