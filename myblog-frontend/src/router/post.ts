import { RouteRecordRaw } from "vue-router";
export const postRoutes :RouteRecordRaw[] = [
  {
    path: 'new',
    component: () => import('../views/Post/components/forms/PostForm.vue'),
    name: 'post.create'
  },
  {
    path: ':id',
    component: () => import('../views/Post/PostView.vue'),
    name: 'post.id'
  },
]