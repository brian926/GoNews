<template>
  <div>
    <h1>Top News</h1>
    <div class="controls">
      <label for="language">Select Language:</label>
      <select id="language" v-model="selectedLanguage" @change="fetchNews">
        <option v-for="lang in languages" :value="lang.value" :key="lang.value">{{ lang.label }}</option>
      </select>
    </div>
    <div>
      <label for="country">Select Country:</label>
      <select id="country" v-model="selectedCountry" @change="fetchNews">
        <option v-for="country in countries" :value="country.value" :key="country.value">{{ country.label }}</option>
      </select>
    </div>
    <div v-if="loading">Loading news...</div>
    <div v-if="error">{{ error }}</div>
    <div v-if="articles.length > 0" class="article-list">
      <div v-for="article in articles" :key="article.link" class="article-item">
        <img :src="article.image_url || 'https://via.placeholder.com/80x80'" :alt="article.title" class="thumbnail" />
        <div class="article-details">
          <h2 class="title">
            <a :href="article.link" target="_blank">{{ article.title }}</a>
          </h2>
          <p class="description">{{ article.description || 'No description available' }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import '@/assets/styles/articles.css';

export default {
  name: 'TopNews',
  data() {
    return {
      articles: [],
      loading: false,
      error: null,
      selectedLanguage: 'en',
      selectedCountry: 'us',
      languages: [
        { value: 'en', label: 'English' },
        { value: 'es', label: 'Spanish' },
        { value: 'fr', label: 'French' },
        { value: 'de', label: 'German' },
        { value: 'pt', label: 'Portuguese' },
      ],
      countries: [
        { value: 'us', label: 'United States' },
        { value: 'br', label: 'Brazil' },
        { value: 'es', label: 'Spain' },
        { value: 'de', label: 'Germany' },
        { value: 'fr', label: 'France' },
      ],
    };
  },
  methods: {
    async fetchNews() {
      this.loading = true;
      this.error = null;
      try {
        const response = await fetch(
          `http://localhost:5432/news?lang=${this.selectedLanguage}&cat=${this.selectedCountry}`
        );
        if (!response.ok) {
          throw new Error('Failed to fetch news.');
        }
        this.articles = await response.json();
      } catch (error) {
        this.error = error.message;
      } finally {
        this.loading = false;
      }
    },
  },
  mounted() {
    this.fetchNews();
  },
};
</script>
