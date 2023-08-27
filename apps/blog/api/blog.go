package api

import (
	"github.com/gin-gonic/gin"
	"go-vblog/apps/blog"
	"net/http"
	"strconv"
)

func (h *HTTPAPI) CreateBlog(ctx *gin.Context) {
	req := blog.NewCreateBlogRequest()
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	instance, err := h.service.CreateBlog(ctx.Request.Context(), req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, instance)
}

func (h *HTTPAPI) DescribeBlog(ctx *gin.Context) {
	blogId := ctx.Param("id")
	bId, err := strconv.Atoi(blogId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	req := blog.NewDescribeBlogRequest(bId)
	instance, err := h.service.DescribeBlog(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, instance)
}

func (h *HTTPAPI) DeleteBlog(c *gin.Context) {
	// "/:id"   /abc   id=abc
	blogIdStr := c.Param("id")
	bid, err := strconv.Atoi(blogIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	req := blog.NewDeleteBlogRequest(bid)
	ins, err := h.service.DeleteBlog(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ins)
}

func (h *HTTPAPI) PutBlog(c *gin.Context) {
	// "/:id"   /abc   id=abc
	// 覆盖掉来自body里面的id参数
	blogIdStr := c.Param("id")
	bid, err := strconv.Atoi(blogIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	req := blog.NewPutUpdateBlogRequest(bid)

	// 读取来自body的json参数
	if err := c.BindJSON(req.CreateBlogRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ins, err := h.service.UpdateBlog(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ins)
}

func (h *HTTPAPI) PatchBlog(c *gin.Context) {
	// "/:id"   /abc   id=abc
	// 覆盖掉来自body里面的id参数
	blogIdStr := c.Param("id")
	bid, err := strconv.Atoi(blogIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	req := blog.NewPatchUpdateBlogRequest(bid)

	if err := c.BindJSON(req.CreateBlogRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ins, err := h.service.UpdateBlog(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ins)
}
