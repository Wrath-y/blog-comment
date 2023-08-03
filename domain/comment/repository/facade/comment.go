package facade

import "comment/domain/comment/repository/po"

type CommentRepositoryI interface {
	Insert(po.Comment) error
}
