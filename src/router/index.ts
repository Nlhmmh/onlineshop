import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Home from '../views/Home.vue'
import About from '../views/About.vue'
import Test from '../views/Test.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: {
      auth: false,
      title: 'Home'
    }
  },
  {
    path: '/about',
    name: 'About',
    component: About,
    meta: {
      auth: false,
      title: 'About'
    }
  },
  {
    path: '/test',
    name: 'Test',
    component: Test,
    meta: {
      auth: false,
      title: 'Test'
    }
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
