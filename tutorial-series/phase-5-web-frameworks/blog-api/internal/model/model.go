package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"` // 用户名，唯一索引
	Password string `gorm:"not null" json:"-"`                    // 密码，不返回给前端
	Email    string `gorm:"uniqueIndex;not null" json:"email"`    // 邮箱，唯一索引
	Posts    []Post `gorm:"foreignKey:UserID" json:"posts,omitempty"` // 一对多：一个用户有多篇文章
}

// Post 文章模型
type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`           // 文章标题
	Content string `gorm:"type:text;not null" json:"content"` // 文章内容
	UserID  uint   `gorm:"not null;index" json:"user_id"`   // 外键：作者 ID
	User    User   `gorm:"foreignKey:UserID" json:"user,omitempty"` // 关联：文章作者
}

// BeforeCreate Hook: 在创建用户前自动加密密码
// 这是 GORM 的生命周期钩子，会在执行 INSERT 前自动调用
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// 使用 bcrypt 加密密码
	// bcrypt.DefaultCost 是推荐的加密强度
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword 验证密码是否正确
func (u *User) CheckPassword(password string) bool {
	// bcrypt.CompareHashAndPassword 会自动处理盐值
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
