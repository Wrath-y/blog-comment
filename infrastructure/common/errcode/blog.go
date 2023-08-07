package errcode

var (
	BlogNetworkBusy  = &ErrCode{40000, "网络繁忙，请稍后重试", ""}
	BlogInvalidParam = &ErrCode{40001, "无效的参数", ""}
	BlogInvalidSign  = &ErrCode{40002, "无效的签名", ""}
	BlogBodyTooLarge = &ErrCode{40003, "请求消息体过大", ""}
)

var (
	CommentInsertFailed   = &ErrCode{40300, "评论发布失败", ""}
	CommentsGetFailed     = &ErrCode{40301, "评论列表获取失败", ""}
	CommentCountGetFailed = &ErrCode{40302, "评论数量获取失败", ""}
	CommentArticleIdErr   = &ErrCode{40303, "articleId异常", ""}
)
