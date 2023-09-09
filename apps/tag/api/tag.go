package api

import (
	"go-vblog/apps/tag"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *HTTPAPI) AddTag(c *gin.Context) {
	// 接收请求参数
	req := tag.NewAddTagRequest()
	if err := c.BindJSON(&req.Tags); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	set, err := h.service.AddTag(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, set)
}

func (h *HTTPAPI) RemoveTag(c *gin.Context) {
	// 接收请求参数
	req := tag.NewRemoveTagRequest()
	bint, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	req.TagIds = []int{bint}

	set, err := h.service.RemoveTag(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, set)
}
