# API 文档

## 基础信息

- **Base URL**: `http://localhost:8080/api`
- **认证方式**: JWT Bearer Token

## 认证接口

### 1. 用户注册

**接口**: `POST /auth/register`

**请求体**:
\`\`\`json
{
  "username": "admin",
  "password": "123456",
  "email": "admin@example.com"
}
\`\`\`

**响应**:
\`\`\`json
{
  "code": 0,
  "message": "注册成功",
  "data": null
}
\`\`\`

### 2. 用户登录

**接口**: `POST /auth/login`

**请求体**:
\`\`\`json
{
  "username": "admin",
  "password": "123456"
}
\`\`\`

**响应**:
\`\`\`json
{
  "code": 0,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
\`\`\`

## 文章接口

### 3. 获取文章列表

**接口**: `GET /posts?page=1&page_size=10`

**响应**:
\`\`\`json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "title": "Hello World",
        "content": "This is my first post",
        "user_id": 1,
        "user": {
          "id": 1,
          "username": "admin",
          "email": "admin@example.com"
        },
        "created_at": "2025-12-18T00:00:00Z"
      }
    ],
    "total": 1,
    "page": 1
  }
}
\`\`\`

### 4. 获取文章详情

**接口**: `GET /posts/:id`

### 5. 创建文章（需要认证）

**接口**: `POST /posts`

**请求头**:
\`\`\`
Authorization: Bearer <your_token>
\`\`\`

**请求体**:
\`\`\`json
{
  "title": "My New Post",
  "content": "This is the content of my new post"
}
\`\`\`

### 6. 更新文章（需要认证）

**接口**: `PUT /posts/:id`

### 7. 删除文章（需要认证）

**接口**: `DELETE /posts/:id`

## 错误码

| Code | 说明 |
|------|------|
| 0 | 成功 |
| 400 | 参数错误 |
| 401 | 未认证 |
| 403 | 无权限 |
| 404 | 资源不存在 |
| 500 | 服务器错误 |
