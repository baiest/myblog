import { createRouter,  createWebHistory, RouteRecordRaw} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import { postRoutes } from './post'

const routes:RouteRecordRaw[] = [
  {
    path: '/',
    component: HomeView,
    name: 'home'
  },
  {
    path: '/post',
    component: () => import('../views/PostView.vue'),
    name: 'post',
    children: postRoutes
  },

]
export const router = createRouter({
  history: createWebHistory(),
  routes
})