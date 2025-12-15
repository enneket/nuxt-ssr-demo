package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Article 文章结构体
type Article struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Date    string    `json:"date"`
	Excerpt string    `json:"excerpt"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

// 模拟数据库
var articles = []Article{
	{
		ID:      1,
		Title:   "Nuxt 4 SSR 入门指南",
		Date:    "2025-01-15",
		Excerpt: "本文介绍了 Nuxt 4 SSR 的基本概念、核心功能和使用方法...",
		Content: `# Nuxt 4 SSR 入门指南

本文详细介绍了 Nuxt 4 SSR 的核心概念、核心功能和使用方法。包括：

## 什么是 SSR

服务器端渲染（SSR）是指将 Vue 组件在服务器端渲染成 HTML 字符串，然后发送到客户端，客户端将这些静态标记 "激活" 为完全可交互的应用程序。

## Nuxt 4 的优势

1. 快速首屏加载
2. 更好的 SEO
3. 统一的开发体验
4. 减少客户端资源消耗
5. 更广泛的浏览器支持

## 如何开始使用

使用 npx nuxi init 命令可以快速创建一个 Nuxt 4 项目。`,
		Created: time.Now(),
	},
	{
		ID:      2,
		Title:   "Vue 3 Composition API 最佳实践",
		Date:    "2025-01-10",
		Excerpt: "探索 Vue 3 Composition API 的最佳实践和高级技巧...",
		Content: `# Vue 3 Composition API 最佳实践

Vue 3 的 Composition API 提供了一种更灵活的方式来组织组件逻辑。本文将介绍一些最佳实践：

## 逻辑复用

使用组合式函数（Composables）来复用逻辑：

func useCounter() {
  const count = ref(0)
  const increment = () => count.value++
  const decrement = () => count.value--
  
  return { count, increment, decrement }
}

## 响应式系统

充分利用 ref 和 reactive 来创建响应式数据。`,
		Created: time.Now(),
	},
	{
		ID:      3,
		Title:   "Go 语言 RESTful API 设计",
		Date:    "2025-01-05",
		Excerpt: "学习如何使用 Go 语言设计和实现高质量的 RESTful API...",
		Content: `# Go 语言 RESTful API 设计

Go 语言非常适合构建高性能的 RESTful API。本文将介绍如何设计和实现高质量的 API：

## RESTful 原则

1. 使用合适的 HTTP 方法
2. 设计清晰的资源路径
3. 使用标准的状态码
4. 提供一致的响应格式

## Go 实现

使用 Go 的 net/http 包或第三方框架（如 Gin、Echo）来实现 API。`,
		Created: time.Now(),
	},
}

// corsMiddleware 处理跨域请求
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// getArticles 获取所有文章列表
func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// getArticleById 根据ID获取文章详情
func getArticleById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 从URL路径获取ID参数
	idStr := r.URL.Path[len("/api/articles/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid article ID"})
		return
	}

	// 查找文章
	for _, article := range articles {
		if article.ID == id {
			json.NewEncoder(w).Encode(article)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Article not found"})
}

// getServerInfo 获取服务器信息
func getServerInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	info := map[string]interface{}{
		"server":    "Go HTTP Server",
		"version":   "1.0.0",
		"timestamp": time.Now().Unix(),
	}
	json.NewEncoder(w).Encode(info)
}

func main() {
	// 创建路由器
	mux := http.NewServeMux()

	// API路由
	mux.HandleFunc("/api/articles", getArticles)
	mux.HandleFunc("/api/articles/", getArticleById) // 注意末尾的斜杠，用于匹配动态ID
	mux.HandleFunc("/api/server-info", getServerInfo)

	// 应用中间件
	handler := corsMiddleware(mux)

	// 启动服务器
	port := 8000
	fmt.Printf("Server running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
