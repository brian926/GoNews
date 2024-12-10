import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import TopNews from '../views/TopNews.vue';
import Sources from '../views/Sources.vue';
import Countries from '../views/Countries.vue';

const routes = [
  { path: '/', name: 'Home', component: Home },
  { path: '/top-news', name: 'TopNews', component: TopNews },
  { path: '/sources', name: 'Sources', component: Sources },
  { path: '/countries', name: 'Countries', component: Countries },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
