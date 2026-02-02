<template>
  <div class="performance-container">
    <h1>⚡ SSR 性能监控面板</h1>

    <!-- 性能指标卡片 -->
    <section class="metrics-grid">
      <div class="metric-card">
        <div class="metric-icon">🚀</div>
        <div class="metric-label">TTFB</div>
        <div class="metric-value">{{ metrics.ttfb }}</div>
        <div class="metric-desc">首字节时间</div>
      </div>

      <div class="metric-card">
        <div class="metric-icon">🎨</div>
        <div class="metric-label">FCP</div>
        <div class="metric-value">{{ metrics.fcp }}</div>
        <div class="metric-desc">首次内容绘制</div>
      </div>

      <div class="metric-card">
        <div class="metric-icon">📊</div>
        <div class="metric-label">LCP</div>
        <div class="metric-value">{{ metrics.lcp }}</div>
        <div class="metric-desc">最大内容绘制</div>
      </div>

      <div class="metric-card">
        <div class="metric-icon">⚡</div>
        <div class="metric-label">TTI</div>
        <div class="metric-value">{{ metrics.tti }}</div>
        <div class="metric-desc">可交互时间</div>
      </div>
    </section>

    <!-- 渲染模式对比 -->
    <section class="demo-section">
      <h2>🔄 SSR vs CSR 对比</h2>
      <div class="info-card">
        <table class="comparison-table">
          <thead>
            <tr>
              <th>指标</th>
              <th>SSR（本页面）</th>
              <th>CSR（传统 SPA）</th>
              <th>优势</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>TTFB</td>
              <td :class="getPerformanceClass(metrics.ttfb)">{{ metrics.ttfb }}</td>
              <td class="csr-value">~200ms</td>
              <td>{{ compareMetric('ttfb', metrics.ttfb, 200) }}</td>
            </tr>
            <tr>
              <td>FCP</td>
              <td :class="getPerformanceClass(metrics.fcp)">{{ metrics.fcp }}</td>
              <td class="csr-value">~1500ms</td>
              <td>{{ compareMetric('fcp', metrics.fcp, 1500) }}</td>
            </tr>
            <tr>
              <td>LCP</td>
              <td :class="getPerformanceClass(metrics.lcp)">{{ metrics.lcp }}</td>
              <td class="csr-value">~2500ms</td>
              <td>{{ compareMetric('lcp', metrics.lcp, 2500) }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </section>

    <!-- 性能详情 -->
    <section class="demo-section">
      <h2>📈 详细性能数据</h2>
      <div class="info-card">
        <div class="detail-grid">
          <div class="detail-item">
            <strong>页面加载时间：</strong>
            <span>{{ metrics.loadTime }}</span>
          </div>
          <div class="detail-item">
            <strong>DOM 解析时间：</strong>
            <span>{{ metrics.domParse }}</span>
          </div>
          <div class="detail-item">
            <strong>资源加载时间：</strong>
            <span>{{ metrics.resourceLoad }}</span>
          </div>
          <div class="detail-item">
            <strong>水合时间：</strong>
            <span>{{ metrics.hydrationTime }}</span>
          </div>
        </div>
      </div>
    </section>

    <!-- 性能建议 -->
    <section class="demo-section">
      <h2>💡 SSR 性能优势</h2>
      <div class="info-card">
        <ul class="benefits-list">
          <li>
            <span class="benefit-icon">✅</span>
            <div>
              <strong>更快的首屏渲染</strong>
              <p>服务器直接返回完整 HTML，用户立即看到内容</p>
            </div>
          </li>
          <li>
            <span class="benefit-icon">✅</span>
            <div>
              <strong>更好的 SEO</strong>
              <p>搜索引擎可以直接抓取完整页面内容</p>
            </div>
          </li>
          <li>
            <span class="benefit-icon">✅</span>
            <div>
              <strong>降低客户端负担</strong>
              <p>减少客户端 JavaScript 执行时间</p>
            </div>
          </li>
          <li>
            <span class="benefit-icon">✅</span>
            <div>
              <strong>更好的用户体验</strong>
              <p>特别是在低配置设备和慢速网络环境下</p>
            </div>
          </li>
        </ul>
      </div>
    </section>

    <!-- 性能评分 -->
    <section class="demo-section">
      <h2>🎯 性能评分</h2>
      <div class="info-card">
        <div class="score-container">
          <div class="score-circle" :class="getScoreClass(performanceScore)">
            <div class="score-value">{{ performanceScore }}</div>
            <div class="score-label">分</div>
          </div>
          <div class="score-description">
            <p><strong>{{ getScoreText(performanceScore) }}</strong></p>
            <p>{{ getScoreDescription(performanceScore) }}</p>
          </div>
        </div>
      </div>
    </section>

    <NuxtLink to="/" class="back-link">← 返回首页</NuxtLink>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useHead } from '#app';

