<script setup></script>

<template>
    <nav>
      <h1><router-link :to="{ name: 'home' }">My Blog</router-link></h1>
      <ul class="nav__menu">
        <li><router-link :to="{ name: 'home' }">Home</router-link></li>
        <li><router-link :to="{ name: 'post.create' }">New post</router-link></li>
        <li><router-link :to="{ name: 'post.create' }">My posts</router-link></li>
      </ul>
    </nav>
</template>

<style scoped>
nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  background: var(--second-background-color);
}

.nav__menu {
  display: flex;
  gap: 10px;
  margin: 0;
  padding: 0;
  list-style: none;
}


a {
  color: var(--main-color);
  text-decoration: none;
}
</style>