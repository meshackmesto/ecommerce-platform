// File: frontend/src/components/common/Header.vue
<template>
  <header class="bg-white/80 backdrop-blur-md shadow-sm sticky top-0 z-50">
    <nav class="container mx-auto px-4 py-4 flex justify-between items-center">
      <router-link to="/" class="text-2xl font-bold text-gray-800">
        Gadget & Grove
      </router-link>
      <div class="hidden md:flex items-center space-x-8">
        <router-link to="/" class="text-gray-600 hover:text-teal-500 transition">Home</router-link>
        <router-link to="/products" class="text-gray-600 hover:text-teal-500 transition">Shop</router-link>
      </div>
      <div class="flex items-center space-x-6">
        <router-link to="/cart" class="relative text-gray-600 hover:text-teal-500 transition">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" /></svg>
          <span v-if="cartStore.cartItemCount > 0" class="absolute -top-2 -right-2 bg-teal-500 text-white text-xs rounded-full h-5 w-5 flex items-center justify-center">{{ cartStore.cartItemCount }}</span>
        </router-link>

        <div v-if="authStore.isAuthenticated" class="flex items-center space-x-4">
          <router-link to="/profile" class="font-semibold text-gray-600 hover:text-teal-500">{{ authStore.user.first_name }}</router-link>
          <router-link v-if="authStore.isAdmin" to="/admin" class="text-sm bg-teal-500 text-white px-3 py-1 rounded-full">Admin</router-link>
          <button @click="handleLogout" class="text-gray-600 hover:text-teal-500">Logout</button>
        </div>
        <div v-else class="hidden md:flex items-center space-x-4">
          <router-link to="/login" class="text-gray-600 font-semibold hover:text-teal-500">Login</router-link>
          <router-link to="/register" class="bg-gray-800 text-white px-5 py-2 rounded-lg font-semibold hover:bg-gray-900">Register</router-link>
        </div>
      </div>
    </nav>
  </header>
</template>

<script setup>
import { useAuthStore } from '../../store/modules/auth';
import { useCartStore } from '../../store/modules/cart';
import { useRouter } from 'vue-router';

const authStore = useAuthStore();
const cartStore = useCartStore();
const router = useRouter();

const handleLogout = () => {
  authStore.logout();
  router.push('/login');
};
</script>