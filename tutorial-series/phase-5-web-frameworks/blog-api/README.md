# Blog API - Go Web 开发完整示例

这是一个完整的博客 API 项目，涵盖了 Gin + GORM Web 开发的所有核心知识点。

## 项目特点

- ✅ **完整的分层架构**：Handler -> Service -> Repository
- ✅ **JWT 认证**：用户登录和权限控制
- ✅ **统一响应格式**：规范的 API 响应结构
- ✅ **错误处理**：全局错误处理中间件
- ✅ **文件上传**：支持图片上传
- ✅ **CORS 支持**：跨域配置
- ✅ **优雅关闭**：Graceful Shutdown
- ✅ **限流保护**：基于 IP 的限流
- ✅ **日志管理**：结构化日志
- ✅ **单元测试**：Handler 和 Service 层测试
- ✅ **Docker 部署**：容器化配置

## 快速开始

### 1. 安装依赖

\`\`\`bash
go mod download
\`\`\`

### 2. 运行项目

\`\`\`bash
go run cmd/server/main.go
\`\`\`

服务将在 `http://localhost:8080` 启动

### 3. 测试 API

\`\`\`bash
# 注册用户
curl -X POST http://localhost:8080/api/auth/register \\
  -H "Content-Type: application/json" \\
  -d '{"username":"admin","password":"123456","email":"admin@example.com"}'

# 登录获取 Token
curl -X POST http://localhost:8080/api/auth/login \\
  -H "Content-Type: application/json" \\
  -d '{"username":"admin","password":"123456"}'

# 创建文章（需要 Token）
curl -X POST http://localhost:8080/api/posts \\
  -H "Content-Type: application/json" \\
  -H "Authorization: Bearer YOUR_TOKEN" \\
  -d '{"title":"Hello World","content":"This is my first post"}'
\`\`\`

## 学习路径

本项目按照教程章节组织，您可以按以下顺序学习：

### 阶段 1：Gin 基础（教程 23-24 章）

**学习目标**：掌握 Gin 路由、中间件、参数绑定

**代码位置**：
- \`cmd/server/main.go\` - 路由定义和中间件注册
- \`internal/handler/user.go\` - 路由处理器示例
- \`internal/middleware/auth.go\` - 认证中间件
- \`internal/middleware/logger.go\` - 日志中间件

**关键知识点**：
- 路由参数 (\`:id\`) 和查询参数 (\`?page=1\`)
- 路由分组 (\`/api/v1\`)
- \`ShouldBindJSON\` 参数绑定
- 中间件的执行顺序

### 阶段 2：Gin 进阶（教程 29-31 章）

**学习目标**：错误处理、文件上传、JWT、CORS、优雅关闭

**代码位置**：
- \`pkg/response/response.go\` - 统一响应格式
- \`internal/middleware/error.go\` - 错误处理中间件
- \`internal/handler/upload.go\` - 文件上传
- \`pkg/jwt/jwt.go\` - JWT 工具
- \`internal/middleware/cors.go\` - CORS 配置
- \`cmd/server/main.go\` - 优雅关闭实现

**关键知识点**：
- 自定义错误类型
- \`panic/recover\` 机制
- JWT Token 生成与验证
- \`http.Server.Shutdown()\`

### 阶段 3：GORM 基础（教程 25-27 章）

**学习目标**：模型定义、CRUD、关联关系、事务、Hooks

**代码位置**：
- \`internal/model/model.go\` - 数据模型定义
- \`internal/repository/user.go\` - User CRUD 操作
- \`internal/repository/post.go\` - Post CRUD 和关联查询
- \`internal/model/model.go\` - BeforeCreate Hook 示例

**关键知识点**：
- \`gorm.Model\` 基础模型
- 一对多关联 (\`User has many Posts\`)
- \`Preload\` 预加载
- 事务处理
- 软删除

### 阶段 4：GORM 进阶（教程 32-34 章）

**学习目标**：复杂查询、性能优化、迁移管理

**代码位置**：
- \`internal/repository/post.go\` - 复杂查询示例
- \`internal/repository/user.go\` - 批量操作
- \`cmd/server/main.go\` - 数据库连接池配置

**关键知识点**：
- 子查询、联表查询
- \`CreateInBatches\` 批量插入
- 索引使用
- N+1 问题解决
- 连接池配置

### 阶段 5：项目架构（教程 35 章）

**学习目标**：分层架构、依赖注入、配置管理

**代码位置**：
- 整个项目结构
- \`config/config.yaml\` - 配置文件
- \`internal/service/\` - 业务逻辑层

**关键知识点**：
- Handler -> Service -> Repository 分层
- 接口抽象
- Viper 配置管理

### 阶段 6：测试（教程 36 章）

**学习目标**：单元测试、集成测试、Mock

**代码位置**：
- \`test/handler_test.go\` - Handler 测试
- \`test/service_test.go\` - Service 测试

**关键知识点**：
- \`httptest\` 测试 HTTP 接口
- Mock Repository
- 表格驱动测试

### 阶段 7：部署（教程 37 章）

**学习目标**：Docker 部署、日志、监控

**代码位置**：
- \`docker/Dockerfile\` - Docker 镜像
- \`docker/docker-compose.yml\` - 容器编排
- \`pkg/logger/logger.go\` - 日志管理
- \`internal/middleware/ratelimit.go\` - 限流

**关键知识点**：
- 多阶段构建
- 环境变量管理
- 结构化日志
- 限流策略

## 项目结构说明

\`\`\`
blog-api/
├── cmd/
│   └── server/
│       └── main.go              # 程序入口，路由注册
├── internal/                    # 私有代码
│   ├── handler/                 # HTTP 处理器（Controller 层）
│   │   ├── auth.go              # 认证相关接口
│   │   ├── user.go              # 用户相关接口
│   │   ├── post.go              # 文章相关接口
│   │   └── upload.go            # 文件上传
│   ├── service/                 # 业务逻辑层
│   │   ├── auth.go
│   │   ├── user.go
│   │   └── post.go
│   ├── repository/              # 数据访问层
│   │   ├── user.go
│   │   └── post.go
│   ├── model/                   # 数据模型
│   │   └── model.go
│   └── middleware/              # 中间件
│       ├── auth.go              # JWT 认证
│       ├── cors.go              # 跨域
│       ├── error.go             # 错误处理
│       ├── logger.go            # 日志
│       └── ratelimit.go         # 限流
├── pkg/                         # 公共工具库
│   ├── response/                # 统一响应
│   ├── jwt/                     # JWT 工具
│   └── logger/                  # 日志工具
├── config/
│   └── config.yaml              # 配置文件
├── test/                        # 测试
│   ├── handler_test.go
│   └── service_test.go
├── docker/                      # 部署相关
│   ├── Dockerfile
│   └── docker-compose.yml
└── docs/                        # 文档
    └── API.md                   # API 文档
\`\`\`

## API 文档

详见 [docs/API.md](docs/API.md)

## 技术栈

- **Web 框架**: Gin
- **ORM**: GORM
- **数据库**: SQLite (可切换为 MySQL/PostgreSQL)
- **认证**: JWT
- **日志**: Zap
- **配置**: Viper
- **测试**: testify

## License

MIT
