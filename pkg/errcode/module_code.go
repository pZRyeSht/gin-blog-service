package errcode

var (
	ErrorGetTagListFail = NewError(20210101, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20210102, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20210103, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20210104, "删除标签失败")
	ErrorCountTagFail   = NewError(202100105, "统计标签失败")

	ErrorGetArticleFail    = NewError(20210201, "获取单个文章失败")
	ErrorGetArticlesFail   = NewError(20210202, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(20210203, "创建文章失败")
	ErrorUpdateArticleFail = NewError(20210204, "更新文章失败")
	ErrorDeleteArticleFail = NewError(20210205, "删除文章失败")

	ErrorUploadFileFail = NewError(20210301, "上传文件失败")
)
