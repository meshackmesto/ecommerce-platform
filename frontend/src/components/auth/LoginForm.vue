<template>
  <form class="mt-8 space-y-6" @submit.prevent="handleLogin">
    <div v-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative" role="alert">
      <span class="block sm:inline">{{ error }}</span>
    </div>
    <div class="rounded-md shadow-sm -space-y-px">
      <div>
        <label for="email-address" class="sr-only">Email address</label>
        <input id="email-address" v-model="email" name="email" type="email" autocomplete="email" required class="relative block w-full appearance-none rounded-none rounded-t-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-teal-500 focus:outline-none focus:ring-teal-500 sm:text-sm" placeholder="Email address">
      </div>
      <div>
        <label for="password" class="sr-only">Password</label>
        <input id="password" v-model="password" name="password" type="password" autocomplete="current-password" required class="relative block w-full appearance-none rounded-none rounded-b-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-teal-500 focus:outline-none focus:ring-teal-500 sm:text-sm" placeholder="Password">
      </div>
    </div>

    <div>
      <button type="submit" :disabled="loading" class="group relative flex w-full justify-center rounded-md border border-transparent bg-gray-800 py-2 px-4 text-sm font-medium text-white hover:bg-gray-900 focus:outline-none focus:ring-2 focus:ring-teal-500 focus:ring-offset-2 disabled:bg-gray-400">
        {{ loading ? 'Signing in...' : 'Sign in' }}
      </button>
    </div>
  </form>
</template>

<script setup>
import { ref } from 'vue';
import { useAuthStore } from '../../store/modules/auth';
import { useRouter } from 'vue-router';

const email = ref('');
const password = ref('');
const error = ref(null);
const loading = ref(false);

const authStore = useAuthStore();
const router = useRouter();

const handleLogin = async () => {
  loading.value = true;
  error.value = null;
  try {
    await authStore.login({ email: email.value, password: password.value });
    if (authStore.isAdmin) {
      router.push('/admin');
    } else {
      router.push('/profile');
    }
  } catch (err) {
    error.value = err.message || 'Failed to login. Please check your credentials.';
  } finally {
    loading.value = false;
  }
};
</script>