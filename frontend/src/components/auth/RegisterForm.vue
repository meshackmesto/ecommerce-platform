<template>
  <form class="mt-8 space-y-6" @submit.prevent="handleRegister">
     <div v-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative" role="alert">
      <span class="block sm:inline">{{ error }}</span>
    </div>
    <div class="rounded-md shadow-sm space-y-4">
      <input v-model="firstName" type="text" required placeholder="First Name" class="input-field">
      <input v-model="lastName" type="text" required placeholder="Last Name" class="input-field">
      <input v-model="email" type="email" required placeholder="Email Address" class="input-field">
      <input v-model="phone" type="tel" placeholder="Phone Number (e.g. 254...)" class="input-field">
      <input v-model="password" type="password" required placeholder="Password (min 6 characters)" class="input-field">
    </div>

    <div>
      <button type="submit" :disabled="loading" class="group relative flex w-full justify-center rounded-md border border-transparent bg-gray-800 py-2 px-4 text-sm font-medium text-white hover:bg-gray-900 focus:outline-none focus:ring-2 focus:ring-teal-500 focus:ring-offset-2 disabled:bg-gray-400">
        {{ loading ? 'Creating Account...' : 'Create Account' }}
      </button>
    </div>
  </form>
</template>

<script setup>
import { ref } from 'vue';
import { useAuthStore } from '../../store/modules/auth';
import { useRouter } from 'vue-router';

const firstName = ref('');
const lastName = ref('');
const email = ref('');
const phone = ref('');
const password = ref('');
const error = ref(null);
const loading = ref(false);

const authStore = useAuthStore();
const router = useRouter();

const handleRegister = async () => {
  loading.value = true;
  error.value = null;
  try {
    await authStore.register({
      first_name: firstName.value,
      last_name: lastName.value,
      email: email.value,
      phone: phone.value,
      password: password.value,
    });
    router.push('/');
  } catch (err) {
    error.value = err.message || 'Registration failed.';
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
.input-field {
    @apply relative block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-teal-500 focus:outline-none focus:ring-teal-500 sm:text-sm;
}
</style>