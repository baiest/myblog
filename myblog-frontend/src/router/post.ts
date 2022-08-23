import { RouteRecordRaw } from "vue-router";
export const postRoutes :RouteRecordRaw[] = [
  {
    path: 'new',
    component: () => import('../components/forms/PostForm.vue'),
    name: 'post.create'
  },
  {
    path: ':id',
    component: () => import('../views/PostView.vue'),
    name: 'post-id'
  },
]