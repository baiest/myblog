<script lang="ts" setup>
import PostCard from './PostCard.vue'
import { Post } from '../../models/Post'
import { Request } from '../../models/Request'
import { PostService } from '../../services/post'
import { onMounted, reactive } from 'vue';

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
  <!-- <PostCard v-for="post in post.data"
  :key="post.id"/> -->
  {{ posts }}
</template>