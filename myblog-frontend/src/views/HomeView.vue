<script lang="ts" setup>
import PostList from '../components/PostsList/index.vue'
import { AuthService } from '../services/auth';

const login = async() => {
  try {
    await AuthService.login({email: "correo@gmail.com", password: "1234"})
  } catch (error: Error) {
    console.log("No se pudo logear", error.message)
  }
}
</script>
<template>
  <PostList />
  <button @click="login">Login</button>
</template>