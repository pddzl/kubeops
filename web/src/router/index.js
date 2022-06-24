import { createRouter, createWebHashHistory } from 'vue-router'

export const routes = [
  {
    path: '/login',
    name: 'login',
    hidden: true,
    component: () => import('@/view/login/index.vue')
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
