<script setup></script>

<template>
    <nav>
      <h1>My Blog</h1>
      <ul class="nav__menu">
        <li><router-link to="/">Home</router-link></li>
        <li><router-link to="/post">New post</router-link></li>
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

.nav__menu li a {
  color: var(--main-color);
  text-decoration: none;
}
</style>