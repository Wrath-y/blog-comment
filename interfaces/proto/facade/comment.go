package facade

import (
	"comment/application/service"
	grpcCtx "comment/infrastructure/common/context"
	"comment/infrastructure/util/errcode"
	"comment/interfaces/assembler"
	"comment/interfaces/dto"
	"comment/interfaces/proto"
	"comment/launch/grpc/resp"
	"context"
)

type Comment struct{}

func (*Comment) Insert(ctx context.Context, req *proto.CommentBaseInfo) (*proto.Response, error) {
	if err := service.NewCommentApplicationService(grpcCtx.GetContext(ctx)).Insert(assembler.PbToCommentEntity(req)); err != nil {
		return resp.FailWithErrCode(errcode.CommentInsertFailed)
	}

	return resp.Success(dto.H{})
}
