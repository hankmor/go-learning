package repository

import (
	"blog-api/internal/model"

	"gorm.io/gorm"
)

// PostRepository 文章数据访问层
type PostRepository struct {
	db *gorm.DB
}

// NewPostRepository 创建 PostRepository 实例
func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

// Create 创建文章
func (r *PostRepository) Create(post *model.Post) error {
	return r.db.Create(post).Error
}

// FindByID 根据 ID 查询文章（包含作者信息）
func (r *PostRepository) FindByID(id uint) (*model.Post, error) {
	var post model.Post
	// Preload 预加载关联的 User 数据，避免 N+1 查询问题
	err := r.db.Preload("User").First(&post, id).Error
	return &post, err
}

// FindAll 查询所有文章（分页，包含作者信息）
func (r *PostRepository) FindAll(page, pageSize int) ([]model.Post, int64, error) {
	var posts []model.Post
	var total int64

	// 计算总数
	r.db.Model(&model.Post{}).Count(&total)

	// 分页查询，预加载作者信息
	offset := (page - 1) * pageSize
	err := r.db.Preload("User").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&posts).Error

	return posts, total, err
}

// FindByUserID 查询某个用户的所有文章
func (r *PostRepository) FindByUserID(userID uint) ([]model.Post, error) {
	var posts []model.Post
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&posts).Error
	return posts, err
}

// Update 更新文章
func (r *PostRepository) Update(post *model.Post) error {
	return r.db.Save(post).Error
}

// Delete 删除文章（软删除）
func (r *PostRepository) Delete(id uint) error {
	return r.db.Delete(&model.Post{}, id).Error
}

// Search 搜索文章（按标题或内容）
func (r *PostRepository) Search(keyword string, page, pageSize int) ([]model.Post, int64, error) {
	var posts []model.Post
	var total int64

	// 使用 LIKE 进行模糊搜索
	query := r.db.Model(&model.Post{}).Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%")

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Preload("User").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&posts).Error

	return posts, total, err
}
