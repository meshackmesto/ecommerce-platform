// File: frontend/src/views/admin/AdminDashboard.vue
<template>
  <div class="flex min-h-screen bg-background text-foreground">
    <aside class="w-64 bg-sidebar-background border-r border-sidebar-border flex-shrink-0">
      <div class="p-6">
        <h2 class="text-2xl font-bold text-sidebar-primary-foreground">Admin Panel</h2>
      </div>
      <nav class="mt-6 px-4 space-y-2">
        <router-link :to="{ name: 'AdminOverview' }" class="nav-link">
          <HomeIcon class="h-5 w-5 mr-3" />
          Dashboard
        </router-link>
        <router-link :to="{ name: 'AdminProducts' }" class="nav-link">
          <CubeIcon class="h-5 w-5 mr-3" />
          Products
        </router-link>
        <router-link :to="{ name: 'AdminOrders' }" class="nav-link">
          <ShoppingCartIcon class="h-5 w-5 mr-3" />
          Orders
        </router-link>
      </nav>
    </aside>
    <main class="flex-1 p-8 overflow-y-auto">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { HomeIcon, CubeIcon, ShoppingCartIcon } from '@heroicons/vue/24/outline';
</script>

<style scoped>
.nav-link {
  @apply flex items-center px-4 py-2.5 text-sm font-medium rounded-lg text-sidebar-foreground hover:bg-sidebar-accent hover:text-sidebar-accent-foreground transition-colors;
}
.router-link-exact-active {
  @apply bg-sidebar-primary text-sidebar-primary-foreground;
}
</style>