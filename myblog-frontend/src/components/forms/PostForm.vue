<script lang="ts" setup>
import { reactive, ref } from 'vue';
import { PostService } from '../../services/post';
import {PostDAO} from '../../models/Post'
import './style.css'
import { useRouter } from 'vue-router';
const form = reactive<PostDAO>({
  title: '',
  content: ''
})

const errorResponse = ref('')

const router = useRouter()

const submit = async () => {
    errorResponse.value = "Enviado"
    try {
      await PostService.create(form)
      router.push({
        name: 'home'
      })
      errorResponse.value = "dasd"
    } catch (error: Error) {
      errorResponse.value = error.message
    }
}
</script>
<template>
  <p v-if="errorResponse">{{errorResponse}} dsadasd</p>
  <form @submit.prevent="submit">
    <fieldset>
      <input class="form__input" placeholder=" " type="text" v-model="form.title">
      <label class="form__label" placeholder=" " for="title">Título</label>
    </fieldset>
    <fieldset>
      <input class="form__input" placeholder=" " type="text" v-model="form.content">
      <label class="form__label" placeholder=" " for="title">Contenido</label>
    </fieldset>
    <button type="submit">Crear</button>
  </form>
</template>
