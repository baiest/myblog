import { RouteRecordRaw } from "vue-router";
export const postRoutes :RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('../components/forms/PostForm.vue'),
    name: 'post-create'
  },
  {
    path: ':id',
    component: () => import('../views/PostView.vue'),
    name: 'post-id'
  },
]