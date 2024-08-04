import { createRouter, createWebHistory } from 'vue-router';
import Login from '@/views/Login.vue';
import Protected from '@/views/Protected.vue';

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/protected',
    name: 'Protected',
    component: Protected,
    meta: { requiresAuth: true }
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    const token = sessionStorage.getItem('authToken');
    if (token && token !== "null") {
      next();
    } else {
      next('/login');
    }
  } else {
    next();
  }
});

export default router;
