<template>
  <div class="flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <div class="w-full max-w-md space-y-8 bg-white p-10 rounded-xl shadow-lg">
      <div>
        <h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">Create a new account</h2>
         <p class="mt-2 text-center text-sm text-gray-600">
          Already have an account?
          <router-link to="/login" class="font-medium text-teal-600 hover:text-teal-500">Sign in here</router-link>
        </p>
      </div>
      <RegisterForm />
    </div>
  </div>
</template>

<script setup>
import RegisterForm from '../components/auth/RegisterForm.vue';
</script>