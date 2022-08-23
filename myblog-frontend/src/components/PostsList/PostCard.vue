<script lang="ts" setup>
import { Post } from '../../models/Post';
const { post } = defineProps<{
  post: Post
}>()
</script>
<template>
  <router-link :to="{ name: 'post-id', params: { id:post.id }}" class="card">
    <h2 class="card__title">{{ post.title }}</h2>
    <section class="card__body">  
      <p>{{ post.content }}</p>
      <small>Autor: {{ post.id_user }}</small>
      <small>{{ post.created_at }}</small>
    </section>
  </router-link>
</template>

<style scoped>
  .card {
    border-radius: 12px;
    background: var(--card-background-color);
    overflow: hidden;
    position: relative;
    text-decoration: none;
    color: var(--main-color);
  }
  .card__title{
    margin: 0;
    position: sticky;
    top: 0;
    left: 0;
    right: 0;
    padding: 10px;
    background: var(--second-background-color);
  }
  
  .card__body{
    padding: 10px;
    padding-top: 0;
  }
</style>