// 设置页面元数据
useHead({
  title: '性能监控 - Nuxt 4 SSR Demo',
  meta: [
    { name: 'description', content: '展示 Nuxt 4 SSR 的性能指标和优势' }
  ]
});

// 性能指标
const metrics = ref({
  ttfb: '计算中...',
  fcp: '计算中...',
  lcp: '计算中...',
  tti: '计算中...',
  loadTime: '计算中...',
  domParse: '计算中...',
  resourceLoad: '计算中...',
  hydrationTime: '计算中...'
});

// 性能评分
const performanceScore = ref(0);

// 格式化时间
const formatTime = (ms) => {
  if (ms === null || ms === undefined || isNaN(ms)) return 'N/A';
  return `${Math.round(ms)}ms`;
};

// 获取性能等级样式
const getPerformanceClass = (value) => {
  if (value === '计算中...' || value === 'N/A') return '';
  const ms = parseInt(value);
  if (ms < 500) return 'perf-good';
  if (ms < 1000) return 'perf-ok';
  return 'perf-poor';
};

// 对比指标
const compareMetric = (type, ssrValue, csrValue) => {
  if (ssrValue === '计算中...' || ssrValue === 'N/A') return '-';
  const ssr = parseInt(ssrValue);
  const improvement = Math.round(((csrValue - ssr) / csrValue) * 100);
  return improvement > 0 ? `快 ${improvement}%` : '相当';
};

// 获取评分等级
const getScoreClass = (score) => {
  if (score >= 90) return 'score-excellent';
  if (score >= 70) return 'score-good';
  if (score >= 50) return 'score-ok';
  return 'score-poor';
};

// 获取评分文本
const getScoreText = (score) => {
  if (score >= 90) return '优秀';
  if (score >= 70) return '良好';
  if (score >= 50) return '一般';
  return '需要优化';
};

// 获取评分描述
const getScoreDescription = (score) => {
  if (score >= 90) return 'SSR 性能表现出色，用户体验极佳';
  if (score >= 70) return 'SSR 性能良好，大部分指标达标';
  if (score >= 50) return 'SSR 性能一般，仍有优化空间';
  return 'SSR 性能需要优化，建议检查服务器配置';
};

// 计算性能评分
const calculateScore = () => {
  const ttfb = parseInt(metrics.value.ttfb) || 0;
  const fcp = parseInt(metrics.value.fcp) || 0;
  const lcp = parseInt(metrics.value.lcp) || 0;
  
  let score = 100;
  
  // TTFB 评分（权重 30%）
  if (ttfb > 600) score -= 30;
  else if (ttfb > 300) score -= 15;
  else if (ttfb > 100) score -= 5;
  
  // FCP 评分（权重 35%）
  if (fcp > 3000) score -= 35;
  else if (fcp > 1800) score -= 20;
  else if (fcp > 1000) score -= 10;
  
  // LCP 评分（权重 35%）
  if (lcp > 4000) score -= 35;
  else if (lcp > 2500) score -= 20;
  else if (lcp > 1500) score -= 10;
  
  performanceScore.value = Math.max(0, Math.round(score));
};

