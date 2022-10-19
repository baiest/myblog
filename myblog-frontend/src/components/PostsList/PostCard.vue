<script lang="ts" setup>
import { computed } from '@vue/reactivity';
import { Post } from '../../models/Post';
import { PostService } from '../../services/post'

const { post } = defineProps<{
  post: Post
}>()

defineEmits(['delete'])

const localDate = computed(() => post.created_at.toLocaleString())
</script>
<template>
  <article class="card">
    <div class="card__header">
      <h2 class="card__title">{{ post.title }}</h2>
      <button class="card__delete" @click="$emit('delete')">Borrar</button>
    </div>
    <router-link class="card__body" :to="{ name: 'post.id', params: { id:post.id }}">  
        <div class="card__content" v-html="post.content"></div>
        <span class="card__info">
          <small>Autor: {{ post.user.name }}</small>
          <small>Creado: {{ localDate }}</small>
        </span>
    </router-link>
  </article>
</template>

<style scoped>
  .card {
    background: var(--card-background-color);
    overflow: hidden;
    position: relative;
    display: flex;
    flex-direction: column;
  }
  .card a {
    text-decoration: none;
    color: var(--main-color);
  }
  
  .card__header{
    display: flex;
    justify-content: space-between;
    align-items: center;
    background: var(--second-background-color);
    position: sticky;
    padding: 10px;
    margin: 0;
  }
  .card__title{
    margin: 0;
    top: 0;
    left: 0;
    right: 0;
  }
  
  .card__delete{
    width: fit-content;
  }
  
  .card__body{
    padding: 10px;
    padding-top: 0;
    height: 100%;
    display: flex;
    gap: 10px;
    flex-direction: column;
    justify-content: space-between;
  }
  .card__content {
    text-overflow: ellipsis;
    overflow: hidden;
    font-size: 1.6rem;
    line-height: 2rem;
    display: -webkit-box;
    -webkit-line-clamp: 5;
    -webkit-box-orient: vertical;
  }
  .card__info{
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: flex-end;
    font-weight: 600;
  }
</style>