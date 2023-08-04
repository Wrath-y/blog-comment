package persistence

import (
	"comment/domain/comment/entity"
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

func (*CommentRepository) FindByArticleId(articleId, limit int) ([]*entity.Comment, error) {
	var comments []*entity.Comment
	return comments, db.Orm.Raw("select * from comment where article_id = ? order by id desc limit ?", articleId, limit).Find(&comments).Error
}

func (*CommentRepository) FindByArticleIdLastId(articleId, lastId, limit int) ([]*entity.Comment, error) {
	var comments []*entity.Comment
	return comments, db.Orm.Raw("select * from comment where id < ? and article_id = ? order by id desc limit ?", lastId, articleId, limit).Find(&comments).Error
}

func (*CommentRepository) GetCountByArticleId(articleId int) (int64, error) {
	var count int64
	return count, db.Orm.Raw("select COUNT(id) as comment_count from comment where article_id = ?", articleId).Count(&count).Error
}
