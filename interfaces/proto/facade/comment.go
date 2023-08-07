package facade

import (
	"comment/application/service"
	grpcCtx "comment/infrastructure/common/context"
	"comment/infrastructure/common/errcode"
	"comment/interfaces/assembler"
	"comment/interfaces/dto"
	"comment/interfaces/proto"
	"comment/launch/grpc/resp"
	"context"
)

type Comment struct{}

func (*Comment) GetComments(ctx context.Context, req *proto.GetCommentsReq) (*proto.Response, error) {
	list, err := service.NewCommentApplicationService(grpcCtx.GetContext(ctx)).GetComments(req.ArticleId, req.Id, 15)
	if err != nil {
		return resp.FailWithErrCode(errcode.CommentsGetFailed)
	}
	return resp.Success(assembler.ToDTOs(list))
}

func (*Comment) GetCountByArticleId(ctx context.Context, req *proto.OnlyArticleIdReq) (*proto.Response, error) {
	count, err := service.NewCommentApplicationService(grpcCtx.GetContext(ctx)).GetCountByArticleId(req.ArticleId)
	if err != nil {
		return resp.FailWithErrCode(errcode.CommentCountGetFailed)
	}

	return resp.Success(count)
}

func (*Comment) AddComment(ctx context.Context, req *proto.AddCommentReq) (*proto.Response, error) {
	if err := service.NewCommentApplicationService(grpcCtx.GetContext(ctx)).AddComment(assembler.ToEntity(req)); err != nil {
		return resp.FailWithErrCode(errcode.CommentInsertFailed)
	}
	return resp.Success(dto.H{})
}
