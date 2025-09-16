// File: frontend/src/components/auth/LoginForm.vue
<template>
  <form class="mt-8 space-y-6" @submit.prevent="handleLogin">
    <div v-if="error" class="bg-destructive-light border border-destructive text-destructive px-4 py-3 rounded-md text-sm" role="alert">
      <span>{{ error }}</span>
    </div>
    <div class="space-y-4">
      <div>
        <label for="email-address" class="form-label">Email address</label>
        <input id="email-address" v-model="email" type="email" required class="form-input" placeholder="name@company.com">
      </div>
      <div>
        <label for="password" class="form-label">Password</label>
        <input id="password" v-model="password" type="password" required class="form-input" placeholder="••••••••">
      </div>
    </div>
    <div>
      <button type="submit" :disabled="loading" class="btn-primary w-full justify-center !py-3">
        {{ loading ? 'Signing in...' : 'Sign in' }}
      </button>
    </div>
  </form>
</template>

<script setup>
// Script content remains the same as before
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

<style scoped>
.form-label { @apply block mb-2 text-sm font-medium text-foreground; }
.form-input { @apply bg-secondary border border-border text-foreground text-sm rounded-lg focus:ring-primary focus:border-primary block w-full p-2.5; }
</style>