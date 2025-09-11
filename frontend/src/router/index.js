// File: frontend/src/router/index.js (Corrected Usage)
import { createRouter, createWebHistory } from 'vue-router';
// We can import the store function, but we CANNOT call it here.
import { useAuthStore } from '../store/modules/auth'; 

const routes = [
  // ... (Your complete list of routes)
  { path: '/', name: 'Home', component: () => import('../views/Home.vue') },
  { path: '/login', name: 'Login', component: () => import('../views/Login.vue') },
  { path: '/register', name: 'Register', component: () => import('../views/Register.vue') },
  { path: '/products', name: 'Products', component: () => import('../views/Products.vue') },
  { path: '/profile', name: 'Profile', component: () => import('../views/Profile.vue'), meta: { requiresAuth: true } },
  {
    path: '/admin',
    component: () => import('../views/admin/AdminDashboard.vue'),
    meta: { requiresAuth: true, requiresAdmin: true },
    redirect: { name: 'AdminOverview' },
    children: [
      { path: 'dashboard', name: 'AdminOverview', component: () => import('../components/admin/Dashboard.vue') },
      { path: 'products', name: 'AdminProducts', component: () => import('../components/admin/ProductManagement.vue') },
      { path: 'orders', name: 'AdminOrders', component: () => import('../components/admin/OrderManagement.vue') },
    ]
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior: () => ({ top: 0 }),
});

router.beforeEach((to, from, next) => {
  // CORRECT: The store is instantiated here, *inside* the guard function.
  // This function only runs *after* the app is set up and the user navigates.
  const authStore = useAuthStore();
  
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next({ name: 'Login' });
  } else if (to.meta.requiresAdmin && !authStore.isAdmin) {
    next({ name: 'Home' });
  } else {
    next();
  }
});

export default router;