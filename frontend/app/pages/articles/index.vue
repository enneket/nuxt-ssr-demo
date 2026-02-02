<template>
  <div class="articles-container">
    <h1>文章列表</h1>
    
    <!-- 错误提示 -->
    <div v-if="error" class="error-message">
      <p>⚠️ 加载文章失败，请稍后重试</p>
      <p class="error-detail">{{ error.message }}</p>
    </div>
    
    <!-- 加载中 -->
    <div v-else-if="!articles" class="loading">
      <p>加载中...</p>
    </div>
    
    <!-- 文章列表 -->
    <div v-else class="articles-list">
      <div v-for="article in articles" :key="article.id" class="article-card">
        <NuxtLink :to="`/articles/${article.id}`">
          <h2>{{ article.title }}</h2>
          <p class="article-date">{{ article.date }}</p>
          <p class="article-excerpt">{{ article.excerpt }}</p>
        </NuxtLink>
      </div>
    </div>
    
    <NuxtLink to="/" class="back-link">← 返回首页</NuxtLink>
  </div>
</template>

<script setup>
import { useAsyncData, useHead } from '#app';

// 设置页面元数据
useHead({
  title: '文章列表 - Nuxt 4 SSR Demo',
  meta: [
    { name: 'description', content: '展示 Nuxt 4 SSR 文章列表功能' }
  ]
});

// 获取运行时配置
const runtimeConfig = useRuntimeConfig();
const apiBaseUrl = runtimeConfig.public.apiBaseUrl;

// 获取文章列表数据
const { data: articles, error } = await useAsyncData('articles', async () => {
  try {
    const response = await fetch(`${apiBaseUrl}/articles`);
    if (!response.ok) {
      throw new Error('Failed to fetch articles');
    }
    return await response.json();
  } catch (err) {
    console.error('Error fetching articles:', err);
    throw err;
  }
});
</script>

<style scoped>
.articles-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  color: #333;
  text-align: center;
  margin-bottom: 30px;
}

.articles-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.article-card {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.article-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
}

.article-card a {
  text-decoration: none;
  color: inherit;
  display: block;
}

.article-card h2 {
  margin-top: 0;
  color: #667eea;
}

.article-date {
  color: #666;
  font-size: 0.9rem;
  margin-bottom: 10px;
}

.article-excerpt {
  color: #333;
  line-height: 1.6;
}

.back-link {
  display: inline-block;
  margin-top: 30px;
  color: #667eea;
  text-decoration: none;
  font-weight: bold;
}

.back-link:hover {
  text-decoration: underline;
}

.error-message {
  background: #fee;
  border: 1px solid #fcc;
  border-radius: 8px;
  padding: 20px;
  margin: 20px 0;
  text-align: center;
}

.error-message p {
  margin: 10px 0;
  color: #c33;
}

.error-detail {
  font-size: 0.9rem;
  color: #666;
}

.loading {
  text-align: center;
  padding: 40px;
  color: #667eea;
  font-size: 1.2rem;
}
</style>