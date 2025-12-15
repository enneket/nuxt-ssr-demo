// 文章数据文件
// 存储所有文章数据，供列表页和详情页共享使用
export const articles = [
  {
    id: 1,
    title: 'Nuxt 4 SSR 入门指南',
    date: '2025-01-15',
    excerpt: '本文介绍了 Nuxt 4 SSR 的基本概念、核心特性和快速入门方法...',
    content: '# Nuxt 4 SSR 入门指南\n\n本文详细介绍了 Nuxt 4 SSR 的核心概念和使用方法。\n\n## 什么是 SSR？\n\n服务器端渲染（SSR）是一种将页面在服务器上渲染为 HTML 并返回给客户端的技术。\n\n## Nuxt 4 的优势\n\n- 快速首屏加载\n- 优秀的 SEO 表现\n- 一致的开发体验\n- 减少客户端资源消耗\n\n## 快速开始\n\n1. 安装 Nuxt 4\n2. 创建第一个页面\n3. 配置路由\n4. 添加数据获取逻辑'
  },
  {
    id: 2,
    title: '深入理解 Nuxt 4 的数据获取',
    date: '2025-01-20',
    excerpt: '本文深入探讨了 Nuxt 4 中的数据获取方法，包括 useAsyncData 和 useFetch...',
    content: '# 深入理解 Nuxt 4 的数据获取\n\n在 Nuxt 4 中，有多种方式获取数据，本文将详细介绍。\n\n## useAsyncData\n\nuseAsyncData 是 Nuxt 4 中最常用的数据获取函数，它确保服务器端和客户端数据的一致性。\n\n## useFetch\n\nuseFetch 是对 useAsyncData 和 $fetch 的封装，提供了更简洁的 API。\n\n## 数据预取\n\nNuxt 4 支持在页面导航前预取数据，提升用户体验。'
  },
  {
    id: 3,
    title: 'Nuxt 4 性能优化最佳实践',
    date: '2025-01-25',
    excerpt: '本文分享了 Nuxt 4 应用性能优化的最佳实践，包括代码分割、图片优化等...',
    content: '# Nuxt 4 性能优化最佳实践\n\n性能优化是现代 Web 应用开发的重要环节，本文介绍 Nuxt 4 的优化技巧。\n\n## 代码分割\n\nNuxt 4 自动进行代码分割，减少初始加载体积。\n\n## 图片优化\n\n使用响应式图片和懒加载提升性能。\n\n## 缓存策略\n\n合理配置缓存策略，减少服务器请求。'
  },
  {
    id: 4,
    title: 'Nuxt 4 与 Vue 3 的协同工作',
    date: '2025-01-30',
    excerpt: '本文介绍了 Nuxt 4 如何与 Vue 3 生态系统协同工作，包括组合式 API、Pinia 等...',
    content: '# Nuxt 4 与 Vue 3 的协同工作\n\nNuxt 4 基于 Vue 3 构建，充分利用了 Vue 3 的新特性。\n\n## 组合式 API\n\n使用 Vue 3 的组合式 API 编写更可维护的代码。\n\n## Pinia 状态管理\n\n集成 Pinia 提供现代化的状态管理解决方案。\n\n## 组件开发\n\n使用 Vue 3 的组件模型，构建可复用的 UI 组件。'
  }
];

export const getArticleById = (id) => {
  return articles.find(article => article.id === parseInt(id));
};

export const getRelatedArticles = (id, count = 2) => {
  return articles.filter(article => article.id !== parseInt(id)).slice(0, count);
};