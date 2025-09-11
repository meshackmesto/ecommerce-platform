// File: frontend/src/App.vue
<template>
  <div class="min-h-screen bg-gray-50 font-sans antialiased">
    <Header />
    <main>
      <router-view />
    </main>
    <Footer />
  </div>
</template>

<script setup>
import Header from './components/common/Header.vue';
import Footer from './components/common/Footer.vue';
</script>