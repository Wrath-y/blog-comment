package persistence

import (
	"comment/domain/comment/repository/facade"
	"comment/domain/comment/repository/po"
	"comment/infrastructure/util/db"
)

type CommentRepository struct{}

func NewCommentRepository() facade.CommentRepositoryI {
	return &CommentRepository{}
}

func (*CommentRepository) Insert(comment po.Comment) error {
	return db.Orm.Save(&comment).Error
}
