import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    // ... add other routes here
    // e.g., { path: '/login', name: 'login', component: () => import('../views/Login.vue') }
  ]
})

export default router