// File: frontend/src/components/admin/Dashboard.vue
<template>
    <div>
        <h1 class="text-3xl font-bold text-gray-800">Admin Dashboard</h1>
        <p class.="text-gray-600 mt-2">Welcome back, {{ authStore.user?.first_name }}!</p>

        <div class="mt-8 grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="bg-white p-6 rounded-lg shadow">
                <h3 class="text-lg font-semibold text-gray-500">Total Sales</h3>
                <p class="text-3xl font-bold text-gray-800 mt-2">KES 0.00</p>
            </div>
             <div class="bg-white p-6 rounded-lg shadow">
                <h3 class="text-lg font-semibold text-gray-500">New Orders</h3>
                <p class="text-3xl font-bold text-gray-800 mt-2">0</p>
            </div>
             <div class="bg-white p-6 rounded-lg shadow">
                <h3 class="text-lg font-semibold text-gray-500">Total Products</h3>
                <p class="text-3xl font-bold text-gray-800 mt-2">0</p>
            </div>
        </div>
        <p class="mt-6 text-sm text-gray-500">Summary analytics will be displayed here.</p>
    </div>
</template>

<script setup>
import { useAuthStore } from '../../store/modules/auth';
const authStore = useAuthStore();
</script>