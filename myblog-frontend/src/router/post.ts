import { RouteRecordRaw } from "vue-router";
export const postRoutes :RouteRecordRaw[] = [
  {
    path: 'new',
    component: () => import('../views/Post/components/forms/PostForm.vue'),
    name: 'post.create'
  },
  {
    path: ':id',
    component: () => import('../views/Post/components/PostDetail.vue'),
    props: (route) => ({ id: parseInt(route.params.id as string) }),
    name: 'post.id'
  },
]