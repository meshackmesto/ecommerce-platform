// File: frontend/src/components/admin/Dashboard.vue
<template>
    <div>
        <header class="mb-8">
            <h1 class="text-4xl font-bold text-foreground">Dashboard</h1>
            <p class="text-muted-foreground mt-2">Welcome to your dashboard, {{ authStore.user?.first_name }}</p>
        </header>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="bg-card p-6 rounded-lg shadow-custom-md">
                <p class="text-sm font-medium text-muted-foreground uppercase tracking-wider">Total Sales</p>
                <p class="text-3xl font-bold text-foreground mt-2">KES 0.00</p>
                <div class="w-full bg-secondary h-1 mt-4 rounded-full">
                    <div class="bg-primary h-1 rounded-full w-1/4"></div>
                </div>
            </div>
             <div class="bg-card p-6 rounded-lg shadow-custom-md">
                <p class="text-sm font-medium text-muted-foreground uppercase tracking-wider">New Orders</p>
                <p class="text-3xl font-bold text-foreground mt-2">0</p>
                <div class="w-full bg-secondary h-1 mt-4 rounded-full">
                    <div class="bg-accent h-1 rounded-full w-1/2"></div>
                </div>
            </div>
             <div class="bg-card p-6 rounded-lg shadow-custom-md">
                <p class="text-sm font-medium text-muted-foreground uppercase tracking-wider">Total Products</p>
                <p class="text-3xl font-bold text-foreground mt-2">{{ productStore.products.length }}</p>
                 <div class="w-full bg-secondary h-1 mt-4 rounded-full">
                    <div class="bg-success h-1 rounded-full w-3/4"></div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { onMounted } from 'vue';
import { useAuthStore } from '../../store/modules/auth';
import { useProductStore } from '../../store/modules/products';
const authStore = useAuthStore();
const productStore = useProductStore();

onMounted(() => {
    // Ensure products are loaded for the count
    if (productStore.products.length === 0) {
        productStore.fetchProducts();
    }
});
</script>