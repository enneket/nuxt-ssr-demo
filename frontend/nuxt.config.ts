// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  modules: ['@nuxtjs/sitemap'],
  site: { url: 'https://site.com' },
  runtimeConfig: {
    public: {
      apiBaseUrl: 'http://localhost:8000/api' // 使用默认值，部署时可通过环境变量覆盖
    }
  },
  sitemap: {
    // 使用urls数组配置路由，这是v7版本的正确方式
    urls: async () => {
      interface Article {
        id: number;
        title: string;
        excerpt: string;
        content: string;
        image: string;
        created: string;
      }
      
      // 使用any类型绕过类型检查，解决SitemapUrlInput类型兼容问题
      const urls: any[] = [
        {
          loc: '/',
          changefreq: 'daily',
          priority: 0.5
        },
        {
          loc: '/articles',
          changefreq: 'daily',
          priority: 0.6
        }
      ];
      
      // 添加动态文章路由
      try {
        // 在构建时使用配置的API地址，从runtimeConfig获取
        const apiBaseUrl = 'http://localhost:8000/api'; // 从runtimeConfig获取或使用默认值
        const apiUrl = `${apiBaseUrl}/articles`;
        const response = await fetch(apiUrl);
        if (response.ok) {
          const articles: Article[] = await response.json();
          // 为每个文章生成站点地图条目
          articles.forEach(article => {
            urls.push({
              loc: `/articles/${article.id}`,
              changefreq: 'weekly',
              priority: 0.8,
              lastmod: article.created ? new Date(article.created).toISOString() : undefined
            });
          });
        }
      } catch (error) {
        console.error('Error generating sitemap routes:', error);
        // API调用失败时，只使用静态路由
      }
      
      return urls;
    }
  }
})