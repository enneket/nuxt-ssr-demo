<template>
  <div class="ssr-demo-container">
    <h1>SSR 水合（Hydration）演示</h1>
    
    <!-- 服务端渲染内容 -->
    <section class="demo-section">
      <h2>🖥️ 服务端渲染内容</h2>
      <div class="info-card">
        <p><strong>服务器时间：</strong>{{ serverTime }}</p>
        <p><strong>渲染环境：</strong>{{ renderEnvironment }}</p>
        <p><strong>随机数（服务端生成）：</strong>{{ serverRandomNumber }}</p>
        <p class="hint">💡 这些内容在服务器端生成，刷新页面会看到不同的值</p>
      </div>
    </section>

    <!-- 客户端水合内容 -->
    <section class="demo-section">
      <h2>💻 客户端水合内容</h2>
      <div class="info-card">
        <p><strong>客户端时间：</strong>{{ clientTime || '等待水合...' }}</p>
        <p><strong>水合状态：</strong>
          <span :class="isHydrated ? 'status-success' : 'status-pending'">
            {{ isHydrated ? '✅ 已水合' : '⏳ 未水合' }}
          </span>
        </p>
        <p><strong>随机数（客户端生成）：</strong>{{ clientRandomNumber || '等待水合...' }}</p>
        <p class="hint">💡 这些内容在客户端激活后生成，查看页面源代码不会看到这些值</p>
      </div>
    </section>

    <!-- 交互演示 -->
    <section class="demo-section">
      <h2>🎯 交互演示</h2>
      <div class="info-card">
        <p><strong>点击计数：</strong>{{ clickCount }}</p>
        <button @click="incrementCount" class="demo-button">点击我 +1</button>
        <p class="hint">💡 按钮交互只在客户端水合后才能工作</p>
      </div>
    </section>

    <!-- 生命周期演示 -->
    <section class="demo-section">
      <h2>🔄 生命周期演示</h2>
      <div class="info-card">
        <div class="lifecycle-log">
          <div v-for="(log, index) in lifecycleLogs" :key="index" class="log-item">
            <span class="log-time">{{ log.time }}</span>
            <span :class="['log-type', `log-${log.type}`]">{{ log.type }}</span>
            <span class="log-message">{{ log.message }}</span>
          </div>
        </div>
        <p class="hint">💡 观察组件在服务端和客户端的生命周期执行顺序</p>
      </div>
    </section>

    <!-- 源代码对比 -->
    <section class="demo-section">
      <h2>📄 查看页面源代码</h2>
      <div class="info-card">
        <p>右键点击页面，选择"查看页面源代码"，你会发现：</p>
        <ul>
          <li>✅ 服务端渲染的内容已经在 HTML 中</li>
          <li>❌ 客户端水合的内容不在 HTML 中</li>
          <li>✅ 这就是 SSR 的优势：首屏内容立即可见</li>
        </ul>
      </div>
    </section>

    <NuxtLink to="/" class="back-link">← 返回首页</NuxtLink>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useHead } from '#app';

// 设置页面元数据
useHead({
  title: 'SSR 水合演示 - Nuxt 4 SSR Demo',
  meta: [
    { name: 'description', content: '展示 Nuxt 4 SSR 的水合（Hydration）机制' }
  ]
});

// 服务端数据（在服务端和客户端都会执行）
const serverTime = new Date().toLocaleString('zh-CN');
const renderEnvironment = process.server ? '服务器端' : '客户端';
const serverRandomNumber = Math.floor(Math.random() * 1000);

// 客户端数据（只在客户端执行）
const clientTime = ref('');
const clientRandomNumber = ref('');
const isHydrated = ref(false);
const clickCount = ref(0);
const lifecycleLogs = ref([]);

// 添加生命周期日志
const addLog = (type, message) => {
  lifecycleLogs.value.push({
    time: new Date().toLocaleTimeString('zh-CN'),
    type,
    message
  });
};

// 服务端和客户端都会执行
addLog('setup', 'setup() 函数执行');

// 客户端水合
onMounted(() => {
  addLog('client', 'onMounted() 钩子执行（仅客户端）');
  
  // 设置客户端数据
  clientTime.value = new Date().toLocaleString('zh-CN');
  clientRandomNumber.value = Math.floor(Math.random() * 1000);
  isHydrated.value = true;
  
  addLog('client', '客户端水合完成');
});

// 交互函数
const incrementCount = () => {
  clickCount.value++;
  addLog('interaction', `按钮被点击，计数: ${clickCount.value}`);
};
</script>

<style scoped>
.ssr-demo-container {
  max-width: 900px;
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
  line-height: 1.6;
}

.info-card strong {
  color: #333;
}

.hint {
  margin-top: 15px;
  padding: 10px;
  background: #f0f7ff;
  border-left: 3px solid #667eea;
  font-size: 0.95rem;
  color: #555;
}

.status-success {
  color: #22c55e;
  font-weight: bold;
}

.status-pending {
  color: #f59e0b;
  font-weight: bold;
}

.demo-button {
  margin-top: 15px;
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

.demo-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

.demo-button:active {
  transform: translateY(0);
}

.lifecycle-log {
  max-height: 300px;
  overflow-y: auto;
  background: #f8f9fa;
  padding: 15px;
  border-radius: 6px;
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
  min-width: 100px;
  text-align: center;
}

.log-setup {
  background: #dbeafe;
  color: #1e40af;
}

.log-client {
  background: #dcfce7;
  color: #166534;
}

.log-interaction {
  background: #fef3c7;
  color: #92400e;
}

.log-message {
  color: #374151;
  flex: 1;
}

.info-card ul {
  margin: 15px 0;
  padding-left: 20px;
}

.info-card li {
  margin: 8px 0;
  line-height: 1.6;
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
