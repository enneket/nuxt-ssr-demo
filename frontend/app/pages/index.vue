<template>
  <div class="home-container">
    <!-- SSR特性展示 -->
    <section>
      <h2>SSR 核心特性</h2>
      <div class="features-grid">
        <div class="feature-card">
          <h3>快速首屏加载</h3>
          <p>服务器直接返回完整HTML，减少客户端渲染时间，提升用户体验</p>
        </div>
        <div class="feature-card">
          <h3>SEO优化</h3>
          <p>搜索引擎能直接抓取完整页面内容，提升网站排名</p>
        </div>
        <div class="feature-card">
          <h3>统一开发体验</h3>
          <p>使用同一套代码库开发服务器端和客户端渲染</p>
        </div>
        <div class="feature-card">
          <h3>减少客户端资源消耗</h3>
          <p>部分渲染工作在服务器完成，降低客户端CPU和内存占用</p>
        </div>
        <div class="feature-card">
          <h3>实时数据获取</h3>
          <p>服务器端可以获取最新数据，确保首屏展示最新内容</p>
        </div>
        <div class="feature-card">
          <h3>跨设备兼容性</h3>
          <p>减少客户端JavaScript依赖，提高低配置设备兼容性</p>
        </div>
      </div>
    </section>

    <!-- 服务器数据展示 -->
    <section>
      <h2>服务器端数据</h2>
      <div v-if="serverData" class="data-cards-container">
        <div class="data-card">
          <h3>服务器时间</h3>
          <p class="server-time">{{ serverData.time }}</p>
        </div>
        <div class="data-card">
          <h3>Node.js版本</h3>
          <p>{{ serverData.serverInfo }}</p>
        </div>
        <div class="data-card">
          <h3>当前用户</h3>
          <p>{{ serverData.currentUser }}</p>
        </div>
        <div class="data-card">
          <h3>会话ID</h3>
          <p>{{ serverData.sessionId }}</p>
        </div>
        <div class="data-card">
          <h3>天气信息</h3>
          <p>{{ serverData.apiName }}: {{ serverData.temperature }}，{{ serverData.description }}</p>
        </div>
      </div>
      <div v-else class="loading">加载中...</div>
    </section>

    <!-- 文章入口 -->
    <div class="articles-link">
      <h2>文章列表</h2>
      <div class="articles-intro">
        <p>查看完整的文章列表，了解更多关于Nuxt 4 SSR的技术细节和最佳实践。</p>
        <NuxtLink to="/articles" class="articles-button">查看文章列表</NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup>
const { useAsyncData, useHead } = await import('#app');

// 设置页面元数据
useHead({
  title: 'Nuxt 4 SSR Demo - 首页',
  meta: [
    { name: 'description', content: '这是Nuxt 4 SSR演示项目的首页，展示了服务器端渲染的核心特性' }
  ]
});

// 服务器端数据获取
const { data: serverData } = await useAsyncData('serverData', async () => {
  // 模拟从数据库获取数据
  const databaseData = {
    currentUser: 'Demo User',
    sessionId: `session_${Math.random().toString(36).substr(2, 9)}`
  };

  // 模拟调用外部API（这里使用静态数据模拟）
  const externalApiData = {
    apiName: '模拟天气API',
    temperature: `${Math.round(15 + Math.random() * 10)}°C`,
    description: ['晴天', '多云', '阴天', '小雨'][Math.floor(Math.random() * 4)]
  };

  return {
    time: new Date().toLocaleString('zh-CN'),
    serverInfo: `Node.js ${process.version}`,
    ...databaseData,
    ...externalApiData
  };
});
</script>

<style scoped>
.home-container {
  display: flex;
  flex-direction: column;
  gap: 30px;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 20px;
  margin-top: 20px;
}

.feature-card {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 1px 5px rgba(0, 0, 0, 0.1);
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.feature-card h3 {
  margin-top: 0;
  color: #667eea;
}

.data-cards-container {
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
}

.data-card {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 1px 5px rgba(0, 0, 0, 0.1);
  min-width: 200px;
}

.data-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.server-time {
  font-size: 1.5rem;
  font-weight: bold;
  color: #667eea;
}

.loading {
  margin-top: 20px;
  color: #667eea;
  font-style: italic;
}

.articles-link {
  background: #f5f5f5;
  padding: 30px;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.articles-intro {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
  max-width: 600px;
  margin: 0 auto;
}

.articles-intro p {
  font-size: 1.1rem;
  color: #333;
  line-height: 1.6;
  margin: 0;
}

.articles-button {
  display: inline-block;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 12px 24px;
  border-radius: 25px;
  text-decoration: none;
  font-weight: bold;
  font-size: 1.1rem;
  transition: all 0.3s ease;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
}

.articles-button:hover {
  transform: translateY(-3px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);
  text-decoration: none;
  color: white;
}
</style>