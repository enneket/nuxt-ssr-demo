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

// 获取运行时配置
const runtimeConfig = useRuntimeConfig();
const apiBaseUrl = runtimeConfig.public.apiBaseUrl;

// 获取文章详情数据
const { data: article } = await useAsyncData(`article-${id}`, async () => {
  // 从后端API获取文章详情
  const response = await fetch(`${apiBaseUrl}/articles/${id}`);
  if (!response.ok) {
    throw new Error('Failed to fetch article');
  }
  return await response.json();
}, {
  // API失败时的默认值
  default: () => ({
    id: id,
    title: '默认文章标题 - Nuxt 4 SSR Demo',
    excerpt: '这是一篇默认的文章摘要，用于演示Nuxt 4 SSR的动态SEO功能。',
    content: '# 默认文章标题\n\n这是一篇默认的文章内容，用于演示Nuxt 4 SSR的动态SEO功能。\n\n## 章节标题\n\n这是章节内容。',
    date: new Date().toLocaleDateString('zh-CN'),
    image: 'https://site.com/og-image.jpg',
    author: 'Nuxt 4 SSR Demo',
    category: '技术博客',
    created: new Date().toISOString(),
    updated: new Date().toISOString(),
    tags: ['Nuxt', 'SSR', 'Vue']
  })
});

// 生成结构化数据
const jsonLd = computed(() => {
  // 确保article.value存在，否则返回空字符串
  if (!article.value) return '';
  
  // 获取必要的属性，确保所有值都有合理的默认值
  const title = article.value.title || '默认文章标题';
  const excerpt = article.value.excerpt || '';
  const content = article.value.content || '';
  const imageUrl = article.value.image || 'https://site.com/og-image.jpg';
  const authorName = article.value.author || 'Nuxt 4 SSR Demo';
  const category = article.value.category || '技术博客';
  const tags = article.value.tags || [];
  
  // 计算阅读时间（假设平均阅读速度为每分钟200个字）
  const contentWords = content.split(/\s+/).length || 0;
  const readTimeMinutes = Math.ceil(contentWords / 200);
  
  // 处理日期
  const createdDate = article.value.created ? new Date(article.value.created).toISOString() : new Date().toISOString();
  const updatedDate = article.value.updated ? new Date(article.value.updated).toISOString() : createdDate;
  
  const data = {
    "@context": "https://schema.org",
    "@type": "Article",
    "headline": title,
    "description": excerpt,
    "datePublished": createdDate,
    "dateModified": updatedDate,
    "articleSection": category,
    "wordCount": contentWords,
    "timeRequired": `PT${readTimeMinutes}M`, // ISO 8601 格式的阅读时间
    "author": {
      "@type": "Person",
      "name": authorName,
      "url": "https://site.com/author"
    },
    "publisher": {
      "@type": "Organization",
      "name": "Nuxt 4 SSR Demo",
      "logo": {
        "@type": "ImageObject",
        "url": "https://site.com/logo.png",
        "width": 600,
        "height": 60
      }
    },
    "image": {
      "@type": "ImageObject",
      "url": imageUrl,
      "width": 1200,
      "height": 630,
      "caption": title
    },
    "mainEntityOfPage": {
      "@type": "WebPage",
      "@id": `https://site.com/articles/${id}`
    },
    "inLanguage": "zh-CN",
    "isAccessibleForFree": true, // 表示内容是免费访问的
    "license": "https://creativecommons.org/licenses/by/4.0/", // 示例许可证
    "keywords": tags // 如果有标签的话
  };
  
  return JSON.stringify(data);
});

// 获取相关文章（从所有文章中筛选）
const { data: relatedArticles } = await useAsyncData(`related-articles-${id}`, async () => {
  // 从后端API获取所有文章
  const response = await fetch(`${apiBaseUrl}/articles`);
  if (!response.ok) {
    throw new Error('Failed to fetch related articles');
  }
  const allArticles = await response.json();
  // 筛选出除当前文章外的前2篇作为相关文章
  return allArticles.filter(a => a.id !== id).slice(0, 2);
}, {
  // API失败时的默认值
  default: () => [
    {
      id: 1,
      title: '相关文章标题1',
      date: new Date().toLocaleDateString('zh-CN')
    },
    {
      id: 2,
      title: '相关文章标题2',
      date: new Date().toLocaleDateString('zh-CN')
    }
  ]
});

