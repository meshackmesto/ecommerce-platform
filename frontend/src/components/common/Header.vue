// File: frontend/src/components/common/Header.vue
<template>
  <Disclosure as="nav" class="bg-card/80 backdrop-blur-md shadow-sm sticky top-0 z-50" v-slot="{ open }">
    <div class="container mx-auto px-4">
      <div class="relative flex h-16 items-center justify-between">
        <div class="absolute inset-y-0 left-0 flex items-center sm:hidden">
          <DisclosureButton class="inline-flex items-center justify-center rounded-md p-2 text-muted-foreground hover:bg-secondary hover:text-foreground focus:outline-none focus:ring-2 focus:ring-inset focus:ring-primary">
            <span class="sr-only">Open main menu</span>
            <Bars3Icon v-if="!open" class="block h-6 w-6" aria-hidden="true" />
            <XMarkIcon v-else class="block h-6 w-6" aria-hidden="true" />
          </DisclosureButton>
        </div>
        <div class="flex flex-1 items-center justify-center sm:items-stretch sm:justify-start">
          <div class="flex flex-shrink-0 items-center">
            <router-link to="/" class="text-2xl font-bold text-foreground">
              Gadget & Grove
            </router-link>
          </div>
          <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
            <router-link to="/" class="nav-link">Home</router-link>
            <router-link to="/products" class="nav-link">Shop</router-link>
          </div>
        </div>
        <div class="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0">
          <router-link to="/cart" class="relative rounded-full p-1 text-muted-foreground hover:text-foreground">
             <ShoppingBagIcon class="h-6 w-6" aria-hidden="true" />
             <span v-if="cartStore.cartItemCount > 0" class="absolute -top-1 -right-1 bg-primary text-primary-foreground text-xs rounded-full h-4 w-4 flex items-center justify-center">{{ cartStore.cartItemCount }}</span>
          </router-link>
          <div v-if="authStore.isAuthenticated" class="hidden sm:flex items-center space-x-4 ml-4">
              <router-link to="/profile" class="font-semibold text-sm text-foreground hover:text-primary">{{ authStore.user.first_name }}</router-link>
              <router-link v-if="authStore.isAdmin" to="/admin" class="btn-secondary text-xs !py-1 !px-3">Admin</router-link>
              <button @click="handleLogout" class="text-sm text-muted-foreground hover:text-foreground">Logout</button>
            </div>
            <div v-else class="hidden sm:flex items-center space-x-2 ml-4">
              <router-link to="/login" class="btn-secondary !py-2 !px-4">Login</router-link>
              <router-link to="/register" class="btn-primary !py-2 !px-4">Register</router-link>
            </div>
        </div>
      </div>
    </div>

    <DisclosurePanel class="sm:hidden">
      <div class="space-y-1 px-2 pb-3 pt-2">
        <DisclosureButton as="router-link" to="/" class="nav-link-mobile">Home</DisclosureButton>
        <DisclosureButton as="router-link" to="/products" class="nav-link-mobile">Shop</DisclosureButton>
        <div v-if="!authStore.isAuthenticated" class="border-t border-border pt-4 mt-4 flex flex-col space-y-2">
            <DisclosureButton as="router-link" to="/login" class="btn-secondary w-full">Login</DisclosureButton>
            <DisclosureButton as="router-link" to="/register" class="btn-primary w-full">Register</DisclosureButton>
        </div>
         <div v-else class="border-t border-border pt-4 mt-4 space-y-2">
            <DisclosureButton as="router-link" to="/profile" class="nav-link-mobile">Profile</DisclosureButton>
            <DisclosureButton v-if="authStore.isAdmin" as="router-link" to="/admin" class="nav-link-mobile">Admin Panel</DisclosureButton>
            <button @click="handleLogout" class="nav-link-mobile text-left w-full">Logout</button>
        </div>
      </div>
    </DisclosurePanel>
  </Disclosure>
</template>

<script setup>
import { Disclosure, DisclosureButton, DisclosurePanel } from '@headlessui/vue'
import { Bars3Icon, XMarkIcon, ShoppingBagIcon } from '@heroicons/vue/24/outline'
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

<style scoped>
/* Here we use a scoped style block for component-specific styles */
.nav-link {
    @apply inline-flex items-center border-b-2 border-transparent px-1 pt-1 text-sm font-medium text-muted-foreground hover:border-border hover:text-foreground;
}
.router-link-exact-active.nav-link {
    @apply border-primary text-foreground;
}
.nav-link-mobile {
    @apply block rounded-md px-3 py-2 text-base font-medium text-muted-foreground hover:bg-secondary hover:text-foreground;
}
.router-link-exact-active.nav-link-mobile {
    @apply bg-accent-light text-accent;
}
</style>