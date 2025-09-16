// File: frontend/src/views/Profile.vue
<template>
  <div class="container mx-auto px-4 py-12">
    <header class="text-center mb-12">
        <h1 class="text-4xl font-bold text-foreground">My Account</h1>
        <p class="text-muted-foreground mt-2">View your order history and manage your details.</p>
    </header>

    <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
        <div class="lg:col-span-1">
            <div class="bg-card p-6 rounded-lg shadow-custom-md">
                <h2 class="text-xl font-semibold text-card-foreground">{{ authStore.user?.first_name }} {{ authStore.user?.last_name }}</h2>
                <p class="text-sm text-muted-foreground">{{ authStore.user?.email }}</p>
                <div class="border-t border-border my-4"></div>
                <button class="text-sm font-medium text-primary hover:text-primary-hover transition-colors w-full text-left">
                    Edit Profile
                </button>
                 <button @click="handleLogout" class="mt-2 text-sm font-medium text-destructive hover:text-red-500 transition-colors w-full text-left">
                    Logout
                </button>
            </div>
        </div>

        <div class="lg:col-span-3">
            <OrderHistory />
        </div>
    </div>
  </div>
</template>

<script setup>
import { useAuthStore } from '../store/modules/auth';
import OrderHistory from '../components/order/OrderHistory.vue';
import { useRouter } from 'vue-router';

const authStore = useAuthStore();
const router = useRouter();

const handleLogout = () => {
  authStore.logout();
  router.push('/login');
};
</script>