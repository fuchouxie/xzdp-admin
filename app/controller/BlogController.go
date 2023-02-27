package controller

import (
	"github.com/gin-gonic/gin"
	"xzdp-admin/app/dto"
	"xzdp-admin/app/service"
	"xzdp-admin/system/core/myGin"
)

type BlogController struct {
	blogService service.BlogService
}

func (c *BlogController) BlogList(ctx *gin.Context) {
	var req dto.BlogListReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	res, err := c.blogService.BlogList(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, res)
}

func (c *BlogController) GetBlogInfo(ctx *gin.Context) {
	var req dto.GetBlogInfoReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	res, err := c.blogService.GetBlog(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, res)
}

func (c *BlogController) Remove(ctx *gin.Context) {
	var req dto.RemoveBlogReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	err := c.blogService.Remove(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}

func (c *BlogController) BatchRemove(ctx *gin.Context) {
	var req dto.BatchRemoveBlogReq
	if err := ctx.ShouldBind(&req); err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	err := c.blogService.BatchRemove(req)
	if err != nil {
		myGin.Failure(ctx, err.Error())
		return
	}
	myGin.Success(ctx, nil)
}
