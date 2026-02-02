<template>
  <div class="cache-demo-container">
    <h1>SSR 缓存策略演示</h1>

    <!-- 缓存机制说明 -->
    <section class="demo-section">
      <h2>📚 缓存机制说明</h2>
      <div class="info-card">
        <p>Nuxt 的 <code>useAsyncData</code> 提供了强大的缓存机制：</p>
        <ul>
          <li>✅ 服务端渲染时获取数据</li>
          <li>✅ 数据自动传递到客户端（避免重复请求）</li>
          <li>✅ 客户端导航时复用缓存数据</li>
          <li>✅ 可手动刷新数据</li>
        </ul>
      </div>
    </section>

    <!-- 基础缓存演示 -->
    <section class="demo-section">
      <h2>🔄 基础缓存演示</h2>
      <div class="info-card">
        <div v-if="pending" class="loading">加载中...</div>
        <div v-else-if="error" class="error">加载失败: {{ error.message }}</div>
        <div v-else>
          <p><strong>数据获取时间：</strong>{{ data?.time }}</p>
          <p><strong>随机数：</strong>{{ data?.random }}</p>
          <p><strong>请求次数：</strong>{{ data?.requestCount }}</p>
          <p class="hint">💡 离开页面再返回，数据会从缓存中读取，不会重新请求</p>
          
          <div class="button-group">
            <button @click="refresh" :disabled="pending" class="demo-button">
              {{ pending ? '刷新中...' : '🔄 刷新数据' }}
            </button>
            <button @click="clearCache" class="demo-button secondary">
              🗑️ 清除缓存
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- 多个缓存实例 -->
    <section class="demo-section">
      <h2>📦 多个独立缓存</h2>
      <div class="cache-grid">
        <div class="cache-card">
          <h3>缓存 A</h3>
          <div v-if="cacheA">
            <p><strong>时间：</strong>{{ cacheA.time }}</p>
            <p><strong>值：</strong>{{ cacheA.value }}</p>
          </div>
          <button @click="refreshCacheA" :disabled="pendingA" class="demo-button small">
            {{ pendingA ? '刷新中...' : '刷新 A' }}
          </button>
        </div>

        <div class="cache-card">
          <h3>缓存 B</h3>
          <div v-if="cacheB">
            <p><strong>时间：</strong>{{ cacheB.time }}</p>
            <p><strong>值：</strong>{{ cacheB.value }}</p>
          </div>
          <button @click="refreshCacheB" :disabled="pendingB" class="demo-button small">
            {{ pendingB ? '刷新中...' : '刷新 B' }}
          </button>
        </div>

        <div class="cache-card">
          <h3>缓存 C</h3>
          <div v-if="cacheC">
            <p><strong>时间：</strong>{{ cacheC.time }}</p>
            <p><strong>值：</strong>{{ cacheC.value }}</p>
          </div>
          <button @click="refreshCacheC" :disabled="pendingC" class="demo-button small">
            {{ pendingC ? '刷新中...' : '刷新 C' }}
          </button>
        </div>
      </div>
      <p class="hint">💡 每个缓存独立管理，可以单独刷新</p>
    </section>

    <!-- 缓存策略对比 -->
    <section class="demo-section">
      <h2>⚡ 缓存策略对比</h2>
      <div class="info-card">
        <table class="comparison-table">
          <thead>
            <tr>
              <th>策略</th>
              <th>说明</th>
              <th>适用场景</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>默认缓存</code></td>
              <td>数据在客户端导航时复用</td>
              <td>大部分场景</td>
            </tr>
            <tr>
              <td><code>server: false</code></td>
              <td>仅在客户端获取数据</td>
              <td>用户特定数据</td>
            </tr>
            <tr>
              <td><code>lazy: true</code></td>
              <td>不阻塞导航，异步加载</td>
              <td>非关键数据</td>
            </tr>
            <tr>
              <td><code>watch</code></td>
              <td>监听依赖变化自动刷新</td>
              <td>动态查询</td>
            </tr>
          </tbody>
        </table>
      </div>
    </section>

    <!-- 请求日志 -->
    <section class="demo-section">
      <h2>📝 请求日志</h2>
      <div class="info-card">
        <div class="log-container">
          <div v-for="(log, index) in requestLogs" :key="index" class="log-item">
            <span class="log-time">{{ log.time }}</span>
            <span :class="['log-type', `log-${log.type}`]">{{ log.type }}</span>
            <span class="log-message">{{ log.message }}</span>
          </div>
          <div v-if="requestLogs.length === 0" class="empty-log">
            暂无请求日志
          </div>
        </div>
        <button @click="clearLogs" class="demo-button secondary small">清除日志</button>
      </div>
    </section>

    <NuxtLink to="/" class="back-link">← 返回首页</NuxtLink>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useAsyncData, useHead } from '#app';

// 设置页面元数据
useHead({
  title: '缓存策略演示 - Nuxt 4 SSR Demo',
  meta: [
    { name: 'description', content: '展示 Nuxt 4 SSR 的缓存策略和数据管理' }
  ]
});

// 请求日志
const requestLogs = ref([]);
const addLog = (type, message) => {
  requestLogs.value.unshift({
    time: new Date().toLocaleTimeString('zh-CN'),
    type,
    message
  });
  // 限制日志数量
  if (requestLogs.value.length > 20) {
    requestLogs.value.pop();
  }
};

// 模拟 API 请求计数器
let requestCounter = 0;

