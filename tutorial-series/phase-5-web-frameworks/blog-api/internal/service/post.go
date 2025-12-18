package service

import (
	"blog-api/internal/model"
	"blog-api/internal/repository"
	"errors"
)

// PostService 文章服务层
type PostService struct {
	postRepo *repository.PostRepository
}

// NewPostService 创建 PostService 实例
func NewPostService(postRepo *repository.PostRepository) *PostService {
	return &PostService{postRepo: postRepo}
}

// CreatePost 创建文章
func (s *PostService) CreatePost(userID uint, title, content string) error {
	// 业务逻辑：检查标题长度
	if len(title) < 5 {
		return errors.New("标题至少需要 5 个字符")
	}

	if len(content) < 10 {
		return errors.New("内容至少需要 10 个字符")
	}

	post := &model.Post{
		UserID:  userID,
		Title:   title,
		Content: content,
	}

	return s.postRepo.Create(post)
}

// GetPost 获取文章详情
func (s *PostService) GetPost(id uint) (*model.Post, error) {
	return s.postRepo.FindByID(id)
}

// GetPosts 获取文章列表（分页）
func (s *PostService) GetPosts(page, pageSize int) ([]model.Post, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	return s.postRepo.FindAll(page, pageSize)
}

// GetUserPosts 获取用户的文章列表
func (s *PostService) GetUserPosts(userID uint) ([]model.Post, error) {
	return s.postRepo.FindByUserID(userID)
}

// UpdatePost 更新文章
func (s *PostService) UpdatePost(postID, userID uint, title, content string) error {
	// 1. 查询文章
	post, err := s.postRepo.FindByID(postID)
	if err != nil {
		return errors.New("文章不存在")
	}

	// 2. 权限检查：只有作者才能修改
	if post.UserID != userID {
		return errors.New("无权修改此文章")
	}

	// 3. 更新字段
	post.Title = title
	post.Content = content

	return s.postRepo.Update(post)
}

// DeletePost 删除文章
func (s *PostService) DeletePost(postID, userID uint) error {
	// 1. 查询文章
	post, err := s.postRepo.FindByID(postID)
	if err != nil {
		return errors.New("文章不存在")
	}

	// 2. 权限检查：只有作者才能删除
	if post.UserID != userID {
		return errors.New("无权删除此文章")
	}

	return s.postRepo.Delete(postID)
}

// SearchPosts 搜索文章
func (s *PostService) SearchPosts(keyword string, page, pageSize int) ([]model.Post, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	return s.postRepo.Search(keyword, page, pageSize)
}
