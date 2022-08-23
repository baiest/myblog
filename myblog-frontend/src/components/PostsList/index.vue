<script lang="ts" setup>
import PostCard from './PostCard.vue'
import { Post } from '../../models/Post'
import { Request } from '../../models/Request'
import { PostService } from '../../services/post'
import { onMounted, reactive, ref, watch } from 'vue';

const posts = reactive<Request<Post[]>>({
  loading: false,
  data: [],
  error: null
})

onMounted(async() => {
  try {
    posts.loading = true
    const response = await PostService.getPosts()
    posts.data = response
  } catch (error: Error) {
    posts.error = error.message
  }
  posts.loading = false
})
</script>
<template>
  <p v-if="posts.loading">Cargando...</p>
  <section class="posts__container">
    <PostCard v-for="post in posts.data"
  :key="post.id" :post="post"/>
  </section>
</template>

<style scoped>
.posts__container{
  display: grid;
  gap: 10px;
  grid-template-columns: repeat(auto-fit, 250px);
  justify-content: center;
  padding: 10px;
}
</style>

