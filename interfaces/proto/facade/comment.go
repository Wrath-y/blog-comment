package facade

import (
	"comment/interfaces/dto"
	"comment/interfaces/proto"
	"comment/launch/grpc/resp"
	"context"
)

type Comment struct{}

func (*Comment) GetComments(ctx context.Context, req *proto.GetCommentsReq) (*proto.Response, error) {
	return resp.Success(dto.H{})
}

func (*Comment) GetCountByArticleId(ctx context.Context, req *proto.OnlyArticleIdReq) (*proto.Response, error) {
	return resp.Success(dto.H{})
}

func (*Comment) AddComment(ctx context.Context, req *proto.AddCommentReq) (*proto.Response, error) {
	return resp.Success(dto.H{})
}
