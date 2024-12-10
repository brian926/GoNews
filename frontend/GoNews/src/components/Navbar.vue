<template>
  <nav>
    <ul>
      <li><router-link to="/">Home</router-link></li>
      <li><router-link to="/top-news">Top News</router-link></li>
      <li><router-link to="/sources">Sources</router-link></li>
      <li><router-link to="/countries">Countries</router-link></li>
    </ul>
  </nav>
</template>

<script>
import '@/assets/styles/navbar.css';

export default {
  name: 'Navbar',
};
</script>
