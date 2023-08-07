package assembler

import (
	"comment/domain/comment/entity"
	"comment/interfaces/dto"
	"comment/interfaces/proto"
)

func ToDTOs(comments []entity.Comment) []dto.Comment {
	list := make([]dto.Comment, 0, len(comments))
	for _, v := range comments {
		list = append(list, dto.Comment{
			Id:         v.Id,
			Pid:        v.Pid,
			ArticleId:  v.ArticleId,
			Name:       v.Name,
			Email:      v.Email,
			Url:        v.Url,
			Content:    v.Content,
			CreateTime: v.CreateTime,
		})
	}
	return list
}

func ToEntity(comment *proto.AddCommentReq) entity.Comment {
	return entity.Comment{
		Pid:       comment.Pid,
		ArticleId: comment.ArticleId,
		Name:      comment.Name,
		Email:     comment.Email,
		Url:       comment.Url,
		Content:   comment.Content,
	}
}
