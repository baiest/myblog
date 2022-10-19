<script lang="ts" setup>
import { onMounted, reactive } from 'vue';
import { Post } from '../../../models/Post';
import { Request } from '../../../models/Request';
import { PostService } from '../../../services/post';

const {id} = defineProps<{
  id: number
}>()
const post = reactive<Request<Post | null>>({
  loading: false,
  data: null,
  error: null
})
onMounted(async() => {
  try {
    post.data = await PostService.getById(id)
  } catch (error: Error) {
    post.error = error.message
  }
})
</script>

<template>
<main>
  <p v-if="post.error">{{post.error}}</p>
  <h1>{{post.data?.title}}</h1>
  <p>Author: {{post.data?.user.name}}</p>
  <p>Creado: {{post.data?.created_at}}</p>
  <div v-html="post.data?.content" ></div>
</main>
</template>

<style scope>
  main {
    max-width: 120rem;
    margin: 0 auto;
  }
  
  main * {
    width: 100%;
    max-height: 450px;
    object-fit: contain;
  }
</style>