// 客户端挂载后获取性能数据
onMounted(() => {
  // 使用 Performance API 获取性能数据
  if (window.performance) {
    const perfData = window.performance.timing;
    const navigation = window.performance.getEntriesByType('navigation')[0];
    
    // 计算各项指标
    metrics.value.ttfb = formatTime(perfData.responseStart - perfData.requestStart);
    metrics.value.loadTime = formatTime(perfData.loadEventEnd - perfData.navigationStart);
    metrics.value.domParse = formatTime(perfData.domContentLoadedEventEnd - perfData.domLoading);
    metrics.value.resourceLoad = formatTime(perfData.loadEventEnd - perfData.domContentLoadedEventEnd);
    
    // 使用 PerformanceObserver 获取 FCP 和 LCP
    if ('PerformanceObserver' in window) {
      try {
        // FCP
        const fcpObserver = new PerformanceObserver((list) => {
          for (const entry of list.getEntries()) {
            if (entry.name === 'first-contentful-paint') {
              metrics.value.fcp = formatTime(entry.startTime);
            }
          }
        });
        fcpObserver.observe({ entryTypes: ['paint'] });
        
        // LCP
        const lcpObserver = new PerformanceObserver((list) => {
          const entries = list.getEntries();
          const lastEntry = entries[entries.length - 1];
          metrics.value.lcp = formatTime(lastEntry.startTime);
        });
        lcpObserver.observe({ entryTypes: ['largest-contentful-paint'] });
        
        // TTI（简化估算）
        setTimeout(() => {
          metrics.value.tti = formatTime(performance.now());
          metrics.value.hydrationTime = formatTime(performance.now() - (perfData.domContentLoadedEventEnd - perfData.navigationStart));
          
          // 计算评分
          calculateScore();
        }, 100);
      } catch (e) {
        console.error('Performance Observer error:', e);
      }
    }
    
    // 如果 PerformanceObserver 不可用，使用备用方案
    setTimeout(() => {
      if (metrics.value.fcp === '计算中...') {
        metrics.value.fcp = formatTime(perfData.domContentLoadedEventEnd - perfData.navigationStart);
      }
      if (metrics.value.lcp === '计算中...') {
        metrics.value.lcp = formatTime(perfData.loadEventEnd - perfData.navigationStart);
      }
      if (metrics.value.tti === '计算中...') {
        metrics.value.tti = formatTime(performance.now());
      }
      if (metrics.value.hydrationTime === '计算中...') {
        metrics.value.hydrationTime = formatTime(performance.now() - (perfData.domContentLoadedEventEnd - perfData.navigationStart));
      }
      
      // 计算评分
      calculateScore();
    }, 1000);
  }
});
</script>

<style scoped>
.performance-container {
  max-width: 1100px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  color: #333;
  text-align: center;
  margin-bottom: 30px;
}

.metrics-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

.metric-card {
  background: white;
  padding: 25px;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  text-align: center;
  transition: all 0.3s ease;
}

.metric-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.15);
}

.metric-icon {
  font-size: 2.5rem;
  margin-bottom: 10px;
}

.metric-label {
  font-size: 0.9rem;
  color: #6b7280;
  margin-bottom: 8px;
  font-weight: 600;
}

.metric-value {
  font-size: 1.8rem;
  font-weight: bold;
  color: #667eea;
  margin-bottom: 5px;
}

.metric-desc {
  font-size: 0.85rem;
  color: #9ca3af;
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

.comparison-table {
  width: 100%;
  border-collapse: collapse;
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

.perf-good {
  color: #22c55e;
  font-weight: bold;
}

.perf-ok {
  color: #f59e0b;
  font-weight: bold;
}

.perf-poor {
  color: #ef4444;
  font-weight: bold;
}

.csr-value {
  color: #9ca3af;
}

.detail-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 15px;
}

.detail-item {
  padding: 15px;
  background: #f8f9fa;
  border-radius: 6px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.detail-item strong {
  color: #374151;
}

.detail-item span {
  color: #667eea;
  font-weight: bold;
}

.benefits-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.benefits-list li {
  display: flex;
  gap: 15px;
  padding: 15px 0;
  border-bottom: 1px solid #e5e7eb;
}

.benefits-list li:last-child {
  border-bottom: none;
}

.benefit-icon {
  font-size: 1.5rem;
  flex-shrink: 0;
}

.benefits-list strong {
  color: #374151;
  display: block;
  margin-bottom: 5px;
}

.benefits-list p {
  color: #6b7280;
  margin: 0;
  font-size: 0.95rem;
}

.score-container {
  display: flex;
  gap: 30px;
  align-items: center;
  flex-wrap: wrap;
}

.score-circle {
  width: 150px;
  height: 150px;
  border-radius: 50%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border: 8px solid;
  flex-shrink: 0;
}

.score-excellent {
  border-color: #22c55e;
  background: #f0fdf4;
}

.score-good {
  border-color: #3b82f6;
  background: #eff6ff;
}

.score-ok {
  border-color: #f59e0b;
  background: #fffbeb;
}

.score-poor {
  border-color: #ef4444;
  background: #fef2f2;
}

.score-value {
  font-size: 3rem;
  font-weight: bold;
  color: inherit;
}

.score-label {
  font-size: 1rem;
  color: #6b7280;
}

.score-description {
  flex: 1;
}

.score-description p {
  margin: 8px 0;
  line-height: 1.6;
}

.score-description strong {
  font-size: 1.3rem;
  color: #374151;
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
