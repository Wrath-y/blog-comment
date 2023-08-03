package assembler

import (
	"comment/domain/comment/entity"
	"comment/interfaces/dto"
	"comment/interfaces/proto"
	"time"
)

func ToCommentEntity(comment dto.Comment) entity.Comment {
	return entity.Comment{
		Pid:       comment.Pid,
		ArticleId: comment.ArticleId,
		Name:      comment.Name,
		Email:     comment.Email,
		Url:       comment.Url,
		Content:   comment.Content,
	}
}

func ToCommentDTO(comment entity.Comment) dto.Comment {
	return dto.Comment{
		Pid:       comment.Pid,
		ArticleId: comment.ArticleId,
		Name:      comment.Name,
		Email:     comment.Email,
		Url:       comment.Url,
		Content:   comment.Content,
	}
}
func PbToCommentEntity(comment *proto.CommentBaseInfo) entity.Comment {
	return entity.Comment{
		Pid:        comment.Pid,
		ArticleId:  comment.ArticleId,
		Name:       comment.Name,
		Email:      comment.Email,
		Url:        comment.Url,
		Content:    comment.Content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}
