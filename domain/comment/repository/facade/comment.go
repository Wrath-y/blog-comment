package facade

import (
	"comment/domain/comment/repository/po"
)

type CommentRepositoryI interface {
	FindByArticleId(articleId int64, limit int8) ([]po.Comment, error)
	FindByArticleIdLastId(articleId, id int64, limit int8) ([]po.Comment, error)
	GetCountByArticleId(articleId int64) (int64, error)
	Insert(po.Comment) error
}
