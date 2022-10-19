<script lang="ts" setup>
import PostCard from './PostCard.vue'
import { Post } from '../../models/Post'
import { Request } from '../../models/Request'
import { PostService } from '../../services/post'
import { onMounted, reactive } from 'vue';
import { computed } from '@vue/reactivity';

const posts = reactive<Request<Post[]>>({
  loading: false,
  data: [],
  error: null
})

const noData = computed(() => !posts.loading && !!!posts.data.length)

const handleDelete = async (id: number) => {
  try {
    await PostService.delete(id)
    const postIndex = posts.data.findIndex(post => post.id === id)
    posts.data.splice(postIndex, 1)
  } catch (error: Error) {
    posts.error = error.message
  }
  
}

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
  <p v-if="posts.error">{{posts.error}}</p>
  <p v-if="posts.loading">Cargando...</p>
  <section class="posts__container">
    <PostCard v-for="post in posts.data"
    :key="post.id" :post="post" @delete="handleDelete(post.id)"/>
  </section>
  <p v-if="noData">No hay datos</p>
</template>

<style scoped>
.posts__container{
  display: grid;
  gap: 10px;
  grid-template-columns: repeat(auto-fit,  minmax(250px, 350px));
  grid-auto-rows: minmax(150px, 250px);
  justify-content: center;
  padding: 10px;
}
</style>

