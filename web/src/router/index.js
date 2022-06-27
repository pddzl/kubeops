import { createRouter, createWebHashHistory } from 'vue-router'

/* Layout */
import Layout from '@/view/layout/index.vue'

export const routes = [
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        component: () => import('@/view/dashboard/index.vue'),
        name: 'Dashboard',
        meta: { title: 'Dashboard', affix: true }
      }
    ]
  },
  {
    path: '/login',
    name: 'Login',
    hidden: true,
    component: () => import('@/view/login/index.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