// 基础缓存演示
const { data, pending, error, refresh, clear } = await useAsyncData('cache-demo', async () => {
  requestCounter++;
  addLog('fetch', `发起数据请求 #${requestCounter}`);
  
  // 模拟 API 延迟
  await new Promise(resolve => setTimeout(resolve, 500));
  
  const result = {
    time: new Date().toLocaleString('zh-CN'),
    random: Math.floor(Math.random() * 1000),
    requestCount: requestCounter
  };
  
  addLog('success', `数据获取成功，随机数: ${result.random}`);
  return result;
});

// 清除缓存
const clearCache = () => {
  clear();
  addLog('clear', '缓存已清除');
};

// 多个独立缓存
const { data: cacheA, pending: pendingA, refresh: refreshCacheA } = await useAsyncData('cache-a', async () => {
  await new Promise(resolve => setTimeout(resolve, 300));
  const result = {
    time: new Date().toLocaleTimeString('zh-CN'),
    value: Math.floor(Math.random() * 100)
  };
  addLog('fetch', `缓存 A 刷新，值: ${result.value}`);
  return result;
});

const { data: cacheB, pending: pendingB, refresh: refreshCacheB } = await useAsyncData('cache-b', async () => {
  await new Promise(resolve => setTimeout(resolve, 300));
  const result = {
    time: new Date().toLocaleTimeString('zh-CN'),
    value: Math.floor(Math.random() * 100)
  };
  addLog('fetch', `缓存 B 刷新，值: ${result.value}`);
  return result;
});

const { data: cacheC, pending: pendingC, refresh: refreshCacheC } = await useAsyncData('cache-c', async () => {
  await new Promise(resolve => setTimeout(resolve, 300));
  const result = {
    time: new Date().toLocaleTimeString('zh-CN'),
    value: Math.floor(Math.random() * 100)
  };
  addLog('fetch', `缓存 C 刷新，值: ${result.value}`);
  return result;
});

// 清除日志
const clearLogs = () => {
  requestLogs.value = [];
};
</script>

<style scoped>
.cache-demo-container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  color: #333;
  text-align: center;
  margin-bottom: 30px;
}

.demo-section {
  margin-bottom: 30px;
}

.demo-section h2 {
  color: #667eea;
  margin-bottom: 15px;
  font-size: 1.5rem;
}

.info-card {
  background: white;
  padding: 25px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.info-card p {
  margin: 12px 0;
  font-size: 1.05rem;
}

.info-card ul {
  margin: 15px 0;
  padding-left: 20px;
}

.info-card li {
  margin: 8px 0;
  line-height: 1.6;
}

code {
  background: #f1f5f9;
  padding: 2px 6px;
  border-radius: 3px;
  font-family: 'Courier New', monospace;
  font-size: 0.9em;
  color: #667eea;
}

.hint {
  margin-top: 15px;
  padding: 10px;
  background: #f0f7ff;
  border-left: 3px solid #667eea;
  font-size: 0.95rem;
  color: #555;
}

.loading {
  color: #667eea;
  font-style: italic;
  padding: 20px;
  text-align: center;
}

.error {
  color: #dc2626;
  padding: 20px;
  background: #fee;
  border-radius: 6px;
  text-align: center;
}

.button-group {
  display: flex;
  gap: 15px;
  margin-top: 20px;
  flex-wrap: wrap;
}

.demo-button {
  padding: 12px 24px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  font-weight: bold;
  cursor: pointer;
  transition: all 0.3s ease;
}

.demo-button:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

.demo-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.demo-button.secondary {
  background: #6b7280;
}

.demo-button.small {
  padding: 8px 16px;
  font-size: 0.9rem;
}

.cache-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
  margin-bottom: 15px;
}

.cache-card {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  border: 2px solid #e5e7eb;
}

.cache-card h3 {
  margin-top: 0;
  color: #667eea;
  font-size: 1.2rem;
}

.cache-card p {
  margin: 8px 0;
  font-size: 0.95rem;
}

.cache-card button {
  margin-top: 15px;
  width: 100%;
}

.comparison-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 15px;
}

.comparison-table th,
.comparison-table td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid #e5e7eb;
}

.comparison-table th {
  background: #f8f9fa;
  font-weight: bold;
  color: #374151;
}

.comparison-table tr:hover {
  background: #f9fafb;
}

.log-container {
  max-height: 300px;
  overflow-y: auto;
  background: #f8f9fa;
  padding: 15px;
  border-radius: 6px;
  margin-bottom: 15px;
  font-family: 'Courier New', monospace;
  font-size: 0.9rem;
}

.log-item {
  padding: 8px 0;
  border-bottom: 1px solid #e5e7eb;
  display: flex;
  gap: 10px;
  align-items: center;
}

.log-item:last-child {
  border-bottom: none;
}

.log-time {
  color: #6b7280;
  min-width: 80px;
}

.log-type {
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: bold;
  font-size: 0.85rem;
  min-width: 80px;
  text-align: center;
}

.log-fetch {
  background: #dbeafe;
  color: #1e40af;
}

.log-success {
  background: #dcfce7;
  color: #166534;
}

.log-clear {
  background: #fef3c7;
  color: #92400e;
}

.log-message {
  color: #374151;
  flex: 1;
}

.empty-log {
  text-align: center;
  color: #9ca3af;
  padding: 20px;
}

.back-link {
  display: inline-block;
  margin-top: 30px;
  color: #667eea;
  text-decoration: none;
  font-weight: bold;
  font-size: 1.1rem;
}

.back-link:hover {
  text-decoration: underline;
}
</style>