// 设置页面元数据
useHead({
  title: computed(() => article.value?.title || '文章详情 - Nuxt 4 SSR Demo'),
  meta: [
    { 
      name: 'description', 
      content: computed(() => article.value?.excerpt || 'Nuxt 4 SSR 文章详情') 
    },
    { 
      name: 'keywords', 
      content: computed(() => {
        if (!article.value) return 'Nuxt, SSR, Vue';
        // 从文章内容中提取关键词（简单示例）
        const content = article.value.content.toLowerCase();
        const keywords = [];
        
        // 检查常见技术关键词
        const techKeywords = ['nuxt', 'vue', 'ssr', 'server-side rendering', '前端', 'web开发', 'javascript', 'typescript', '前端开发'];
        for (const keyword of techKeywords) {
          if (content.includes(keyword)) {
            keywords.push(keyword);
          }
        }
        
        // 如果没有匹配的关键词，使用默认值
        return keywords.length > 0 ? keywords.join(', ') : 'Nuxt, SSR, Vue, 前端开发';
      }) 
    },
    // Open Graph 标签（用于社交媒体分享）
    { 
      property: 'og:title', 
      content: computed(() => article.value?.title || '文章详情 - Nuxt 4 SSR Demo') 
    },
    { 
      property: 'og:description', 
      content: computed(() => article.value?.excerpt || 'Nuxt 4 SSR 文章详情') 
    },
    { 
      property: 'og:type', 
      content: 'article' 
    },
    { 
      property: 'og:url', 
      content: computed(() => `https://site.com/articles/${id}`) 
    },
    { 
      property: 'og:site_name', 
      content: 'Nuxt 4 SSR Demo' 
    },
    { 
      property: 'og:locale', 
      content: 'zh_CN' 
    },
    { 
      property: 'og:image', 
      content: computed(() => article.value?.image || 'https://site.com/og-image.jpg') 
    },
    // 文章特定的 Open Graph 标签
    { 
      property: 'article:section', 
      content: computed(() => article.value?.category || '技术博客') 
    },
    { 
      property: 'article:published_time', 
      content: computed(() => {
        if (!article.value || !article.value.created) return '';
        return new Date(article.value.created).toISOString();
      }) 
    },
    { 
      property: 'article:modified_time', 
      content: computed(() => {
        if (!article.value || !article.value.updated) return '';
        return new Date(article.value.updated).toISOString();
      }) 
    },
    { 
      property: 'article:author', 
      content: computed(() => article.value?.author || 'Nuxt 4 SSR Demo') 
    },
    // Twitter 卡片标签
    { 
      name: 'twitter:card', 
      content: 'summary_large_image' 
    },
    { 
      name: 'twitter:title', 
      content: computed(() => article.value?.title || '文章详情 - Nuxt 4 SSR Demo') 
    },
    { 
      name: 'twitter:description', 
      content: computed(() => article.value?.excerpt || 'Nuxt 4 SSR 文章详情') 
    },
    { 
      name: 'twitter:image', 
      content: computed(() => article.value?.image || 'https://site.com/og-image.jpg') 
    },
    { 
      name: 'twitter:site', 
      content: '@NuxtDemo' 
    },
    { 
      name: 'twitter:creator', 
      content: computed(() => article.value?.author || '@NuxtDemo') 
    }
  ],
  // 结构化数据 - 使用script选项注入JSON-LD
  script: [
    {
      type: 'application/ld+json',
      // 使用innerHTML而不是content，因为content会被转义
      innerHTML: jsonLd,
      // 确保脚本只在服务端渲染时添加
      once: true
    }
  ]
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