# nuxt-ssr-demo

Nuxt SSR 结合 Go 后端 API 的示例项目。

## 项目结构

```
ssr-demo/
├── backend/             # Go 后端服务
│   ├── main.go         # 后端主程序
│   └── go.mod          # Go 模块文件
├── frontend/           # Nuxt 4 前端应用
│   ├── app/            # 应用源码
│   │   ├── pages/      # 页面目录
│   │   │   ├── index.vue          # 首页
│   │   │   ├── ssr-demo.vue       # SSR 水合演示
│   │   │   ├── cache-demo.vue     # 缓存策略演示
│   │   │   ├── performance.vue    # 性能监控面板
│   │   │   └── articles/          # 文章相关页面
│   │   │       ├── index.vue      # 文章列表页
│   │   │       └── [id].vue       # 文章详情页
│   │   └── app.vue                # 根组件
│   ├── nuxt.config.ts             # Nuxt 配置
│   ├── package.json               # 项目依赖
│   ├── .env.example               # 环境变量示例
│   └── public/                    # 静态资源
└── README.md                      # 项目文档
```

## 安装和运行

### 1. 安装依赖

#### 前端依赖
```bash
cd frontend
npm install
```

#### 后端依赖
```bash
cd backend
go mod tidy
```

### 2. 配置环境变量

复制前端环境变量示例文件：
```bash
cd frontend
cp .env.example .env
```

根据实际情况修改 `.env` 文件中的配置：
- `NUXT_PUBLIC_API_BASE_URL`: 后端 API 地址（默认：http://localhost:8000/api）
- `NUXT_PUBLIC_SITE_URL`: 站点 URL（默认：http://localhost:3001）

### 3. 运行项目

#### 启动后端服务
```bash
cd backend
go run main.go
# 服务将运行在 http://localhost:8000
```

#### 启动前端开发服务器
```bash
cd frontend
npm run dev
# 服务将运行在 http://localhost:3001
```

#### 构建生产版本
```bash
cd frontend
npm run build
```

#### 静态站点生成
```bash
cd frontend
npm run generate
# 生成的静态文件将放在 .output/public 目录
# 预览静态站点
npx serve frontend/.output/public
```

## 功能特性

### SSR 核心功能演示
- **水合（Hydration）演示**：展示服务端渲染内容如何在客户端激活
- **缓存策略演示**：演示 `useAsyncData` 的缓存机制和数据管理
- **性能监控面板**：实时显示 TTFB、FCP、LCP 等性能指标

### 文章系统
- 文章列表页面
- 文章详情页面（动态路由）
- 完整的 SEO 优化（meta 标签、Open Graph、JSON-LD）

### SEO 优化
- 动态生成 sitemap
- 每个页面独立的 meta 标签
- 结构化数据（Schema.org）

## API 接口

后端提供以下 API 接口：

### GET /api/articles
获取所有文章列表

### GET /api/articles/:id
获取单篇文章详情

## 许可证

MIT License
