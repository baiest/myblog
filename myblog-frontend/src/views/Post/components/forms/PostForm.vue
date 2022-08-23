<script lang="ts" setup>
import { reactive, ref } from 'vue';
import { PostService } from '../../../../services/post';
import {PostDAO} from '../../../..//models/Post'
import './style.css'
import { useRouter } from 'vue-router';
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';

const form = reactive<PostDAO>({
  title: '',
  content: ''
})
const errorResponse = ref('')

const router = useRouter()

const submit = async () => {
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
      <input class="form__input" placeholder=" " id="title" type="text" v-model="form.title">
      <label class="form__label" placeholder=" " for="title">Título</label>
    </fieldset>
    <fieldset>
      <QuillEditor theme="snow" v-model:content="form.content" content-type="html"/>
    </fieldset>
    <button type="submit" class="form__submit">Crear</button>
  </form>
</template>

<style scoped>
  .form__input, .form__label{
    font-size: 5rem;
  }
  .form__input:focus+.form__label,
  .form__input:not(:placeholder-shown)+.form__label {
    transform: translateY(-4rem);
  }
  .ql-editor{
    font-size: 2rem;
  }
</style>