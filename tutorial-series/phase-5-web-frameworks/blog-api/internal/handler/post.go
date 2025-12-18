package handler

import (
	"blog-api/internal/service"
	"blog-api/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostHandler 文章处理器
type PostHandler struct {
	postService *service.PostService
}

// NewPostHandler 创建 PostHandler 实例
func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

// CreatePostRequest 创建文章请求
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=5"`
	Content string `json:"content" binding:"required,min=10"`
}

// Create 创建文章
// POST /api/posts
// 需要认证
func (h *PostHandler) Create(c *gin.Context) {
	var req CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	// 从上下文获取用户 ID（由 JWT 中间件设置）
	userID := c.GetUint("user_id")

	if err := h.postService.CreatePost(userID, req.Title, req.Content); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.SuccessWithMessage(c, "文章创建成功", nil)
}

// Get 获取文章详情
// GET /api/posts/:id
func (h *PostHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	post, err := h.postService.GetPost(uint(id))
	if err != nil {
		response.Error(c, 404, "文章不存在")
		return
	}

	response.Success(c, post)
}

// List 获取文章列表
// GET /api/posts?page=1&page_size=10
func (h *PostHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	posts, total, err := h.postService.GetPosts(page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{
		"list":  posts,
		"total": total,
		"page":  page,
	})
}

// Update 更新文章
// PUT /api/posts/:id
// 需要认证
func (h *PostHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("user_id")

	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	if err := h.postService.UpdatePost(uint(id), userID, req.Title, req.Content); err != nil {
		response.Error(c, 403, err.Error())
		return
	}

	response.SuccessWithMessage(c, "文章更新成功", nil)
}

// Delete 删除文章
// DELETE /api/posts/:id
// 需要认证
func (h *PostHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("user_id")

	if err := h.postService.DeletePost(uint(id), userID); err != nil {
		response.Error(c, 403, err.Error())
		return
	}

	response.SuccessWithMessage(c, "文章删除成功", nil)
}
