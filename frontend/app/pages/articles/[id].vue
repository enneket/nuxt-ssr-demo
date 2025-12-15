<template>
  <div class="article-detail">
    <div v-if="article" class="article-content">
      <h1>{{ article.title }}</h1>
      <p class="article-date">{{ article.date }}</p>
      <div class="article-body" v-html="renderMarkdown(article.content)"></div>
      
      <!-- 相关文章 -->
      <div class="related-articles">
        <h2>相关文章</h2>
        <div class="related-list">
          <div v-for="related in relatedArticles" :key="related.id" class="related-card">
            <NuxtLink :to="`/articles/${related.id}`">
              <h3>{{ related.title }}</h3>
              <p class="related-date">{{ related.date }}</p>
            </NuxtLink>
          </div>
        </div>
      </div>
      
      <div class="article-actions">
        <NuxtLink to="/articles" class="back-link">← 返回文章列表</NuxtLink>
        <NuxtLink to="/" class="home-link">返回首页</NuxtLink>
      </div>
    </div>
    <div v-else class="loading">
      加载中...
    </div>
  </div>
</template>

<script setup>
import { useAsyncData, useHead, useRoute } from '#app';
import { computed } from 'vue';

// 获取路由参数
const route = useRoute();
const id = parseInt(route.params.id);

// 获取文章详情数据
const { data: article } = await useAsyncData(`article-${id}`, async () => {
  // 从后端API获取文章详情
  const response = await fetch(`http://localhost:8000/api/articles/${id}`);
  if (!response.ok) {
    throw new Error('Failed to fetch article');
  }
  return await response.json();
});

// 设置页面元数据
useHead({
  title: computed(() => article.value?.title || '文章详情 - Nuxt 4 SSR Demo')
});

// 获取相关文章（从所有文章中筛选）
const { data: relatedArticles } = await useAsyncData(`related-articles-${id}`, async () => {
  // 从后端API获取所有文章
  const response = await fetch('http://localhost:8000/api/articles');
  if (!response.ok) {
    throw new Error('Failed to fetch related articles');
  }
  const allArticles = await response.json();
  // 筛选出除当前文章外的前2篇作为相关文章
  return allArticles.filter(a => a.id !== id).slice(0, 2);
});

// 模拟Markdown渲染
const renderMarkdown = (content) => {
  return content
    .replace(/^# (.*$)/gm, '<h1>$1</h1>')
    .replace(/^## (.*$)/gm, '<h2>$1</h2>')
    .replace(/\n/g, '<br>');
};
</script>

<style scoped>
.article-detail {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.article-content {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.article-date {
  color: #666;
  font-size: 0.9rem;
  margin-bottom: 20px;
}

.article-body {
  line-height: 1.8;
  color: #333;
  margin-bottom: 40px;
}

.article-body h1 {
  font-size: 1.8rem;
  margin-top: 0;
  margin-bottom: 20px;
  color: #333;
}

.article-body h2 {
  font-size: 1.5rem;
  margin-top: 30px;
  margin-bottom: 15px;
  color: #444;
}

.related-articles {
  margin-top: 50px;
  padding-top: 30px;
  border-top: 1px solid #eee;
}

.related-articles h2 {
  font-size: 1.5rem;
  margin-bottom: 20px;
  color: #333;
}

.related-list {
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
}

.related-card {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  flex: 1;
  min-width: 250px;
  transition: all 0.3s ease;
}

.related-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
}

.related-card a {
  text-decoration: none;
  color: inherit;
  display: block;
}

.related-card h3 {
  margin-top: 0;
  color: #667eea;
  margin-bottom: 10px;
}

.related-date {
  color: #666;
  font-size: 0.85rem;
}

.article-actions {
  margin-top: 40px;
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
}

.back-link, .home-link {
  display: inline-block;
  background: #667eea;
  color: white;
  padding: 10px 20px;
  border-radius: 4px;
  text-decoration: none;
  font-weight: bold;
  transition: all 0.3s ease;
}

.back-link:hover, .home-link:hover {
  background: #764ba2;
  text-decoration: none;
  color: white;
}

.loading {
  margin-top: 20px;
  color: #667eea;
  font-style: italic;
  text-align: center;
  font-size: 1.2rem;
}
</style>