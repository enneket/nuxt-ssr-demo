<template>
  <div class="container">
    <header>
      <h1>Nuxt 4 SSR Demo</h1>
      <p>展示服务器端渲染的强大功能</p>
    </header>
    
    <main>
      <NuxtPage />
    </main>
    
    <footer>
      <p>© 2025 Nuxt 4 SSR Demo</p>
    </footer>
  </div>
</template>

<script setup>
const { useAsyncData, useHead } = await import('#app');

// 设置页面元数据
useHead({
  title: 'Nuxt 4 SSR Demo - 展示服务器端渲染的强大功能',
  meta: [
    { name: 'description', content: '这是一个 Nuxt 4 SSR 演示项目，展示了服务器端渲染的核心特性和优势' }
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
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  font-family: Arial, sans-serif;
}

header {
  text-align: center;
  padding: 40px 0;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border-radius: 10px;
  margin-bottom: 30px;
}

header h1 {
  margin: 0;
  font-size: 2.5rem;
}

header p {
  margin: 10px 0 0;
  font-size: 1.2rem;
}

main {
  display: flex;
  flex-direction: column;
  gap: 30px;
}

section {
  background: #f5f5f5;
  padding: 30px;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

section h2 {
  margin-top: 0;
  color: #333;
  border-bottom: 2px solid #667eea;
  padding-bottom: 10px;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 20px;
  margin-top: 20px;
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.data-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.feature-card {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 1px 5px rgba(0, 0, 0, 0.1);
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

footer {
  text-align: center;
  padding: 20px 0;
  margin-top: 30px;
  color: #666;
  border-top: 1px solid #eee;
}
</style>