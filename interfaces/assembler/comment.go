package assembler

import (
	"comment/domain/comment/entity"
	"comment/interfaces/dto"
	"comment/interfaces/proto"
)

func ToEntity(comment dto.Comment) entity.Comment {
	return entity.Comment{
		Pid:       comment.Pid,
		ArticleId: comment.ArticleId,
		Name:      comment.Name,
		Email:     comment.Email,
		Url:       comment.Url,
		Content:   comment.Content,
	}
}

func ToDTO(comment entity.Comment) dto.Comment {
	return dto.Comment{
		Pid:       comment.Pid,
		ArticleId: comment.ArticleId,
		Name:      comment.Name,
		Email:     comment.Email,
		Url:       comment.Url,
		Content:   comment.Content,
	}
}

func AddCommentReqToEntity(comment *proto.AddCommentReq) entity.Comment {
	return entity.Comment{
		Pid:       comment.Pid,
		ArticleId: comment.ArticleId,
		Name:      comment.Name,
		Email:     comment.Email,
		Url:       comment.Url,
		Content:   comment.Content,
	}
}
