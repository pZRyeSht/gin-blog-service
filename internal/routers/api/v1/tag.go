package v1

import (
	"gin_blog_service/global"
	"gin_blog_service/pkg/app"
	"gin_blog_service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func NewTag() Tag {
	return Tag{}
}

type CountTagRequest struct {
	State uint8 `form:"state, default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state, default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required, min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state, default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state, default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"requied,min=3.max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required, gte=1"`
}

// @Summary 获取单个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (a Tag) Get(c *gin.Context) {}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (a Tag) List(c *gin.Context) {
	param := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state, default=1" binding:"oneof=0 1"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParam.WithDetails(errs.Errors()...))
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 新增标签
// @Produce json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (a Tag) Create(c *gin.Context) {}

// @Summary 更新标签
// @Produce json
// @Param id path int true "标签ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [put]
func (a Tag) Update(c *gin.Context) {}

// @Summary 删除标签
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {string} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [delete]
func (a Tag) Delete(c *gin.Context) {}
